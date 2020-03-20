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

package beater

import (
	"fmt"
	"sync"

	"github.com/elastic/beats/v7/filebeat/channel"
	"github.com/elastic/beats/v7/filebeat/input"
	"github.com/elastic/beats/v7/filebeat/input/file"
	"github.com/elastic/beats/v7/filebeat/registrar"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/cfgfile"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
)

type crawler struct {
	inputs          map[uint64]*input.Runner
	inputConfigs    []*common.Config
	out             channel.Factory
	wg              sync.WaitGroup
	inputsFactory   cfgfile.RunnerFactory
	modulesFactory  cfgfile.RunnerFactory
	modulesReloader *cfgfile.Reloader
	inputReloader   *cfgfile.Reloader
	once            bool
	beatDone        chan struct{}
}

func newCrawler(
	inputFactory, module cfgfile.RunnerFactory,
	out channel.Factory,
	inputConfigs []*common.Config,
	beatDone chan struct{},
	once bool,
) (*crawler, error) {
	return &crawler{
		out:            out,
		inputs:         map[uint64]*input.Runner{},
		inputsFactory:  inputFactory,
		modulesFactory: module,
		inputConfigs:   inputConfigs,
		once:           once,
		beatDone:       beatDone,
	}, nil
}

// Start starts the crawler with all inputs
func (c *crawler) Start(
	pipeline beat.Pipeline,
	r *registrar.Registrar,
	configInputs *common.Config,
	configModules *common.Config,
) error {

	logp.Info("Loading Inputs: %v", len(c.inputConfigs))

	// Prospect the globs/paths given on the command line and launch harvesters
	for _, inputConfig := range c.inputConfigs {
		err := c.startInput(pipeline, inputConfig, r.GetStates())
		if err != nil {
			return err
		}
	}

	if configInputs.Enabled() {
		c.inputReloader = cfgfile.NewReloader(pipeline, configInputs)
		if err := c.inputReloader.Check(c.inputsFactory); err != nil {
			return err
		}

		go func() {
			c.inputReloader.Run(c.inputsFactory)
		}()
	}

	if configModules.Enabled() {
		c.modulesReloader = cfgfile.NewReloader(pipeline, configModules)
		if err := c.modulesReloader.Check(c.modulesFactory); err != nil {
			return err
		}

		go func() {
			c.modulesReloader.Run(c.modulesFactory)
		}()
	}

	logp.Info("Loading and starting Inputs completed. Enabled inputs: %v", len(c.inputs))

	return nil
}

func (c *crawler) startInput(
	pipeline beat.Pipeline,
	config *common.Config,
	states []file.State,
) error {
	if !config.Enabled() {
		return nil
	}

	connector := c.out(pipeline)
	p, err := input.New(config, connector, c.beatDone, states, nil)
	if err != nil {
		return fmt.Errorf("Error while initializing input: %s", err)
	}
	p.Once = c.once

	if _, ok := c.inputs[p.ID]; ok {
		return fmt.Errorf("Input with same ID already exists: %d", p.ID)
	}

	c.inputs[p.ID] = p

	p.Start()

	return nil
}

func (c *crawler) Stop() {
	logp.Info("Stopping Crawler")

	asyncWaitStop := func(stop func()) {
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			stop()
		}()
	}

	logp.Info("Stopping %v inputs", len(c.inputs))
	for _, p := range c.inputs {
		// Stop inputs in parallel
		asyncWaitStop(p.Stop)
	}

	if c.inputReloader != nil {
		asyncWaitStop(c.inputReloader.Stop)
	}

	if c.modulesReloader != nil {
		asyncWaitStop(c.modulesReloader.Stop)
	}

	c.WaitForCompletion()

	logp.Info("Crawler stopped")
}

func (c *crawler) WaitForCompletion() {
	c.wg.Wait()
}
