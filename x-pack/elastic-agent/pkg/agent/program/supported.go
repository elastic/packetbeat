// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/elastic-agent/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// spec/apm-server.yml
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/fleet-server.yml
	// spec/heartbeat.yml
	// spec/metricbeat.yml
	// spec/osquerybeat.yml
	// spec/packetbeat.yml
	unpacked := packer.MustUnpack("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")
	SupportedMap = make(map[string]Spec)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Cmd)] = s
	}
}
