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
