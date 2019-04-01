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
