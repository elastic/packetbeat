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

package pipeline

import (
	"errors"
	"fmt"
	"sync"

	"github.com/elastic/beats/v7/libbeat/common/atomic"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/publisher"
	"github.com/elastic/beats/v7/libbeat/publisher/queue"
)

// eventConsumer collects and forwards events from the queue to the outputs work queue.
// The eventConsumer is managed by the controller and receives additional pause signals
// from the retryer in case of too many events failing to be send or if retryer
// is receiving cancelled batches from outputs to be closed on output reloading.
type eventConsumer struct {
	logger *logp.Logger
	ctx    *batchContext

	pause atomic.Bool
	wait  atomic.Bool
	sig   chan consumerSignal
	wg    sync.WaitGroup

	queue queue.Queue

	out *outputGroup
}

type consumerSignal struct {
	tag consumerEventTag
	out *outputGroup
}

type consumerEventTag uint8

const (
	sigConsumerCheck consumerEventTag = iota
	sigConsumerUpdateOutput
	sigStop
)

var errStopped = errors.New("stopped")

func newEventConsumer(
	log *logp.Logger,
	queue queue.Queue,
	ctx *batchContext,
) *eventConsumer {
	c := &eventConsumer{
		logger: log,
		sig:    make(chan consumerSignal, 3),
		out:    nil,

		queue: queue,
		ctx:   ctx,
	}

	c.pause.Store(true)

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		c.loop()
	}()
	return c
}

func (c *eventConsumer) close() {
	//c.consumer.Close()
	c.sig <- consumerSignal{tag: sigStop}
	c.wg.Wait()
}

func (c *eventConsumer) sigWait() {
	c.wait.Store(true)
	c.sigHint()
}

func (c *eventConsumer) sigUnWait() {
	c.wait.Store(false)
	c.sigHint()
}

func (c *eventConsumer) sigPause() {
	c.pause.Store(true)
	c.sigHint()
}

func (c *eventConsumer) sigContinue() {
	c.pause.Store(false)
	c.sigHint()
}

func (c *eventConsumer) sigHint() {
	// send signal to unblock a consumer trying to publish events.
	// With flags being set atomically, multiple signals can be compressed into one
	// signal -> drop if queue is not empty
	select {
	case c.sig <- consumerSignal{tag: sigConsumerCheck}:
	default:
	}
}

func (c *eventConsumer) updOutput(grp *outputGroup) {
	// update output
	c.sig <- consumerSignal{
		tag: sigConsumerUpdateOutput,
		out: grp,
	}
}

func (c *eventConsumer) loop() { //consumer queue.Consumer) {
	defer fmt.Printf("eventConsumer.loop returning GOODBYE\n")
	log := c.logger
	consumer := c.queue.Consumer()

	log.Debug("start pipeline event consumer")

	var (
		// The channel batches should be sent to; output workers are listening
		// on this channel.
		out chan publisher.Batch

		// The batch waiting to be sent, or nil if we don't yet have one
		batch Batch

		// The value of the atomic `c.paused`, the last time we read it.
		// Updated when receiving a signal on the eventConsumer's signal
		// channel `c.sig`.
		paused = true
	)

	// handleSignal always returns nil unless it receives sigStop.
	// it updates `consumer`, `c.out`, `paused`, and `out`.
	handleSignal := func(sig consumerSignal) error {
		switch sig.tag {
		case sigStop:
			fmt.Printf("got sigStop\n")
			consumer.Close()
			return errStopped

		case sigConsumerCheck:

		case sigConsumerUpdateOutput:
			consumer.Close()
			c.out = sig.out
			consumer = c.queue.Consumer()
		}

		paused = c.paused()
		if c.out != nil && batch != nil {
			out = c.out.workQueue
		} else {
			out = nil
		}
		return nil
	}

	for {
		fmt.Printf("loop iter\n")
		// If we want a batch but don't yet have one
		if !paused && c.out != nil && consumer != nil && batch == nil {
			out = c.out.workQueue
			queueBatch, err := consumer.Get(c.out.batchSize)
			if err != nil {
				out = nil
				consumer = nil
				continue
			}
			if queueBatch != nil {
				batch = newBatch(c.ctx, queueBatch, c.out.timeToLive)
			}

			paused = c.paused()
			if paused || batch == nil {
				out = nil
			}
		}

		select {
		case sig := <-c.sig:
			if err := handleSignal(sig); err != nil {
				return
			}
			continue
		default:
		}

		select {
		case sig := <-c.sig:
			if err := handleSignal(sig); err != nil {
				return
			}
		case out <- batch:
			batch = nil
			if paused {
				out = nil
			}
		}
	}
}

func (c *eventConsumer) paused() bool {
	return c.pause.Load() || c.wait.Load()
}
