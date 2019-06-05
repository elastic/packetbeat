// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build integration

package tablespace

import (
	"github.com/elastic/beats/x-pack/metricbeat/module/oracle"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "gopkg.in/goracle.v2"

	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

func TestData(t *testing.T) {
	f := mbtest.NewReportingMetricSetV2Error(t, getConfig())

	events, errs := mbtest.ReportingFetchV2Error(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)

	if err := mbtest.WriteEventsReporterV2Error(f, t, ""); err != nil {
		t.Fatal("write", err)
	}
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "oracle",
		"metricsets": []string{"tablespace"},
		"hosts":      []string{oracle.GetOracleConnectionDetails()},
		"username":   "sys",
		"password":   "Oradoc_db1",
	}
}
