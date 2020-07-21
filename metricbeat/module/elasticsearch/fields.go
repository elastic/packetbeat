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
