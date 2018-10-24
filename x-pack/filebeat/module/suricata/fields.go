// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package suricata

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "suricata", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
