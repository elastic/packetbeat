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

//go:build integration

package status

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/v7/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func TestFetch(t *testing.T) {
	service := compose.EnsureUp(t, "apache")

	f := mbtest.NewReportingMetricSetV2Error(t, getConfig(service.Host()))
	events, errs := mbtest.ReportingFetchV2Error(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)
	event := events[0]

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)

	// Check number of fields.
	if len(event.MetricSetFields) < 11 {
		t.Fatal("Too few top-level elements in the event")
	}
}

func TestFetchFleetMode(t *testing.T) {
	service := compose.EnsureUp(t, "apache")

	f := mbtest.NewReportingMetricSetV2Error(t, getConfig(service.Host()))
	f.(*MetricSet).isFleetMode = true // silently simulate running in the fleet mode

	events, errs := mbtest.ReportingFetchV2Error(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)
	event := events[0]

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)

	// Check number of fields.
	if len(event.MetricSetFields) < 11 {
		t.Fatal("Too few top-level elements in the event")
	}

	_, err := event.MetricSetFields.GetValue("total_kbytes")
	assert.Equal(t, mapstr.ErrKeyNotFound, err, "apache.status.total_kbytes shouldn't be present in the fleet mode")

	totalBytes, err := event.MetricSetFields.GetValue("total_bytes")
	assert.NoError(t, err, "apache.status.total_bytes should be present in the fleet mode")
	assert.GreaterOrEqual(t, totalBytes.(int64), int64(0), "apache.status.total_bytes should be non-negative")

	_, err = event.MetricSetFields.GetValue("hostname")
	assert.Equal(t, mapstr.ErrKeyNotFound, err, "apache.status.hostname shouldn't be present in the fleet mode")
}

func getConfig(host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     "apache",
		"metricsets": []string{"status"},
		"hosts":      []string{host},
	}
}
