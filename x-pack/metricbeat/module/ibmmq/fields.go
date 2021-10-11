// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package ibmmq

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "ibmmq", asset.ModuleFieldsPri, AssetIbmmq); err != nil {
		panic(err)
	}
}

// AssetIbmmq returns asset data.
// This is the base64 encoded zlib format compressed contents of module/ibmmq.
func AssetIbmmq() string {
	return "eJxUzrGugkAQheF+n+LPNjSXF9jiFnYWFNbGAmTEjbuwMkPB2xvFRJ1ycr6cU3OTNRC7nO8OLFqSQLXfNTSHysEsSVqVQCfWOuhFz3MsFqcx8O8AtjB56pckDlTM4jho4OhVk//DX82KPzm4REm9hperGdssn/Ln2VokMMzTUt6fb7Gp30mPAAAA//+eWDd6"
}
