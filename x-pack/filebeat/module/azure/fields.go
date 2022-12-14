// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package azure

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded zlib format compressed contents of module/azure.
func AssetAzure() string {
	return "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"
}
