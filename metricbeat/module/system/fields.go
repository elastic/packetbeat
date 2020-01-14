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
