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
