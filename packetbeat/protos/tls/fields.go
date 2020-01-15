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

package tls

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "tls", asset.ModuleFieldsPri, AssetTls); err != nil {
		panic(err)
	}
}

// AssetTls returns asset data.
// This is the base64 encoded gzipped contents of protos/tls.
func AssetTls() string {
	return "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"
}
