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

package collstats

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
	"github.com/elastic/beats/metricbeat/module/mongodb/mtest"
)

func TestCollstats(t *testing.T) {
	mtest.Runner.Run(t, compose.Suite{
		"Fetch": testFetch,
		"Data":  testData,
	})
}

func testFetch(t *testing.T, r compose.R) {
	f := mbtest.NewEventsFetcher(t, mtest.GetConfig("collstats", r.Host()))
	events, err := f.Fetch()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	for _, event := range events {
		t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)

		// Check a few event Fields
		db := event["db"].(string)
		assert.NotEqual(t, db, "")

		collection := event["collection"].(string)
		assert.NotEqual(t, collection, "")
	}
}

func testData(t *testing.T, r compose.R) {
	f := mbtest.NewEventsFetcher(t, mtest.GetConfig("collstats", r.Host()))
	err := mbtest.WriteEvents(f, t)
	if err != nil {
		t.Fatal("write", err)
	}
}
