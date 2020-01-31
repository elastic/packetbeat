// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "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"
}
