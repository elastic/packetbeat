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
	if err := asset.SetFields("metricbeat", "logstash", asset.ModuleFieldsPri, AssetLogstash); err != nil {
		panic(err)
	}
}

// AssetLogstash returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/logstash.
func AssetLogstash() string {
	return "eJy0k8GOmzAQhu88xS/ODQ/Aoae2aqpW7SmXqqoQTIg3xoM8A6u8/YoEIiBmN5ts5hApNnz/N/awwp4OKSyXopnsIkCNWkoR/+yX4ggoSHJvajXsUnyOAGDYRsVFYykCPFnKhFKUWQQIqRpXSoq/sYiNPyHeqdbxvwjYGrKFpEfOCi6raGLQlR7qjuS5qfuVgMOUNKY5Lui8GKItEk81e3/a2lDz8LHAjkWT7qf7O3lisNnT4Zl9Mdt7xamr7yyKC+QQKuRbk1PSkhfD7uNiz5e9CZCH9Ke2CibOj/2KvB+bX1i7Lc82Qgc+Ngj3/VbvV/h0FWp9HF57zkkkqU0o4CRg2ZW3pf850bH+Epz1RDRTuXfi/x8pqEi9ySW56wOgltzEaNnqiiP4eqRh3uWSxFjELM/D7dexdjlXxpV9m8i5cUo+WbTgRh+g8bvRkt+jsTVWydMjJvRbj75weQkAAP//rimNWQ=="
}
