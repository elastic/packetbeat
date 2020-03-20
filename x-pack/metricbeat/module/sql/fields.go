// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package sql

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "sql", asset.ModuleFieldsPri, AssetSql); err != nil {
		panic(err)
	}
}

// AssetSql returns asset data.
// This is the base64 encoded gzipped contents of module/sql.
func AssetSql() string {
	return "eJykkkFuwjAQRfc5xRfLSnAAL7rqskJCPUDl2B9wcWJij2lz+4okRqlSpKJ6+T0z72XiNU7sFVLnK0CceCqsUudXFRDpqRMVaoquAMtkojuLC63CcwUAb7tXNMFmT+wp5siEhhKdSdjH0EAPFVaLrnViBewdvU1qaF6j1Q0L/HqkP1PhEEM+T8m8ft5jo7sw3uLSemL/GaKd5b9Il/MyzEBOtJAAftFkIeRIdJmx3yyoQ/w/6O46orAGrgne00hZ3JJaLtrcMDqzeVoYhPqDRmbxGLyPtzbk2vNvetuRcfuLkxztfa0k0bWHh60e2to2tOvp83HRPvOO2c8n+x0AAP//MiXQYA=="
}
