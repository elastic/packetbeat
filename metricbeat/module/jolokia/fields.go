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
	if err := asset.SetFields("metricbeat", "jolokia", asset.ModuleFieldsPri, AssetJolokia); err != nil {
		panic(err)
	}
}

// AssetJolokia returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/jolokia.
func AssetJolokia() string {
	return "eJyEkVFOwzAQRP9zipG/aQ/gDw7AFRBCJtmk29reyLstye1RmgS5AsR+jtczTzsHXGj2OEuUC4cGMLZIHu5lVVwDdKRt4dFYssdzAwDbK5J010gNoCcp9t5K7nnw6EPURS0UKSh5DIu1khnnQT1enWp0T3Ans9G9NUDPFDv1d/MDckhUQy1j87gYFbmOm/IL1zrbR7SSLXBWJLLCrYKmUZQ63Dh8L4WBsm2fa4qa5L5zvFFRllwFrUwXmj+ldJX+JxmwmSBf0wcVSP8Icqyy12bS9COwPsI/cXtR5zTtVyCrNh4b2me/w1cAAAD//5henSM="
}
