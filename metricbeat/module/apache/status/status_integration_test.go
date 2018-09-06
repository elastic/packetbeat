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

package status

import (
	"testing"

	"github.com/elastic/beats/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
	"github.com/elastic/beats/metricbeat/module/apache/mtest"

	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	mtest.Runner.Run(t, compose.Suite{
		"Data": func(t *testing.T, r compose.R) {
			f := mbtest.NewEventFetcher(t, getConfig(r.Host()))

			err := mbtest.WriteEvent(f, t)
			if err != nil {
				t.Fatal("write", err)
			}
		},

		"Fetch": func(t *testing.T, r compose.R) {
			f := mbtest.NewEventFetcher(t, getConfig(r.Host()))
			event, err := f.Fetch()
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)

			// Check number of fields.
			if len(event) < 11 {
				t.Fatal("Too few top-level elements in the event")
			}
		},
	})
}

func getConfig(host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     "apache",
		"metricsets": []string{"status"},
		"hosts":      []string{host},
	}
}
