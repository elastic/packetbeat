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
