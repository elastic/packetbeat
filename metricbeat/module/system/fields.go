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
