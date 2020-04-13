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

package fortigate

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "fortigate", asset.ModuleFieldsPri, AssetFortigate); err != nil {
		panic(err)
	}
}

// AssetFortigate returns asset data.
// This is the base64 encoded gzipped contents of module/fortigate.
func AssetFortigate() string {
	return "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"
}
