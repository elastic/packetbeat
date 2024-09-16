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

package prometheus

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "prometheus", asset.ModuleFieldsPri, AssetPrometheus); err != nil {
		panic(err)
	}
}

// AssetPrometheus returns asset data.
// This is the base64 encoded zlib format compressed contents of module/prometheus.
func AssetPrometheus() string {
	return "eJzMk7mO20AMhns9xQ+lW3j3AVSkSxfkQMogEMYjSprsXCGpNfz2gU7LsYNcTabkP0N+/Dl8xDOdK2ROgbSnQQpAnXqqUH7YgmUBNCSWXVaXYoXXBQB8UqMCsWwyNWg5BRhcXoFik5OL+lQA0ifW2qbYuq5Ca7xQATB5MkIVOjPeIVUXO6nwuRTx5QFlr5rLLwXQOvKNVFPdR0QTqEIgZWeltmmIOimAnjNV8Cl2S+AO93jeDeFIjNSuWZCJ8cYbUWeFDNseTbJDoJn/UvXKq0vJjtOQl8gedjyv8J4bYjiBCzmxmqjoiekAb47kBSfnPYJR26N1LHqA9gQmURgmNGk4etryrSjz46eHTVhh0vErWd2F50A9q890PiVudvJPTBrPbp6zU0vVG5jFxz+m+aG3K7UOJmcXu+Vq+VD+JfQN7beB+Py/sb4YP0xTH7yuGzXKH9/u+K935k5XNz3tv+YvYKYE6yhp78O9src/fQVhCkmpPrFT+heeOQ+mPNumbr4stgnxC/Fvs34PAAD//zFqb0I="
}
