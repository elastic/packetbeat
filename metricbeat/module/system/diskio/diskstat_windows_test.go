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

// +build !integration
// +build windows

package diskio

import (
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCDriveFilterOnWindowsTestEnv(t *testing.T) {
	conf := map[string]interface{}{
		"module":                 "system",
		"metricsets":             []string{"diskio"},
		"diskio.include_devices": []string{"C:"},
	}

	f := mbtest.NewReportingMetricSetV2(t, conf)
	data, errs := mbtest.ReportingFetchV2(f)
	assert.Empty(t, errs)
	assert.Equal(t, 1, len(data))
}

func TestAllDrivesOnWindowsTestEnv(t *testing.T) {
	conf := map[string]interface{}{
		"module":     "system",
		"metricsets": []string{"diskio"},
	}

	f := mbtest.NewReportingMetricSetV2(t, conf)
	data, errs := mbtest.ReportingFetchV2(f)
	assert.Empty(t, errs)
	assert.True(t, len(data)>= 1)
}

