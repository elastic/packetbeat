// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package gcp

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "gcp", asset.ModuleFieldsPri, AssetGcp); err != nil {
		panic(err)
	}
}

// AssetGcp returns asset data.
// This is the base64 encoded zlib format compressed contents of module/gcp.
func AssetGcp() string {
	return "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"
}
