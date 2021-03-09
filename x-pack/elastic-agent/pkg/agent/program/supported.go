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
