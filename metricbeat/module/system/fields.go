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
