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

package uwsgi

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "uwsgi", asset.ModuleFieldsPri, AssetUwsgi); err != nil {
		panic(err)
	}
}

// AssetUwsgi returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/uwsgi.
func AssetUwsgi() string {
	return "eJzEl8tu6zYQhvd+ikFWLWAb7daLAgGSRRdtUzvnslNocWQTpkhlOLQiP/0BdfFVVowDGoc7yfT/f8PLzGgCG6xm4Eu3UiMAVqxxBg/188MIQKJLSRWsrJnBXyMAaOZCbqXXOAIg1CgczmAlRgCZQi3drJ44ASNyPIiHwVURppL1Rfumx+FU5ljKsWDv9q/79MI4ZepGr1MzasRpow45MqnUIbcYRzPPuY7Z2LLQU8J3j47dyZSOU1uzOvthACqM1yAKnSishZEa5YA/fqRYq8UluCJ76l2SYkyQyNKd4q8doMfhfBOEvA+HkJ+7F0pGcn0hm6JzoPp3vLS0QZpGs/tW633iFi+61q64KUiR1gfwwuNnzRcsGMFmrf4Y/gSVgWOlNey9wGB5OHyW10ilcgh/DJFGzgD/+nyJFEj3IA5pixKWFfBaudZ3CEmiZpH8IjAQGSN1D8qF7GyFRAnlGk24U+kaJfzz+D2ZP///5XnxuhiKJXqCe94LAgnlrmTX1n0tSGwUqSS13nAkgieyRYHysI5hAVWO1vMQi1MrI3S0ZcgVM0poVeEywF775N2jx+hHqqMohaovIltY4mD566jOO4QDzwar0tJ59rotSbWtwW/pGkUxhkJ4h+NAOYald9UYlNT4+2BacBGp5uiURMOwQIaF2uEUcswtVZB6IjSsK/CuuYqiy7DTkOGELkXlYIdkgamCt8mk+eeEsLDEb2Br17AJx41bb0xbt4u0818VsRf6Mh6ndgjChfOAMhyDO4ZD3hhlVkm4e5EreCsNF9IXxcMVojRR08u80fz8RvNHrG6JhHH7fLIbjFhsVwnFCvVxiyRWCGEZrXF4fb1TSxivc2qL299P172aKUnE5lCEi95Zvwx5d4VlWjeo0ZM1nzbpTRdwA01IqiqNXztqWciUxoalv1k/ZSHrGWNtzbwWg96Gq9/eZlnTFEUi+K/TuwHiDp9th83IhNKh1QsegwsR+4vtAiFYuNGPAAAA//9dpnuJ"
}
