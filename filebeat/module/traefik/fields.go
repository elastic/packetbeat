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

package traefik

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "traefik", asset.ModuleFieldsPri, AssetTraefik); err != nil {
		panic(err)
	}
}

// AssetTraefik returns asset data.
// This is the base64 encoded gzipped contents of module/traefik.
func AssetTraefik() string {
	return "eJyslk9vozwQxu/5FKPei1S97ymHvVSqtIe9rHpHLgzgjfGw43Ervv3KJKQU2ZQQfAvEz++ZPx7zCCfsjyCssNKnA4BoMXiEh9fzk4cDQImuYN2JJnuEHwcAgF9UeoNQEUOn2GlbgzQIl01gqIZKG3TZAaDSaEp3HPY9glUtTnlhSd/hEWom312eRJBhvQxSUDG1aV5YU+aUq4oCnbs+jqEX8GE9kxWlrbsghhRMrZwJwdHVTMzQ1JR3yLku0YquNPKX/4wOT9h/EJezdws+w/rpBmu/X57h6f+n/+DMkB6oGl4URqOVqCfGvx6d5AX52T9GR4ZsfZud1wbB+vYNORi4EFwUXzFZQVvm4ed+CRkcqBbHBIyYUIIyauRNFafgw7PZaCNqwrMZPVwI8NEg45gV0ENnfSguJ8biDqnsc4dWsrde0EVdKqPV/E2npDlCI9JljK4j6zALWlGZVteszlkV9phomZYEc93daMGR5wIzVZb89WyuBQ/nJ9koaXDYl0X2rWG2KA3Ny74y2UOFs6jCqnATrbgQKJuMWNfaqvnWNcBgO39HdprslojjW9e11Lkx84LKW6v7tbOdKPEuprPOR4XMidG8st4JjTV4Vc/H9LrmzoeNt5Y+fcbSPuYXKCSuvKlkie+6mFdjObRoeAmd60lVf2ie9A2UuMwVou0ukKjMCOmUFM39kLjMCInMww2MqMqIoJjEjQByWeWNiY39KSjfqfrkvmmAgNqnBwJqsQ3IxaPeQlpOX3xubCHNlaIzpkZKfDxsGS8FWdEWrdyTrcu3SY2UfaP3ifVWuM+1o9htswn8jeKINlQM8/x+ZFLp80KsNdmdMrskdk2rln6vQialZtHtV8KU4OFfAAAA///Vouzi"
}
