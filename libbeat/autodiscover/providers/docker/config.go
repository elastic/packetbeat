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

package docker

import (
	"fmt"
	"strings"

	"github.com/elastic/beats/libbeat/autodiscover/template"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/docker"
)

// Config for docker autodiscover provider
type Config struct {
	Host         string                  `config:"host"`
	TLS          *docker.TLSConfig       `config:"ssl"`
	Prefix       string                  `config:"prefix"`
	Separator    string                  `config:"separator"`
	HintsEnabled bool                    `config:"hints.enabled"`
	Builders     []*common.Config        `config:"builders"`
	Appenders    []*common.Config        `config:"appenders"`
	Templates    template.MapperSettings `config:"templates"`
}

func defaultConfig() *Config {
	return &Config{
		Host:      "unix:///var/run/docker.sock",
		Prefix:    "co.elastic",
		Separator: "/",
	}
}

// Validate ensures correctness of config
func (c *Config) Validate() error {
	// Make sure that prefix doesn't ends with a '.'
	if c.Prefix[len(c.Prefix)-1] == '.' && c.Prefix != "." {
		c.Prefix = c.Prefix[:len(c.Prefix)-2]
	}
	if len(c.Separator) != 1 {
		return fmt.Errorf("separator '%v' must be a single character", c.Separator)
	}
	// Make sure that separator isn't included in prefix
	if strings.Contains(c.Prefix, c.Separator) {
		return fmt.Errorf("separator '%v' must not be included in prefix '%v'", c.Prefix, c.Separator)
	}
	return nil
}
