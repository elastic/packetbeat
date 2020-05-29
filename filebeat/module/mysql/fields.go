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
	if err := asset.SetFields("filebeat", "mysql", asset.ModuleFieldsPri, AssetMysql); err != nil {
		panic(err)
	}
}

// AssetMysql returns asset data.
// This is the base64 encoded gzipped contents of module/mysql.
func AssetMysql() string {
	return "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"
}
