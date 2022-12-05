// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

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
