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

package docs

import (
	"github.com/magefile/mage/mg"

	"github.com/elastic/beats/dev-tools/mage"
)

var (
	docsDeps []interface{}
)

// RegisterDeps registers dependencies of the Docs target.
func RegisterDeps(deps ...interface{}) {
	docsDeps = append(docsDeps, deps...)
}

// Docs generates the documentation for the Beat. Set PREVIEW=true to
// automatically open the browser to the docs.
func Docs() error {
	mg.SerialDeps(docsDeps...)
	return mage.Docs.AsciidocBook()
}
