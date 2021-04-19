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

package decode_xml_wineventlog

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/libbeat/common/jsontransform"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/processors"
	"github.com/elastic/beats/v7/libbeat/processors/checks"
	jsprocessor "github.com/elastic/beats/v7/libbeat/processors/script/javascript/module/processor"
	"github.com/elastic/beats/v7/winlogbeat/sys/winevent"
)

var (
	errFieldIsNotString = errors.New("field value is not a string")
)

const (
	procName = "decode_xml_wineventlog"
	logName  = "processor." + procName
)

func init() {
	processors.RegisterPlugin(procName,
		checks.ConfigChecked(New,
			checks.RequireFields("field", "target_field"),
			checks.AllowedFields(
				"field", "target_field",
				"overwrite_keys", "document_id",
				"ignore_missing", "ignore_failure",
				"map_ecs_fields",
			)))
	jsprocessor.RegisterPlugin(procName, New)
}

type processor struct {
	config

	decoder decoder
	log     *logp.Logger
}

type decoder interface {
	decode(data []byte) (win, ecs common.MapStr, err error)
}

// New constructs a new decode_xml processor.
func New(c *common.Config) (processors.Processor, error) {
	config := defaultConfig()

	if err := c.Unpack(&config); err != nil {
		return nil, fmt.Errorf("fail to unpack the "+procName+" processor configuration: %s", err)
	}

	return newProcessor(config)
}

func newProcessor(config config) (processors.Processor, error) {
	cfgwarn.Experimental("The " + procName + " processor is experimental.")

	return &processor{
		config:  config,
		decoder: newDecoder(),
		log:     logp.NewLogger(logName),
	}, nil
}

func (p *processor) Run(event *beat.Event) (*beat.Event, error) {
	if err := p.run(event); err != nil && !p.IgnoreFailure {
		err = fmt.Errorf("failed in decode_xml on the %q field: %w", p.Field, err)
		event.PutValue("error.message", err.Error())
		return event, err
	}
	return event, nil
}

func (p *processor) run(event *beat.Event) error {
	data, err := event.GetValue(p.Field)
	if err != nil {
		if p.IgnoreMissing && err == common.ErrKeyNotFound {
			return nil
		}
		return err
	}

	text, ok := data.(string)
	if !ok {
		return errFieldIsNotString
	}

	win, ecs, err := p.decoder.decode([]byte(text))
	if err != nil {
		return fmt.Errorf("error decoding XML field: %w", err)
	}

	var id string
	if tmp, err := common.MapStr(win).GetValue(p.DocumentID); err == nil {
		if v, ok := tmp.(string); ok {
			id = v
			common.MapStr(win).Delete(p.DocumentID)
		}
	}

	if p.Target != "" {
		if _, err = event.PutValue(p.Target, win); err != nil {
			return fmt.Errorf("failed to put value %v into field %q: %w", win, p.Target, err)
		}
	} else {
		jsontransform.WriteJSONKeys(event, win, false, p.OverwriteKeys, !p.IgnoreFailure)
	}

	if p.MapECSFields {
		jsontransform.WriteJSONKeys(event, ecs, false, p.OverwriteKeys, !p.IgnoreFailure)
	}

	if id != "" {
		event.SetID(id)
	}

	return nil
}

func (p *processor) String() string {
	json, _ := json.Marshal(p.config)
	return procName + "=" + string(json)
}

func fields(evt winevent.Event) (common.MapStr, common.MapStr) {
	win := evt.Fields()

	ecs := common.MapStr{}

	ecs.Put("event.kind", "event")
	ecs.Put("event.code", evt.EventIdentifier.ID)
	ecs.Put("event.provider", evt.Provider.Name)
	winevent.AddOptional(ecs, "event.action", evt.Task)
	winevent.AddOptional(ecs, "host.name", evt.Computer)
	winevent.AddOptional(ecs, "event.outcome", getValue(win, "outcome"))
	winevent.AddOptional(ecs, "log.level", getValue(win, "level"))
	winevent.AddOptional(ecs, "message", getValue(win, "message"))
	winevent.AddOptional(ecs, "error.code", getValue(win, "error.code"))
	winevent.AddOptional(ecs, "error.message", getValue(win, "error.message"))

	return win, ecs
}

func getValue(m common.MapStr, key string) interface{} {
	v, _ := m.GetValue(key)
	return v
}
