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
<<<<<<< HEAD
	unpacked := packer.MustUnpack("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")
=======
	unpacked := packer.MustUnpack("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")
>>>>>>> b318f6ead2 (Propagate input id from the Agent policy into Filebeat configuration (#30386))
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
