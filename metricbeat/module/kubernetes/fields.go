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

package kubernetes

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded zlib format compressed contents of module/kubernetes.
func AssetKubernetes() string {
	return "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"
}
