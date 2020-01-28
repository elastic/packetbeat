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

package generate

import (
	"fmt"
	"os"

	devtools "github.com/elastic/beats/dev-tools/mage"
	genfields "github.com/elastic/beats/filebeat/generator/fields"
)

// Fields creates a new fields.yml for an existing Filebeat fileset.
// Use MODULE=module to specify the name of the existing module
// Use FILESET=fileset to specify the name of the existing fileset
func Fields() error {
	targetModule := os.Getenv("MODULE")
	targetFileset := os.Getenv("FILESET")

	if targetModule == "" || targetFileset == "" {
		return fmt.Errorf("You must specify the module and fileset: MODULE=module FILESET=fileset mage generate:fields")
	}

	curDir := devtools.CWD()

	return genfields.Generate(curDir, targetModule, targetFileset, false)
}
