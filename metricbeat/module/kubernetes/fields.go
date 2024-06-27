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
