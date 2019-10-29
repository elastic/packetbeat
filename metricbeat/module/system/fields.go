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

package system

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "system", asset.ModuleFieldsPri, AssetSystem); err != nil {
		panic(err)
	}
}

// AssetSystem returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/system.
func AssetSystem() string {
	return "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"
}
