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

package auditd

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "auditd", asset.ModuleFieldsPri, AssetAuditd); err != nil {
		panic(err)
	}
}

// AssetAuditd returns asset data.
// This is the base64 encoded gzipped contents of module/auditd.
func AssetAuditd() string {
	return "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"
}
