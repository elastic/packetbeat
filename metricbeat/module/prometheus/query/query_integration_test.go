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

package query

import (
	"testing"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"

	"github.com/stretchr/testify/assert"
)

func TestData(t *testing.T) {
	service := compose.EnsureUp(t, "prometheus")

	config := map[string]interface{}{
		"module":     "prometheus",
		"metricsets": []string{"query"},
		"hosts":      []string{service.Host()},
		"queries": []common.MapStr{
			common.MapStr{
				"query_name": "go_threads",
				"path":       "/api/v1/query",
				"query_params": common.MapStr{
					"query": "go_threads",
				},
			},
		},
	}
	ms := mbtest.NewReportingMetricSetV2Error(t, config)
	err := mbtest.WriteEventsReporterV2Error(ms, t, "")
	if err == nil {
		return
	}
	t.Fatal("write", err)
}

func TestQueryFetch(t *testing.T) {
	service := compose.EnsureUp(t, "prometheus")

	config := map[string]interface{}{
		"module":     "prometheus",
		"metricsets": []string{"query"},
		"hosts":      []string{service.Host()},
		"queries": []common.MapStr{
			common.MapStr{
				"query_name": "go_threads",
				"path":       "/api/v1/query",
				"query_params": common.MapStr{
					"query": "go_threads",
				},
			},
		},
	}
	f := mbtest.NewReportingMetricSetV2Error(t, config)
	events, errs := mbtest.ReportingFetchV2Error(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 errors, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)
	event := events[0].MetricSetFields
	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)
}
