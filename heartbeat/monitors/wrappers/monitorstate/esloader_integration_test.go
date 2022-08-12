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
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"

	"github.com/elastic/beats/v7/heartbeat/esutil"
)

func TestStatesESLoader(t *testing.T) {
	testStart := time.Now()
	etc := newESTestContext(t)

	// Create three monitors in ES, load their states, and make sure we track them correctly
	// We create a few to make sure the query isolates the monitors correctly
	// and alternate between testing monitors that start up or down
	for i := 0; i < 10; i++ {
		testStatus := StatusUp
		if i%2 == 1 {
			testStatus = StatusDown
		}

		monID := etc.createTestMonitorStateInES(t, testStatus)
		// Since we've continued this state it should register the initial state
		ms := etc.tracker.getCurrentState(monID)
		require.True(t, ms.StartedAt.After(testStart.Add(-time.Nanosecond)), "timestamp for new state is off")
		requireMSStatusCount(t, ms, testStatus, 1)

		// Write the state a few times, enough to guarantee a stable state
		count := FlappingThreshold * 2
		var lastId string
		for i := 0; i < count; i++ {
			ms = etc.tracker.RecordStatus(monID, testStatus)
			if i == 0 {
				lastId = ms.ID
			}
			require.Equal(t, lastId, ms.ID, "state ID should not change within state")
		}
		// The initial state adds 1 to count
		requireMSStatusCount(t, ms, testStatus, count+1)

		// now change the state
		if testStatus == StatusUp {
			testStatus = StatusDown
		} else {
			testStatus = StatusUp
		}

		origMsId := ms.ID
		for i := 0; i < count; i++ {
			ms = etc.tracker.RecordStatus(monID, testStatus)
			require.NotEqual(t, origMsId, ms.ID)
			if i == 0 {
				lastId = ms.ID
				require.Equal(t, origMsId, ms.Ends.ID, "transition should point to the prior state")
			}
			require.Equal(t, lastId, ms.ID, "state ID should not change within state")
		}
		requireMSStatusCount(t, ms, testStatus, count)
	}
}

type esTestContext struct {
	namespace string
	esc       *elasticsearch.Client
	loader    StateLoader
	tracker   *Tracker
}

func newESTestContext(t *testing.T) *esTestContext {
	namespace, _ := uuid.NewV4()
	esc := integES(t)
	etc := &esTestContext{
		namespace: namespace.String(),
		esc:       esc,
		loader:    MakeESLoader(esc, fmt.Sprintf("synthetics-*-%s", namespace.String())),
	}

	etc.tracker = NewTracker(etc.loader)

	return etc
}

func (etc *esTestContext) createTestMonitorStateInES(t *testing.T, s StateStatus) (id string) {
	mUUID, _ := uuid.NewV4()
	mID := mUUID.String()
	mType := "testtyp"
	initState := newMonitorState(mID, s, 0)
	etc.setInitialState(t, mType, mID, initState)
	return mID
}

func (etc *esTestContext) setInitialState(t *testing.T, typ string, monitorID string, ms *State) {
	idx := fmt.Sprintf("synthetics-%s-%s", typ, etc.namespace)

	type Mon struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	}

	reqBodyRdr, err := esutil.ToJsonRdr(struct {
		Ts      time.Time `json:"@timestamp"`
		Monitor Mon       `json:"monitor"`
		State   *State    `json:"state"`
	}{
		Ts:      time.Now(),
		Monitor: Mon{Id: monitorID, Type: typ},
		State:   ms,
	})

	_, err = esutil.CheckRetResp(etc.esc.Index(idx, reqBodyRdr, func(request *esapi.IndexRequest) {
		// Refresh the index since we tend to re-query immediately, otherwise this would miss
		request.Refresh = "true"

	}))
	require.NoError(t, err)
}

func integES(t *testing.T) (esc *elasticsearch.Client) {
	esc, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://127.0.0.1:9200"},
		Username:  "admin",
		Password:  "testing",
	})
	require.NoError(t, err)
	respBody, err := esc.Cluster.Health()
	healthRaw, err := esutil.CheckRetResp(respBody, err)
	require.NoError(t, err)

	healthResp := struct {
		Status string `json:"status"`
	}{}
	err = json.Unmarshal(healthRaw, &healthResp)
	require.NoError(t, err)
	require.Contains(t, []string{"green", "yellow"}, healthResp.Status)

	return esc
}
