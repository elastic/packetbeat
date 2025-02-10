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
