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

//go:build !integration

package module_test

import (
	stdjson "encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/outputs/codec/json"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/mb/module"
	conf "github.com/elastic/elastic-agent-libs/config"
)

// ExampleWrapper demonstrates how to create a single Wrapper
// from configuration, start the module, and consume events generated by the
// module.
func ExampleWrapper() {
	// Build a configuration object.
	config, err := conf.NewConfigFrom(map[string]interface{}{
		"module":     moduleName,
		"metricsets": []string{reportingFetcherName},
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a new Wrapper based on the configuration.
	m, err := module.NewWrapper(config, mb.Registry, module.WithMetricSetInfo())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Run the module until done is closed.
	done := make(chan struct{})
	output := m.Start(done)

	// Process events from the output channel until it is closed.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range output {
			event.Fields.Put("event.duration", 111)

			output, err := encodeEvent(event)
			if err == nil {
				fmt.Println(output)
			}
		}
	}()

	// Simulate running for a while.
	time.Sleep(50 * time.Millisecond)

	// When finished with the module, close the done channel. When the Module
	// stops it will automatically close its output channel so that the output
	// for loop stops.
	close(done)
	wg.Wait()

	// Output:
	// {
	//   "@metadata": {
	//     "beat": "noindex",
	//     "type": "_doc",
	//     "version": "1.2.3"
	//   },
	//   "@timestamp": "2016-05-10T23:27:58.485Z",
	//   "event": {
	//     "dataset": "fake.reportingfetcher",
	//     "duration": 111,
	//     "module": "fake"
	//   },
	//   "fake": {
	//     "reportingfetcher": {
	//       "metric": 1
	//     }
	//   },
	//   "metricset": {
	//     "name": "reportingfetcher",
	//     "period": 10000
	//   },
	//   "service": {
	//     "type": "fake"
	//   }
	// }
}

// ExampleRunner demonstrates how to use Runner to start and stop
// a module.
func ExampleRunner() {
	// A *beat.Beat is injected into a Beater when it runs and contains the
	// Publisher used to publish events. This Beat pointer is created here only
	// for demonstration purposes.
	var b *beat.Beat

	config, err := conf.NewConfigFrom(map[string]interface{}{
		"module":     moduleName,
		"metricsets": []string{reportingFetcherName},
	})
	if err != nil {
		return
	}

	// Create a new Wrapper based on the configuration.
	m, err := module.NewWrapper(config, mb.Registry, module.WithMetricSetInfo())
	if err != nil {
		return
	}

	connector, err := module.NewConnector(b.Info, b.Publisher, config)
	if err != nil {
		return
	}

	client, err := connector.Connect()
	if err != nil {
		return
	}

	// Create the Runner facade.
	runner := module.NewRunner(client, m)

	// Start the module and have it publish to a new publisher.Client.
	runner.Start()

	// Stop the module. This blocks until all MetricSets in the Module have
	// stopped and the publisher.Client is closed.
	runner.Stop()
}

func encodeEvent(event beat.Event) (string, error) {
	output, err := json.New("1.2.3", json.Config{}).Encode("noindex", &event)
	if err != nil {
		return "", nil
	}

	// FIX: need to parse and re-encode, so fields ordering in final json document
	//      keeps stable.

	var tmp interface{}
	if err := stdjson.Unmarshal(output, &tmp); err != nil {
		panic(err)
	}

	output, err = stdjson.MarshalIndent(tmp, "", "  ")
	return string(output), err
}
