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

package elasticsearch

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded gzipped contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
