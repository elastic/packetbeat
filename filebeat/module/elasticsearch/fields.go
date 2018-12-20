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
	if err := asset.SetFields("filebeat", "elasticsearch", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
