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
	if err := asset.SetFields("heartbeat", "Beat", asset.BeatFieldsPri, AssetBeat); err != nil {
		panic(err)
	}
}

// AssetBeat returns asset data.
// This is the base64 encoded gzipped contents of ./build/fields/beat.yml.
func AssetBeat() string {
	return "eJzMVcFu2zoQvPsrBjkn/gAdHvCQoKhvOeQDTIkraRuKVMmlW/99QZuSpUgOktZA64MtL8WZXc5w9wGvdCxQua5zdgMIi6ECd4+nAFpSXkpSgs5ZFufvNoCmUHnuhZ0tNkDNZHRIT8ADrOqoGN4+xQA59lSg8S72OTKFwH85CGTWvDsjbzd5fUo0JUvwY3Bge6XjD+f1JH6FM31eWhpJ0/aR8kKSvm9GElA5W3MTPekz8oKP9Q3Z6mgMvrkSuyeogBhIozxe1F2pV0evEu4ii6mMixxenCgz8LJtIBRkDeutllPqGGbhgdg427xZmHE/ZRawRceVd4EqZ3VY1haqllbUVIbVlLpX0haI3mwXGzpuzmQFxEf6mCT/a+0phISYU9jii/Ogn6rrDWEvVb+/x15MSD+tSPqrrD4/h/2KTK0L8olCtOsU2z8t5KsLkujhagTyB64IJSW5s/Kkt3hUFiWh4xDYNvfgy7s8F3jclDy5e16pkvtFjXzVg/Ncd8/vZrmbZjXPpM1V3s/wpCXsud+fHZzusSi24RT3FJw5kAb3UFnsOrWUllBF78nKCXWlwiBKZr739D2yJ73Q5Td6wc5qrlRqblyPfa5y0WgclGGthE45DichLimnDoqNKk3qhnlM5AIncyJZAca519h/cDRcMPCZ0TAhGufCeeXaWPhnr8at/X2xWrR6NFzDB7LX7OZleTTvNvaxuQ6+TCbJUkLZlEbtXTdemu3f6vW/AgAA//9pmX1a"
}
