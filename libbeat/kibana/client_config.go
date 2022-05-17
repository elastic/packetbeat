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

package kibana

import (
	"fmt"

	"github.com/elastic/elastic-agent-libs/transport/httpcommon"
)

// ClientConfig to connect to Kibana
type ClientConfig struct {
	Protocol     string `config:"protocol" yaml:"protocol,omitempty"`
	Host         string `config:"host" yaml:"host,omitempty"`
	Path         string `config:"path" yaml:"path,omitempty"`
	SpaceID      string `config:"space.id" yaml:"space.id,omitempty"`
	Username     string `config:"username" yaml:"username,omitempty"`
	Password     string `config:"password" yaml:"password,omitempty"`
	APIKey       string `config:"api_key" yaml:"api_key,omitempty"`
	ServiceToken string `config:"service_token" yaml:"service_token,omitempty"`

	// Headers holds headers to include in every request sent to Kibana.
	Headers map[string]string `config:"headers" yaml:"headers,omitempty"`

	IgnoreVersion bool

	Transport httpcommon.HTTPTransportSettings `config:",inline" yaml:",inline"`
}

// DefaultClientConfig connects to a locally running kibana over HTTP
func DefaultClientConfig() ClientConfig {
	return ClientConfig{
		Protocol:     "http",
		Host:         "localhost:5601",
		Path:         "",
		SpaceID:      "",
		Username:     "",
		Password:     "",
		APIKey:       "",
		ServiceToken: "",
		Transport:    httpcommon.DefaultHTTPTransportSettings(),
	}
}

func (c *ClientConfig) Validate() error {
	if c.APIKey != "" && (c.Username != "" || c.Password != "") {
		return fmt.Errorf("cannot set both api_key and username/password")
	}

	return nil
}
