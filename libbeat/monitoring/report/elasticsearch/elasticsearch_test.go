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

package elasticsearch

import (
	"testing"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/monitoring/report"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/stretchr/testify/require"
)

func TestMakeClientParams(t *testing.T) {
	var params, expected map[string]string
	params = map[string]string{
		"foo": "bar",
	}
	expected = map[string]string{
		"foo": "bar",
	}

	p := makeClientParams(config{
		Params: params,
	})

	require.Equal(t, expected, p)
}

func TestMakeReporter(t *testing.T) {
	c, err := conf.NewConfigFrom(map[string]interface{}{
		"hosts": []string{"127.0.0.1"},
	})
	require.NoError(t, err)

	r, err := makeReporter(beat.Info{}, report.Settings{}, c)
	require.NoError(t, err)
	r.Stop()
}
