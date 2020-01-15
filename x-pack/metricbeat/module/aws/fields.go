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
