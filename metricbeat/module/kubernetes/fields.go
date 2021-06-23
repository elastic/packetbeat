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
// This is the base64 encoded gzipped contents of module/kubernetes.
func AssetKubernetes() string {
	return "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"
}
