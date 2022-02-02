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
// This is the base64 encoded zlib format compressed contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
