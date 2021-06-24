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

// +build !integration
// +build darwin freebsd linux openbsd windows

package memory

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/v7/metricbeat/internal/metrics"
)

func TestGetMemory(t *testing.T) {
	mem, err := Get("")

	assert.NotNil(t, mem)
	assert.NoError(t, err)

	assert.True(t, (mem.Total.ValueOr(0) > 0))
	assert.True(t, (mem.Used.Bytes.ValueOr(0) > 0))
	assert.True(t, (mem.Free.ValueOr(0) >= 0))
	assert.True(t, (mem.Actual.Free.ValueOr(0) >= 0))
	assert.True(t, (mem.Actual.Used.Bytes.ValueOr(0) > 0))
}

func TestGetSwap(t *testing.T) {
	if runtime.GOOS == "freebsd" {
		return //no load data on freebsd
	}

	mem, err := Get("")

	assert.NotNil(t, mem)
	assert.NoError(t, err)

	assert.True(t, (mem.Swap.Total.ValueOr(0) >= 0))
	assert.True(t, (mem.Swap.Used.Bytes.ValueOr(0) >= 0))
	assert.True(t, (mem.Swap.Free.ValueOr(0) >= 0))
}

func TestMemPercentage(t *testing.T) {
	m := Memory{
		Total: metrics.OptUintWith(7),
		Used:  UsedMemStats{Bytes: metrics.OptUintWith(5)},
		Free:  metrics.OptUintWith(2),
	}
	m.fillPercentages()
	assert.Equal(t, m.Used.Pct.ValueOr(0), 0.7143)

	m = Memory{
		Total: metrics.OptUintWith(0),
	}
	m.fillPercentages()
	assert.Equal(t, m.Used.Pct.ValueOr(0), 0.0)
}

func TestActualMemPercentage(t *testing.T) {
	m := Memory{
		Total: metrics.OptUintWith(7),
		Actual: ActualMemoryMetrics{
			Used: UsedMemStats{Bytes: metrics.OptUintWith(5)},
			Free: metrics.OptUintWith(2),
		},
	}

	m.fillPercentages()
	assert.Equal(t, m.Actual.Used.Pct.ValueOr(0), 0.7143)

	m = Memory{
		Total: metrics.OptUintWith(0),
	}
	m.fillPercentages()
	assert.Equal(t, m.Actual.Used.Pct.ValueOr(0), 0.0)
}
