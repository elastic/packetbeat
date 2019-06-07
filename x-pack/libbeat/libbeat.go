// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package main

import (
	"os"

	"github.com/elastic/beats/libbeat/cmd"
	"github.com/elastic/beats/libbeat/mock"
	_ "github.com/elastic/beats/x-pack/libbeat/autodiscover/providers/aws/elb"
	xpackcmd "github.com/elastic/beats/x-pack/libbeat/cmd"
)

// RootCmd to test libbeat
var RootCmd = cmd.GenRootCmdWithSettings(mock.New, mock.Settings)

func main() {
	xpackcmd.AddXPack(RootCmd, mock.Name)
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
