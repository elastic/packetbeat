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
	"sync"

	"github.com/elastic/beats/v7/libbeat/common/atomic"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/reload"
	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/publisher/queue"
)

var _outputGroupID atomic.Uint

// outputController manages the pipelines output capabilities, like:
// - start
// - stop
// - reload
type outputController struct {
	beat     beat.Info
	monitors Monitors
	observer outputObserver

	queue queue.Queue

	retryer  *retryer
	consumer *eventConsumer
	out      *outputGroup

	setInProgress sync.Mutex
}

// outputGroup configures a group of load balanced outputs with shared work queue.
type outputGroup struct {
	id        uint
	workQueue workQueue
	outputs   []outputWorker

	batchSize  int
	timeToLive int // event lifetime
}

type workQueue chan *Batch

// outputWorker instances pass events from the shared workQueue to the outputs.Client
// instances.
type outputWorker interface {
	Close() error
}

func newOutputController(
	beat beat.Info,
	monitors Monitors,
	observer outputObserver,
	b queue.Queue,
) *outputController {
	c := &outputController{
		beat:     beat,
		monitors: monitors,
		observer: observer,
		queue:    b,
	}

	ctx := &batchContext{}
	c.consumer = newEventConsumer(monitors.Logger, b, ctx)
	c.retryer = newRetryer(monitors.Logger, observer, nil, c.consumer)
	ctx.observer = observer
	ctx.retryer = c.retryer

	ln("starting consumer...")
	c.consumer.sigContinue()
	return c
}

func (c *outputController) Close() error {
	c.consumer.sigPause()
	c.consumer.close()
	c.retryer.close()

	if c.out != nil {
		for _, out := range c.out.outputs {
			out.Close()
		}
		close(c.out.workQueue)
	}

	return nil
}

func (c *outputController) Set(outGrp outputs.Group) {
	ln("set called...")
	c.setInProgress.Lock()
	defer c.setInProgress.Unlock()
	ln("setting output...")

	// update consumer and retryer
	// close old group, so events are send to new workQueue via retryer
	ln("pausing consumer...")
	c.consumer.sigPause()
	//if c.out != nil {
	//	for range c.out.outputs {
	//		c.retryer.sigOutputRemoved()
	//	}
	//}

	// create new outputGroup with shared work queue
	clients := outGrp.Clients
	queue := makeWorkQueue()
	worker := make([]outputWorker, len(clients))
	for i, client := range clients {
		worker[i] = makeClientWorker(c.observer, queue, client)
	}
	grp := &outputGroup{
		id:         _outputGroupID.Inc(),
		workQueue:  queue,
		outputs:    worker,
		timeToLive: outGrp.Retry + 1,
		batchSize:  outGrp.BatchSize,
	}

	// close old group, so events are send to new workQueue via retryer
	if c.out != nil {
		for _, w := range c.out.outputs {
			w.Close()
		}
		for drained := false; !drained; {
			select {
			case b := <-c.out.workQueue:
				ln("draining batches from old workqueue to new workqueue...")
				queue <- b
			default:
				ln("workqueue is already drained")
				drained = true
			}
		}
	}

	// update consumer and retryer
	//ln("pausing consumer...")
	//c.consumer.sigPause()
	if c.out != nil {
		for range c.out.outputs {
			c.retryer.sigOutputRemoved()
		}
	}

	c.retryer.updOutput(queue)
	for range clients {
		c.retryer.sigOutputAdded()
	}
	c.consumer.updOutput(grp)

	c.out = grp

	// restart consumer (potentially blocked by retryer)
	ln("continuing consumer...")
	c.consumer.sigContinue()

	c.observer.updateOutputGroup()
}

func makeWorkQueue() workQueue {
	return workQueue(make(chan *Batch, 0))
}

// Reload the output
func (c *outputController) Reload(
	cfg *reload.ConfigWithMeta,
	outFactory func(outputs.Observer, common.ConfigNamespace) (outputs.Group, error),
) error {
	outCfg := common.ConfigNamespace{}
	if cfg != nil {
		if err := cfg.Config.Unpack(&outCfg); err != nil {
			return err
		}
	}

	output, err := loadOutput(c.monitors, func(stats outputs.Observer) (string, outputs.Group, error) {
		name := outCfg.Name()
		out, err := outFactory(stats, outCfg)
		return name, out, err
	})
	if err != nil {
		return err
	}

	c.Set(output)

	return nil
}
