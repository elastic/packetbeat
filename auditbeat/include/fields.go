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

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
