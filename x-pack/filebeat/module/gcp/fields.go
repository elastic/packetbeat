// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package gcp

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "gcp", asset.ModuleFieldsPri, AssetGcp); err != nil {
		panic(err)
	}
}

// AssetGcp returns asset data.
// This is the base64 encoded zlib format compressed contents of module/gcp.
func AssetGcp() string {
	return "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"
}
