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
