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
	if err := asset.SetFields("journalbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
