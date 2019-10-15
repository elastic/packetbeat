// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package netflow

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "netflow", asset.ModuleFieldsPri, AssetNetflow); err != nil {
		panic(err)
	}
}

// AssetNetflow returns asset data.
// This is the base64 encoded gzipped contents of input/netflow.
func AssetNetflow() string {
	return "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"
}
