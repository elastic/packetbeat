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
