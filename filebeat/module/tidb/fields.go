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

package tidb

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "tidb", asset.ModuleFieldsPri, AssetTidb); err != nil {
		panic(err)
	}
}

// AssetTidb returns asset data.
// This is the base64 encoded zlib format compressed contents of module/tidb.
func AssetTidb() string {
	return "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"
}
