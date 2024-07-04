// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build integration

package integration

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/reload"
	lbmanagement "github.com/elastic/beats/v7/libbeat/management"
	"github.com/elastic/beats/v7/x-pack/filebeat/cmd"
	"github.com/elastic/beats/v7/x-pack/libbeat/management"
	"github.com/elastic/beats/v7/x-pack/libbeat/management/tests"
	"github.com/elastic/elastic-agent-client/v7/pkg/client"
	"github.com/elastic/elastic-agent-client/v7/pkg/client/mock"
	"github.com/elastic/elastic-agent-client/v7/pkg/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	conf "github.com/elastic/elastic-agent-libs/config"
)

func TestLogStatusReporter(t *testing.T) {
	unitOneID := mock.NewID()
	unitOutID := mock.NewID()
	token := mock.NewID()

	tests.InitBeatsForTest(t, cmd.Filebeat())
	tmpDir := t.TempDir()
	filename := fmt.Sprintf("test-%d", time.Now().Unix())
	outPath := filepath.Join(tmpDir, filename)
	t.Logf("writing output to file %s", outPath)
	err := os.Mkdir(outPath, 0775)
	require.NoError(t, err)
	defer func() {
		err := os.RemoveAll(outPath)
		require.NoError(t, err)
	}()

	/*
	 * valid input stream, shouldn't raise any error.
	 */
	inputStream := getInputStream(unitOneID, filepath.Join(tmpDir, "*"), 2)

	/*
	 * try to open an irregular file.
	 * This should throw "Tried to open non regular file:" and status to degraded
	 */
	inputStreamIrregular := getInputStream(unitOneID, "/dev/null", 1)

	outputExpectedStream := proto.UnitExpected{
		Id:             unitOutID,
		Type:           proto.UnitType_OUTPUT,
		ConfigStateIdx: 1,
		State:          proto.State_HEALTHY,
		Config: &proto.UnitExpectedConfig{
			DataStream: &proto.DataStream{
				Namespace: "default",
			},
			Type:     "file",
			Revision: 1,
			Meta: &proto.Meta{
				Package: &proto.Package{
					Name:    "system",
					Version: "1.17.0",
				},
			},
			Source: tests.RequireNewStruct(map[string]interface{}{
				"type":            "file",
				"enabled":         true,
				"path":            outPath,
				"filename":        "beat-out",
				"number_of_files": 7,
			}),
		},
	}

	observedStates := make(chan *proto.CheckinObserved)
	expectedUnits := make(chan []*proto.UnitExpected)
	done := make(chan struct{})
	// V2 mock server
	server := &mock.StubServerV2{
		CheckinV2Impl: func(observed *proto.CheckinObserved) *proto.CheckinExpected {
			select {
			case observedStates <- observed:
				return &proto.CheckinExpected{
					Units: <-expectedUnits,
				}
			case <-done:
				return nil
			}
		},
		ActionImpl: func(response *proto.ActionResponse) error {
			return nil
		},
	}
	server.Start()
	defer server.Stop()

	// start the client
	client := client.NewV2(fmt.Sprintf(":%d", server.Port), token, client.VersionInfo{
		Name: "program",
		Meta: map[string]string{
			"key": "value",
		},
	}, client.WithGRPCDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))

	lbmanagement.SetManagerFactory(func(cfg *conf.C, registry *reload.Registry) (lbmanagement.Manager, error) {
		c := management.DefaultConfig()
		if err := cfg.Unpack(&c); err != nil {
			return nil, err
		}
		return management.NewV2AgentManagerWithClient(c, registry, client, management.WithStopOnEmptyUnits)
	})

	go func() {
		t.Logf("Running beats...")
		err := cmd.Filebeat().Execute()
		require.NoError(t, err)
	}()

	scenarios := []struct {
		expectedStatus proto.State
		nextInputunit  *proto.UnitExpected
	}{
		{
			proto.State_HEALTHY,
			&inputStreamIrregular,
		},
		{
			proto.State_DEGRADED,
			&inputStream,
		},
		{
			proto.State_HEALTHY,
			&inputStream,
		},
		// wait for one more checkin, just to be sure it's healthy
		{
			proto.State_HEALTHY,
			&inputStream,
		},
	}

	timer := time.NewTimer(2 * time.Minute)
	id := 0
	for id < len(scenarios) {
		select {
		case observed := <-observedStates:
			state := extractState(observed.GetUnits(), unitOneID)
			expectedUnits <- []*proto.UnitExpected{
				scenarios[id].nextInputunit,
				&outputExpectedStream,
			}
			if state != scenarios[id].expectedStatus {
				continue
			}
			fmt.Println("WHATSPUPP", state, scenarios[id].expectedStatus)
			// always ensure that output is healthy
			outputState := extractState(observed.GetUnits(), unitOutID)
			require.Equal(t, outputState, proto.State_HEALTHY)

			timer.Reset(2 * time.Minute)
			id++
		case <-timer.C:
			t.Fatal("timeout waiting for checkin")
		default:
		}
	}
}

func extractState(units []*proto.UnitObserved, idx string) proto.State {
	for _, unit := range units {
		if unit.Id == idx {
			return unit.GetState()
		}
	}
	return -1
}

func getInputStream(id string, path string, stateIdx int) proto.UnitExpected {
	return proto.UnitExpected{
		Id:             id,
		Type:           proto.UnitType_INPUT,
		ConfigStateIdx: uint64(stateIdx),
		State:          proto.State_HEALTHY,
		Config: &proto.UnitExpectedConfig{
			Streams: []*proto.Stream{{
				Id: "filebeat/log-default-system",
				Source: tests.RequireNewStruct(map[string]interface{}{
					"enabled":  true,
					"symlinks": true,
					"type":     "log",
					"paths":    []interface{}{path},
				}),
			}},
			Type:     "log",
			Id:       "log-input-test",
			Name:     "log-1",
			Revision: 1,
		},
	}
}
