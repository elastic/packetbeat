// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import "github.com/elastic/beats/metricbeat/cmd"

// RootCmd to handle beats cli
var RootCmd = cmd.RootCmd

func init() {
	// TODO inject x-pack features
}
