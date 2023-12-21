// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package salesforce

import (
	"testing"
	"time"

	v2 "github.com/elastic/beats/v7/filebeat/input/v2"
	cursor "github.com/elastic/beats/v7/filebeat/input/v2/input-cursor"
	"github.com/elastic/beats/v7/libbeat/statestore"
	"github.com/elastic/beats/v7/libbeat/statestore/storetest"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/go-concert/unison"
	"github.com/stretchr/testify/assert"
)

// compile-time check if querier implements InputManager
var _ v2.InputManager = InputManager{}

func makeTestStore(data map[string]interface{}) *statestore.Store {
	memstore := &storetest.MapStore{Table: data}
	reg := statestore.NewRegistry(&storetest.MemoryStore{
		Stores: map[string]*storetest.MapStore{
			"test": memstore,
		},
	})
	store, err := reg.Get("test")
	if err != nil {
		panic("failed to create test store")
	}
	return store
}

type stateStore struct{}

func (stateStore) Access() (*statestore.Store, error) {
	return makeTestStore(map[string]interface{}{"hello": "world"}), nil
}
func (stateStore) CleanupInterval() time.Duration { return time.Duration(0) }

// compile-time check if stateStore implements cursor.StateStore
var _ cursor.StateStore = stateStore{}

func TestInputManager(t *testing.T) {
	inputManager := NewInputManager(logp.NewLogger("salesforce_test"), stateStore{})

	var inputTaskGroup unison.TaskGroup
	defer func() {
		_ = inputTaskGroup.Stop()
	}()

	err := inputManager.Init(&inputTaskGroup, v2.ModeRun)
	assert.NoError(t, err)

	config, err := conf.NewConfigFrom(map[string]interface{}{
		"url":     "https://salesforce.com",
		"version": 46,
		"auth": &authConfig{
			JWT: &jwtConfig{
				Enabled:        pointer(true),
				URL:            "https://salesforce.com",
				ClientID:       "xyz",
				ClientUsername: "xyz",
				ClientKeyPath:  "xyz",
			},
		},
		"data_collection_method": &DataCollectionMethod{
			Object: ObjectMethod{Enabled: true, Interval: 4},
		},
	})
	assert.NoError(t, err)

	_, err = inputManager.Create(config)
	assert.NoError(t, err)
}

func TestSource(t *testing.T) {
	want := "https://salesforce.com"
	src := source{cfg: config{URL: want}}
	got := src.Name()
	assert.Equal(t, want, got)
}
