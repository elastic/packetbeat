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

// +build integration,linux

package mgr_cluster_disk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

func TestData(t *testing.T) {
	service := compose.EnsureUpWithTimeout(t, 120, "ceph-mgr")

	f := mbtest.NewReportingMetricSetV2Error(t,
		getConfig(service.HostForPort(8003), getPassword(t, service.HostForPort(5000))))
	err := mbtest.WriteEventsReporterV2Error(f, t, "")
	require.NoError(t, err)
}

func getPassword(t *testing.T, host string) string {
	response, err := http.Get(fmt.Sprintf("http://%s/restful-list-keys.json", host))
	require.NoError(t, err)

	defer response.Body.Close()

	accounts := map[string]string{}
	json.NewDecoder(response.Body).Decode(&accounts)
	return accounts["demo"]
}

func getConfig(host, password string) map[string]interface{} {
	return map[string]interface{}{
		"module":                "ceph",
		"metricsets":            []string{"mgr_cluster_disk"},
		"hosts":                 []string{host},
		"username":              "demo",
		"password":              password,
		"ssl.verification_mode": "none",
	}
}
