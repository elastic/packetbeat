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

//go:build !integration
// +build !integration

package collstats

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/elastic-agent-libs/mapstr"
)

func TestEventMapping(t *testing.T) {

	content, err := ioutil.ReadFile("./_meta/test/input.json")
	assert.NoError(t, err)

	data := mapstr.M{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		t.Fatal(err)
	}

	event, _ := eventMapping("unit.test", data)

	assert.Equal(t, event["total"].(mapstr.M)["count"], float64(1))
}
