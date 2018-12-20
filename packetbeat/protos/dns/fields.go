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

package dns

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "dns", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
