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

package thrift

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "thrift", asset.ModuleFieldsPri, AssetThrift); err != nil {
		panic(err)
	}
}

// AssetThrift returns asset data.
// This is the base64 encoded gzipped contents of ../../packetbeat/protos/thrift.
func AssetThrift() string {
	return "eJyUUrGum0AQ7PmK0atjPoAiTV5jKVKiyH205obHKscduj148d9HHCTGzy6cq2A1Mzu7swf84qVB7pN2uQKyZs8GL6dSOPz4/uWlAhytTTpmjaHB5woArgDYyFY7bcGZIaNTemd1he2rKfgDggzcdVpevoxs8JbiNG6VPWVPGyXJYP/KDy39faeeWHwNzH10aMX7lc/MZNAAQT8NEpAoTs6e6GIaJNc4dsg9cXz9eqPYqadBEiGzqF8onwpwJzsZi1fDe8/AmQljNNOzZ30j9i33TO9qm8Tx1dClOJSfgWbyxtJqMrq6utuEMc3a8vlVLDTEdbB9aqsOxODYaaBbNrNNv078oHtinlL4OYuf/sNCgW9cOpwvH80sGdU49WpQA0Mb3epnS+pG8WNq9y75u2Ux8+TFbLGXS0m0yee1+1VnyUptPU+0MWTRYIV0xTzv908AAAD//0KDB+4="
}
