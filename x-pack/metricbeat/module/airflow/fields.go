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
	return "eJyUkc1qxSAQhfc+xcHNhcDtA7go9EkuXjMJtuqITih5+5K/YkvuIrM83xn8Bu/4otnA+jIE/laAeAlkoD+2RCugp+qKz+I5GbwrANgpIvdTIAUUCmQrGTxJrAIGT6GvZi3fkWyk9pFlZM5kMBae8p60O+3erXtzPCW5/ZJjm5+f5KSJt+Cx0cBpPGePaHP2adyLemnqpnpy8THH5asRlXpm2101HQLbF/C/anfRM5IU71rNv1/1EwAA///euo9U"
}
