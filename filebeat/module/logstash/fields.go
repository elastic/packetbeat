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

package logstash

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "logstash", asset.ModuleFieldsPri, AssetLogstash); err != nil {
		panic(err)
	}
}

// AssetLogstash returns asset data.
// This is the base64 encoded gzipped contents of module/logstash.
func AssetLogstash() string {
	return "eJzslc9u2zAMxu95CqLnNg/gQy/DChTYH2Db3VAsWtYiiYZEJfXbD7bjxHbkDO6CnqZbxIjfj59I6wn22GRgSAUWodoAsGaDGTwMWw8bAImh8LpmTS6D5w0AnE/AV5LR4Aag1GhkyLroEzhhcZK3XdzUmIHyFOvTTiLzNNMs23nvAvplBHqOXQktivXrpZOE0pMFrhCGpF0F29Ff52xjPjtYMV49yh6bI3k5i90AatevCk85gTwURoQAxwo9doh4QMdAXivtBOM2icSVRzGX7ZEY33gdz6sryVvRhkHsKHLH4aNz2qmT1AjQkPobniGVd3UkCWn3G4uVjHtsQDgJB2EigsRdVKql0xf2C0r6FjEEodLXKIwWYRapBVdLp6xWXvSo7COmPcADmpVqhtQ2dW5Jb9AKho6zKVo7Kdcp/o/EPUdieRzeZdgPcQQZbQ1UdgQndZPUsdGwzlP3OTHw2qUbFg6nahOVdnn7437VfRMWh8p6gVvardAdW7GpZ9oZvLo6cniEF20YfXiE75HbnbZZP5HEIiz0JNE+1y632hg9H/me0ZBT6wA/v2ERu85kbRFK8iNW0A56NSzIyQWuk3G18MKmsd5l3U/27Xj03+iJhVCQK7WK/Vfs49uzrzRPvjz/9ig9PU/rnRQKMaCEXTNyYuHB/IDHAhKN6YSjdAMs63ZfmK1cuMuZ+J8AAAD//5+nw5A="
}
