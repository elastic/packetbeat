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
