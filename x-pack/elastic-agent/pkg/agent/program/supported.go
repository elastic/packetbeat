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
var SupportedMap map[string]bool

func init() {
	// Packed Files
	// spec/filebeat.yml
	// spec/metricbeat.yml
	unpacked := packer.MustUnpack("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")
	SupportedMap = make(map[string]bool)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Name)] = true
	}
}
