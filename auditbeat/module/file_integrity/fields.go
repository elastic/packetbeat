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

package file_integrity

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "file_integrity", asset.ModuleFieldsPri, AssetFileIntegrity); err != nil {
		panic(err)
	}
}

// AssetFileIntegrity returns asset data.
// This is the base64 encoded gzipped contents of module/file_integrity.
func AssetFileIntegrity() string {
	return "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"
}
