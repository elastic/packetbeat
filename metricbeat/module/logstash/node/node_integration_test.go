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

// +build integration

package node

import (
	"testing"

	"github.com/elastic/beats/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
	"github.com/elastic/beats/metricbeat/module/logstash"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	logstash.Runner.Run(t, compose.Suite{
		"Fetch": func(t *testing.T, r compose.R) {
			f := mbtest.NewEventFetcher(t, logstash.GetConfig("node", r.Host()))
			event, err := f.Fetch()
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			assert.NotNil(t, event)
			t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)
		},
		"Data": func(t *testing.T, r compose.R) {
			f := mbtest.NewEventFetcher(t, logstash.GetConfig("node", r.Host()))
			err := mbtest.WriteEvent(f, t)
			if err != nil {
				t.Fatal("write", err)
			}

		},
	})
}
