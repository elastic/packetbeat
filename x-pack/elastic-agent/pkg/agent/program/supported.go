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
