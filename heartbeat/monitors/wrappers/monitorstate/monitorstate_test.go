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

package monitorstate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecordingAndFlapping(t *testing.T) {
	monitorID := "test"
	ms := newMonitorState(monitorID, StatusUp, 0)
	recordFlappingSeries(monitorID, ms)
	require.Equal(t, StatusFlapping, ms.Status)
	require.Equal(t, FlappingThreshold+1, ms.Checks)
	require.Equal(t, ms.Up+ms.Down, ms.Checks)

	// Use double the flapping threshold so any transitions after this are stable
	priorChecksCount := ms.Checks
	recordStableSeries(monitorID, ms, FlappingThreshold*2, StatusDown)
	require.Equal(t, StatusDown, ms.Status)
	// The count should be FlappingThreshold+1 since we used double the threshold before
	// This is because we have one full threshold of stable checks, as well as the final check that
	// flipped us out of the threshold, which goes toward the new state.
	requireMSCounts(t, ms, 0, FlappingThreshold+1)
	require.Equal(t, priorChecksCount+FlappingThreshold-1, ms.Ends.Checks)
	// We don't want to store the entire state history!
	require.Empty(t, ms.Ends.FlapHistory)

	// Since we're now in a stable state a single up check should create a new state from a stable one
	ms.recordCheck(monitorID, StatusUp)
	require.Equal(t, StatusUp, ms.Status)
	requireMSCounts(t, ms, 1, 0)
}

// recordFlappingSeries is a helper that should always put the monitor into a flapping state.
func recordFlappingSeries(monitorID string, ms *State) {
	for i := 0; i < FlappingThreshold; i++ {
		if i%2 == 0 {
			ms.recordCheck(monitorID, StatusUp)
		} else {
			ms.recordCheck(monitorID, StatusDown)
		}
	}
}

// recordStableSeries is a test helper for repeatedly recording one status
func recordStableSeries(monitorID string, ms *State, count int, s StateStatus) {
	for i := 0; i < count; i++ {
		ms.recordCheck(monitorID, s)
	}
}

func TestTransitionTo(t *testing.T) {
	id := "mymonitor"
	s := newMonitorState(id, StatusUp, 0)
	first := *s
	s.transitionTo(id, StatusDown)
	second := *s
	s.transitionTo(id, StatusUp)

	require.NotEqual(t, s.ID, second.ID)
	require.NotEqual(t, s.ID, first)

	require.Equal(t, second.ID, s.Ends.ID)
	// Ensure No infinite storage of states
	require.Nil(t, s.Ends.Ends)
}
