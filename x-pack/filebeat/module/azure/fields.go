// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package azure

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded gzipped contents of module/azure.
func AssetAzure() string {
	return "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"
}
