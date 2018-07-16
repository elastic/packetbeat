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

package jolokia

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "jolokia", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJxsjkFuBCEMBO+8osV9P8AhD8gXoihC4CHOMhhhb5T5fUQGRXPYPrbdpbrhTkfAl1S5c3SAsVUK8K9n4x2QSdPgbiwt4MUBwLpil/yo5AD9lGEfSdrGJWCLVWc7qFJUCigTrWTGrWjAm1et/t0BG1PNGv6gN7S401Vmxo4+AUMefTVPfM6sIZI0i9wUO9ngpKCfLkoZ3xz/n2KhZmt8tThNrua/AQAA//9JPFm3"
}
