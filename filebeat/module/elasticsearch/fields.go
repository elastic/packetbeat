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
	"github.com/elastic/beats/libbeat/asset"
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
