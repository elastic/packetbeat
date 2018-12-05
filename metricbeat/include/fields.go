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
	if err := asset.SetFields("metricbeat", "../metricbeat/_meta/fields.common.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJyUlEFu2zAQRfc+xUf20QG0KFCkBbrqKhcYk+OQCMlRyZFV3T4gLSWxEcmIdzQ1//2Z+dIjXnnuYSRGSQdAvQbu8bSeLReT/aBeUo8fBwB4kqTkU1mKcPIcbAGdyQc6BoZPoBDAZ04KnQcu3QHLY/2haTwiUeQekTV7U1i7KHYM3C6/pNbfs+NWBzlBHeNSA3WkeOHEmZRtu2nsbotVz98krbXfhzkpug/7I0WvYGScT4xTlojJeeNuPExUhx8CG2Xb4dn58i7WxoxIM5Iojowhc6mLmBynpmNJ6VoCQQyFMG/2kHVtoa6zR5D0svyR+d/oM9semsc7U/3dEpFlTBaa/QD1scUlepOlsJFky+7WykCGr6y88jxJtvvgv2tpHbGdE0VvPpRvkVV4rzv+T3EIn72V+2myYsa4vhAdfoaJ5oKWKMGDFfNw46JwPnvDn8P6DuZARSuZsnH3W1+DtUguvo9MihOrcVw+glFDd+Pkkm87ZqriW0HY4P9aylYPl6+CT0iUBF/vfAGSUuHN4O1mbKltet3hLQAA///iTYXi"
}
