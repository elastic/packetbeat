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
