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
