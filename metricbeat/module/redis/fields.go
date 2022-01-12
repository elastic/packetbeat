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

package redis

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "redis", asset.ModuleFieldsPri, AssetRedis); err != nil {
		panic(err)
	}
}

// AssetRedis returns asset data.
// This is the base64 encoded zlib format compressed contents of module/redis.
func AssetRedis() string {
	return "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"
}
