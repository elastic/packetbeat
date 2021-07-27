// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package netflow

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "netflow", asset.ModuleFieldsPri, AssetNetflow); err != nil {
		panic(err)
	}
}

// AssetNetflow returns asset data.
// This is the base64 encoded zlib format compressed contents of input/netflow.
func AssetNetflow() string {
	return "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"
}
