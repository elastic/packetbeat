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

package datastore

import (
	"testing"

	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmware/govmomi/simulator"
)

func TestFetchEventContents(t *testing.T) {
<<<<<<< HEAD
	model := simulator.ESX()
	if err := model.Create(); err != nil {
		t.Fatal(err)
	}
=======
	// Creating a new simulator model with VPX server to collect broad range of data.
	model := simulator.VPX()
	err := model.Create()
	require.NoError(t, err, "failed to create model")
	t.Cleanup(func() { model.Remove() })
>>>>>>> 1058457a28 ([vSphere] Add new Cluster metricset (#40536))

	ts := model.Service.NewServer()
	t.Cleanup(func() { ts.Close() })

	f := mbtest.NewReportingMetricSetV2WithContext(t, getConfig(ts))
	events, errs := mbtest.ReportingFetchV2WithContext(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
<<<<<<< HEAD

	assert.NotEmpty(t, events)
=======
	require.Empty(t, errs, "expected no error")

	require.NotEmpty(t, events, "didn't get any event, should have gotten at least X")
>>>>>>> 1058457a28 ([vSphere] Add new Cluster metricset (#40536))

	event := events[0].MetricSetFields

	t.Logf("Fetched event from %s/%s event: %+v", f.Module().Name(), f.Name(), event)

	assert.EqualValues(t, "LocalDS_0", event["name"])
	assert.EqualValues(t, "local", event["fstype"])

	// Values are based on the result 'df -k'.
	fields := []string{"capacity.total.bytes", "capacity.free.bytes",
		"capacity.used.bytes"}
	for _, field := range fields {
		value, err := event.GetValue(field)
		if err != nil {
			t.Error(err)
		} else {
			isNonNegativeInt64(t, field, value)
		}
	}
}

<<<<<<< HEAD
func isNonNegativeInt64(t testing.TB, field string, v interface{}) {
	i, ok := v.(int64)
	if !ok {
		t.Errorf("%v: got %T, but expected int64", field, v)
		return
	}

	if i < 0 {
		t.Errorf("%v: value is negative (%v)", field, i)
		return
	}
}

func TestData(t *testing.T) {
=======
func TestDataStoreMetricSetData(t *testing.T) {
>>>>>>> 1058457a28 ([vSphere] Add new Cluster metricset (#40536))
	model := simulator.ESX()
	err := model.Create()
	require.NoError(t, err, "failed to create model")
	t.Cleanup(func() { model.Remove() })

	ts := model.Service.NewServer()
	t.Cleanup(func() { ts.Close() })

	f := mbtest.NewReportingMetricSetV2WithContext(t, getConfig(ts))

	err = mbtest.WriteEventsReporterV2WithContext(f, t, "")
	assert.NoError(t, err, "failed to write events with reporter")
}

func getConfig(ts *simulator.Server) map[string]interface{} {
	urlSimulator := ts.URL.Scheme + "://" + ts.URL.Host + ts.URL.Path

	return map[string]interface{}{
		"module":     "vsphere",
		"metricsets": []string{"datastore"},
		"hosts":      []string{urlSimulator},
		"username":   "user",
		"password":   "pass",
		"insecure":   true,
	}
}
