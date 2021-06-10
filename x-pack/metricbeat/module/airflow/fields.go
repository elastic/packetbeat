// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package airflow

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "airflow", asset.ModuleFieldsPri, AssetAirflow); err != nil {
		panic(err)
	}
}

// AssetAirflow returns asset data.
// This is the base64 encoded gzipped contents of module/airflow.
func AssetAirflow() string {
	return "eJyUkU1qBCEQhfee4uFmoGFyABeBnGRw7OrGRC3RasLcPvQvEnoWXcv3vsKv8I4fehlYX4bAvwoQL4EM9NeaaAX0VF3xWTwng08FAFuLyP0USAGFAtlKBqNVwOAp9NUs6B3JRmqfmEdeeYYLT3lL2p1279Z9OJ6S3I5m3+bnNzlp4jV4rG3gNJ53j2hz9mncQD2TukFP7t1nv3sxolLPbLurpkNg+6b8r9pd9IwkxbtW8/ioJ4lVfwEAAP//RcaOgA=="
}
