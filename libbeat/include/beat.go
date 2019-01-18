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
	if err := asset.SetFields("libbeat", "Beat", asset.BeatFieldsPri, AssetBeat); err != nil {
		panic(err)
	}
}

// AssetBeat returns asset data.
// This is the base64 encoded gzipped contents of ./build/fields/beat.yml.
func AssetBeat() string {
	return "eJyMUl1OwzAMfu8pvgtsB8gDDyDOgbzU7czSuErSoXF6lKSDpUKFPS22vz+73QEXvhmcmFIHJEmODZ7rq+dog8xJ1Bs8dQDwoj6R+Air06S+4DAIuz6CriSOTo4hHuQc+Mo+Id1mjscO65gpPAd4mrjqHpNM/KmeSwcFYEBOKK6VmdLZVLrt8CRjoOowhYW7hr0qNrR6emeb1lJ9vNXOhW8fGvq19Uv2Jv8SOcCqH2RcQgldxY6tAw5BQ2NgDLrM+yKvGXTfqq2K4kdQ30ueJQfxg+Y1W4oMHapO2XL+PW760U128F28G2pz79j6sZaBm6DlkPnvH0c8a2zmdu9XSDPiH8Q05q9jM7xl/woAAP//8G7fyA=="
}
