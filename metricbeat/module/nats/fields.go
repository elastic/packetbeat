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

package nats

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "nats", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzUlz9z2zwMxnd/Clz2vB/Aw3vXa5cMzdCks0NTiMxGIhQQjE/99D3QkSIrlP+ksc/l4EEi8fwAiA/pa3jCdg7eSJgBiJMK53B1ayRczQAKDJZdI478HP6fAUCaCd+piBXOAB4dVkWYpzfX4E2NfSwd0jY4h5IpNq9PMhF1POiiB7DkxTgfIIgRF8TZALIyAmtkBEZTwCNTDbdvEkOCIYVGCP3THMoOHB1fO5aU8YthYKyMYAE1Cjs7DD6G2AJBfkFeuGLrbQf0hO2aePxuB5aO+xW+hoWbb1lRT+usXGEEj9O62+iIqxHo8TV3sIxGV2TFY6Ozs/oV+fL4XBtkRwXIW9ouwFYf38RrrLPKzguWyMeL28iMXjQwcQsxmDJV4vbL/R00TBZDyJJYYgyfy+JjvURW9YpKZ021EUmVGfIARx9goj22iaepkG3iVnnCzvIIiakWlrxHq9FPVqokVLWbbxYLsJVDL3koxprkdF1jLF0QZCyySh2F81mAsXkdKG9qil5U3nlLtfOl2oAZzc1ZGGztq6CtHZdmX3kOYNzBOSnaUS3b9906LdJ7xY6FopymbRSlpH+gbT3n5bStR5puW6horTYUYo18sp2vKtCrdJZZtUA+OXcWbSXSLBifF+ky8x8TySKyyzLS8hfa8Qe4r17bjCsnemiA6kBN3gmx1m4sOYGnVv77U/nyeEnnA3xMUfAcgBuhDxCGuAznAEw6H+HT33Pwqc6hfJ/puRqvu973d7zNlfNI3+2LNulyOcQDMDtUxueIQYZ/lTJzpzCHqO/7eRjogbAjYChcEHbLqGu02YM2//xxczcRY1ciw2Q2e29y2v4j5sjM4C9cIMefzO0S8Pe6bI7+xfBFwCvHsezJFC8Bfq875+j1zL4E+Nzd4U8AAAD//1Pdpz4="
}
