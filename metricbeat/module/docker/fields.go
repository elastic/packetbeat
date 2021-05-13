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

package docker

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", asset.ModuleFieldsPri, AssetDocker); err != nil {
		panic(err)
	}
}

// AssetDocker returns asset data.
// This is the base64 encoded gzipped contents of module/docker.
func AssetDocker() string {
	return "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"
}
