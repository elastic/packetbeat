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
	if err := asset.SetFields("filebeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded gzipped contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
