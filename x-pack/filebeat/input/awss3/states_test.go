// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package awss3

import (
	"fmt"
	"testing"
	"time"

	"github.com/elastic/beats/v7/filebeat/beater"
	"github.com/elastic/beats/v7/libbeat/statestore"
	"github.com/elastic/beats/v7/libbeat/statestore/storetest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testInputStore struct {
	registry *statestore.Registry
}

func openTestStatestore() beater.StateStore {
	return &testInputStore{
		registry: statestore.NewRegistry(storetest.NewMemoryStoreBackend()),
	}
}

func (s *testInputStore) Close() {
	_ = s.registry.Close()
}

func (s *testInputStore) Access() (*statestore.Store, error) {
	return s.registry.Get("filebeat")
}

func (s *testInputStore) CleanupInterval() time.Duration {
	return 24 * time.Hour
}

func TestStatesAddStateAndIsProcessed(t *testing.T) {
	type stateTestCase struct {
		// An initialization callback to invoke on the (initially empty) states.
		statesEdit func(states *states) error

		// The state to call IsProcessed on and the expected result
		state               state
		expectedIsProcessed bool

		// If true, the test will run statesEdit, then create a new states
		// object from the same persistent store before calling IsProcessed
		// (to test persistence between restarts).
		shouldReload bool
	}
	lastModified := time.Date(2022, time.June, 30, 14, 13, 00, 0, time.UTC)
	testState1 := newState("bucket", "key", "etag", lastModified)
	testState2 := newState("bucket1", "key1", "etag1", lastModified)
	tests := map[string]stateTestCase{
		"with empty states": {
			state:               testState1,
			expectedIsProcessed: false,
		},
		"not existing state": {
			statesEdit: func(states *states) error {
				return states.AddState(testState2)
			},
			state:               testState1,
			expectedIsProcessed: false,
		},
		"existing state": {
			statesEdit: func(states *states) error {
				return states.AddState(testState1)
			},
			state:               testState1,
			expectedIsProcessed: true,
		},
		"existing stored state is persisted": {
			statesEdit: func(states *states) error {
				state := testState1
				state.Stored = true
				return states.AddState(state)
			},
			state:               testState1,
			shouldReload:        true,
			expectedIsProcessed: true,
		},
		"existing failed state is persisted": {
			statesEdit: func(states *states) error {
				state := testState1
				state.Failed = true
				return states.AddState(state)
			},
			state:               testState1,
			shouldReload:        true,
			expectedIsProcessed: true,
		},
		"existing unprocessed state is not persisted": {
			statesEdit: func(states *states) error {
				return states.AddState(testState1)
			},
			state:               testState1,
			shouldReload:        true,
			expectedIsProcessed: false,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			store := openTestStatestore()
			states, err := newStates(nil, store)
			require.NoError(t, err, "states creation must succeed")
			if test.statesEdit != nil {
				err = test.statesEdit(states)
				require.NoError(t, err, "states edit must succeed")
			}
			if test.shouldReload {
				states, err = newStates(nil, store)
				require.NoError(t, err, "states creation must succeed")
			}

			isProcessed := states.IsProcessed(test.state)
			assert.Equal(t, test.expectedIsProcessed, isProcessed)
		})
	}
}

func TestStatesCleanUp(t *testing.T) {
	bucketName := "test-bucket"
	lModifiedTime := time.Unix(0, 0)
	stateA := newState(bucketName, "a", "a-etag", lModifiedTime)
	stateB := newState(bucketName, "b", "b-etag", lModifiedTime)
	stateC := newState(bucketName, "c", "c-etag", lModifiedTime)

	tests := []struct {
		name       string
		initStates []state
		knownIDs   []string
		expectIDs  []string
	}{
		{
			name:       "No cleanup if not missing from known list",
			initStates: []state{stateA, stateB, stateC},
			knownIDs:   []string{stateA.ID(), stateB.ID(), stateC.ID()},
			expectIDs:  []string{stateA.ID(), stateB.ID(), stateC.ID()},
		},
		{
			name:       "Clean up if missing from known list",
			initStates: []state{stateA, stateB, stateC},
			knownIDs:   []string{stateA.ID()},
			expectIDs:  []string{stateA.ID()},
		},
		{
			name:       "Clean up everything",
			initStates: []state{stateA, stateC}, // given A, C
			knownIDs:   []string{stateB.ID()},   // but known B
			expectIDs:  []string{},              // empty state & store
		},
		{
			name:       "Empty known IDs are valid",
			initStates: []state{stateA}, // given A
			knownIDs:   []string{},      // Known nothing
			expectIDs:  []string{},      // empty state & store
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := openTestStatestore()
			statesInstance, err := newStates(nil, store)
			require.NoError(t, err, "states creation must succeed")

			for _, s := range test.initStates {
				err := statesInstance.AddState(s)
				require.NoError(t, err, "state initialization must succeed")
			}

			// perform cleanup
			err = statesInstance.CleanUp(test.knownIDs)
			require.NoError(t, err, "state cleanup must succeed")

			// validate
			for _, id := range test.expectIDs {
				// must be in local state
				_, ok := statesInstance.states[id]
				require.True(t, ok, fmt.Errorf("expected id %s in state, but got missing", id))

				// must be in store
				ok, err := statesInstance.store.Has(getStoreKey(id))
				require.NoError(t, err, "state has must succeed")
				require.True(t, ok, fmt.Errorf("expected id %s in store, but got missing", id))
			}
		})
	}

}
