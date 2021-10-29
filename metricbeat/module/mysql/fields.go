// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package mysql

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "mysql", asset.ModuleFieldsPri, AssetMysql); err != nil {
		panic(err)
	}
}

// AssetMysql returns asset data.
// This is the base64 encoded zlib format compressed contents of module/mysql.
func AssetMysql() string {
	return "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"
}
