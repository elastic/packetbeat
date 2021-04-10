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
