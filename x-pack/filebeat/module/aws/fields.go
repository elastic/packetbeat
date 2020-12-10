// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "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"
}
