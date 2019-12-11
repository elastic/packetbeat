// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package input

import (
	"fmt"
	"sync"

	"github.com/elastic/beats/libbeat/processors/add_formatted_index"

	"github.com/elastic/beats/libbeat/common/fmtstr"

	"github.com/gofrs/uuid"

	"github.com/elastic/beats/journalbeat/checkpoint"
	"github.com/elastic/beats/journalbeat/reader"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/processors"
)

// Input manages readers and forwards entries from journals.
type Input struct {
	readers    []*reader.Reader
	done       chan struct{}
	config     Config
	pipeline   beat.Pipeline
	client     beat.Client
	states     map[string]checkpoint.JournalState
	id         uuid.UUID
	logger     *logp.Logger
	eventMeta  common.EventMetadata
	processors beat.ProcessorList
}

// New returns a new Inout
func New(
	c *common.Config,
	b *beat.Beat,
	done chan struct{},
	states map[string]checkpoint.JournalState,
) (*Input, error) {
	config := DefaultConfig
	if err := c.Unpack(&config); err != nil {
		return nil, err
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("error while generating ID for input: %v", err)
	}

	logger := logp.NewLogger("input").With("id", id)

	var readers []*reader.Reader
	if len(config.Paths) == 0 {
		cfg := reader.Config{
			Path:               reader.LocalSystemJournalID, // used to identify the state in the registry
			Backoff:            config.Backoff,
			MaxBackoff:         config.MaxBackoff,
			Seek:               config.Seek,
			CursorSeekFallback: config.CursorSeekFallback,
			Matches:            config.Matches,
			SaveRemoteHostname: config.SaveRemoteHostname,
		}

		state := states[reader.LocalSystemJournalID]
		r, err := reader.NewLocal(cfg, done, state, logger)
		if err != nil {
			return nil, fmt.Errorf("error creating reader for local journal: %v", err)
		}
		readers = append(readers, r)
	}

	for _, p := range config.Paths {
		cfg := reader.Config{
			Path:               p,
			Backoff:            config.Backoff,
			MaxBackoff:         config.MaxBackoff,
			Seek:               config.Seek,
			CursorSeekFallback: config.CursorSeekFallback,
			Matches:            config.Matches,
			SaveRemoteHostname: config.SaveRemoteHostname,
		}
		state := states[p]
		r, err := reader.New(cfg, done, state, logger)
		if err != nil {
			return nil, fmt.Errorf("error creating reader for journal: %v", err)
		}
		readers = append(readers, r)
	}

	inputProcessors, err := processorsForInput(b.Info, config)
	if err != nil {
		return nil, err
	}

	logger.Debugf("New input is created for paths %v", config.Paths)

	return &Input{
		readers:    readers,
		done:       done,
		config:     config,
		pipeline:   b.Publisher,
		states:     states,
		id:         id,
		logger:     logger,
		eventMeta:  config.EventMetadata,
		processors: inputProcessors,
	}, nil
}

// Run connects to the output, collects entries from the readers
// and then publishes the events.
func (i *Input) Run() {
	var err error
	i.client, err = i.pipeline.ConnectWith(beat.ClientConfig{
		PublishMode: beat.GuaranteedSend,
		Processing: beat.ProcessingConfig{
			EventMetadata: i.eventMeta,
			Meta:          nil,
			Processor:     i.processors,
		},
		ACKCount: func(n int) {
			i.logger.Infof("journalbeat successfully published %d events", n)
		},
	})
	if err != nil {
		i.logger.Error("Error connecting to output: %v", err)
		return
	}

	i.publishAll()
}

// publishAll reads events from all readers and publishes them.
func (i *Input) publishAll() {
	out := make(chan *beat.Event)
	defer close(out)

	var wg sync.WaitGroup
	defer wg.Wait()
	for _, r := range i.readers {
		wg.Add(1)
		r := r
		go func() {
			defer wg.Done()

			for {
				select {
				case <-i.done:
					return
				default:
				}

				event, err := r.Next()
				if event == nil {
					if err != nil {
						i.logger.Errorf("Error while reading event: %v", err)
					}
					continue
				}

				select {
				case <-i.done:
				case out <- event:
				}
			}
		}()
	}

	for {
		select {
		case <-i.done:
			return
		case e := <-out:
			i.client.Publish(*e)
		}
	}
}

// Stop stops all readers of the input.
func (i *Input) Stop() {
	for _, r := range i.readers {
		r.Close()
	}
	i.client.Close()
}

// Wait waits until all readers are done.
func (i *Input) Wait() {
	i.Stop()
}

func processorsForInput(beatInfo beat.Info, config Config) (*processors.Processors, error) {
	procs := processors.NewList(nil)

	// Processor ordering is important:
	// 1. Index configuration
	if !config.Index.IsEmpty() {
		staticFields := fmtstr.FieldsForBeat(beatInfo.Beat, beatInfo.Version)
		timestampFormat, err :=
			fmtstr.NewTimestampFormatString(&config.Index, staticFields)
		if err != nil {
			return nil, err
		}
		indexProcessor := add_formatted_index.New(timestampFormat)
		procs.List = append(procs.List, indexProcessor)
	}

	// 2. User processors
	userProcessors, err := processors.New(config.Processors)
	if err != nil {
		return nil, err
	}
	// Subtlety: it is important here that we append the individual elements of
	// userProcessors, rather than userProcessors itself, even though
	// userProcessors implements the processors.Processor interface. This is
	// because the contents of what we return are later pulled out into a
	// processing.group rather than a processors.Processors, and the two have
	// different error semantics: processors.Processors aborts processing on
	// any error, whereas processing.group only aborts on fatal errors. The
	// latter is the most common behavior, and the one we are preserving here for
	// backwards compatibility.
	// We are unhappy about this and have plans to fix this inconsistency at a
	// higher level, but for now we need to respect the existing semantics.
	procs.List = append(procs.List, userProcessors.List...)

	return procs, nil
}
