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

package pensando

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "pensando", asset.ModuleFieldsPri, AssetPensando); err != nil {
		panic(err)
	}
}

// AssetPensando returns asset data.
// This is the base64 encoded gzipped contents of module/pensando.
func AssetPensando() string {
	return "eJy0k8GK2zAQhu9+inmBzQP4sFAICz0UQrfQo5la40Ss7BHSGG/evkhRbMWxySZNdZRG///NL80LfNCxBEudx05xASBaDJWwSztQACjytdNWNHclvBYAMF6AH6x6QwVAo8koX8bTF+iwpQvZsORoqYS9496mnQXlsN6iFjSO2wnE8N5vUkludmGIR8OoKofDeDZz2RlCT4BK5ftZNX1ia2MEU2Xay6pOvQh9yhWEaga4qsy7DstRVC/hDwlesDbYG6lihyU0aDyttfKaHYyZsZsi2779zkrmoeXMWM9SmMg/6DiwU7OzdZC4vkU94A7kQNAYHjawbGxtpefiJ2PdCe3JwZ3O1hpdY7T/vl02VeRFd7GmQqUceT93uej9XoQkyU3utBJAzmLZyTLIY1nsgt7XKLSj05M97w9sR01uxm+wbG8dC9dsnui+O0veNHe9oUpfPfK//ICfvaHw++SAAgN6aFHqA6mV9D15H95/DeKx138/qQaOmxmcCbyg0BODODNE3RxjLQnuXU3/YyjfozLgNJtfQ3n2TCYOm0bzBoTolrxga5cRVEj1Pv9fo2JyN7zfQPE3AAD//3eQEuo="
}
