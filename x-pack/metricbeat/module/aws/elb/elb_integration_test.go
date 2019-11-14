// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build integration

package elb

import (
	"fmt"
	"testing"

	"github.com/elastic/beats/libbeat/common"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
	"github.com/elastic/beats/x-pack/metricbeat/module/aws/mtest"
)

func TestData(t *testing.T) {
	namespaceIs := func(namespace string) func(e common.MapStr) bool {
		return func(e common.MapStr) bool {
			v, err := e.GetValue("aws.cloudwatch.namespace")
			return err == nil && v == namespace
		}
	}

	dataFiles := []struct {
		namespace string
		path      string
	}{
		{"AWS/ELB", "./_meta/data_elb.json"},
		{"AWS/ApplicationELB", "./_meta/data_alb.json"},
		{"AWS/NetworkELB", "./_meta/data_nlb.json"},
	}

	config, info := mtest.GetConfigForTest("elb", "300s")
	if info != "" {
		t.Skip("Skipping TestData: " + info)
	}

	for _, df := range dataFiles {
		metricSet := mbtest.NewFetcher(t, config)

		event, errs := metricSet.FetchEvents()
		if errs != nil {
			t.Skip("Skipping TestData: ", errs[0].Error())
		}

		if len(event) == 0 {
			t.Skip("Skipping TestData: FetchEvents returns no event")
		}

		eventRootFields := event[0].RootFields
		if eventRootFields == nil {
			t.Skip("Skipping TestData: event root field is empty")
		}

		t.Run(fmt.Sprintf("namespace: %s", df.namespace), func(t *testing.T) {
			metricSet.WriteEventsCond(t, df.path, namespaceIs(df.namespace))
		})
	}
}
