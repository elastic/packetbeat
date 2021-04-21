// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package threatintel

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "threatintel", asset.ModuleFieldsPri, AssetThreatintel); err != nil {
		panic(err)
	}
}

// AssetThreatintel returns asset data.
// This is the base64 encoded gzipped contents of module/threatintel.
func AssetThreatintel() string {
	return "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"
}
