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
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", asset.ModuleFieldsPri, AssetDocker); err != nil {
		panic(err)
	}
}

// AssetDocker returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/docker.
func AssetDocker() string {
	return "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"
}
