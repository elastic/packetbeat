// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package suricata

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "suricata", asset.ModuleFieldsPri, AssetSuricata); err != nil {
		panic(err)
	}
}

// AssetSuricata returns asset data.
// This is the base64 encoded gzipped contents of module/suricata.
func AssetSuricata() string {
	return "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"
}
