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

package mongodb

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "mongodb", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
