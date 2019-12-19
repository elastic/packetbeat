// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package elb

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/bus"
	awsauto "github.com/elastic/beats/x-pack/libbeat/autodiscover/providers/aws"
	"github.com/elastic/beats/x-pack/libbeat/autodiscover/providers/aws/test"
)

func Test_internalBuilder(t *testing.T) {
	lbl := fakeLbl()
	lbls := []*lbListener{lbl}
	fetcher := newMockFetcher(lbls, nil)
	pBus := bus.New("test")

	cfg := &awsauto.Config{
		Regions: []string{"us-east-1a", "us-west-1b"},
		Period:  time.Nanosecond,
	}

	uuid, _ := uuid.NewV4()
	provider, err := internalBuilder(uuid, pBus, cfg, fetcher)
	require.NoError(t, err)

	startListener := pBus.Subscribe("start")
	stopListener := pBus.Subscribe("stop")
	listenerDone := make(chan struct{})
	defer close(listenerDone)

	var events test.TestEventAccumulator
	go func() {
		for {
			select {
			case e := <-startListener.Events():
				events.Add(e)
			case e := <-stopListener.Events():
				events.Add(e)
			case <-listenerDone:
				return
			}
		}
	}()

	// Let run twice to ensure that duplicates don't create two start events
	// Since we're turning a list of assets into a list of changes the second once() call should be a noop
	provider.watcher.once()
	provider.watcher.once()
	events.WaitForNumEvents(t, 1, time.Second)

	assert.Equal(t, 1, events.Len())

	expectedStartEvent := bus.Event{
		"id":       lbl.arn(),
		"provider": uuid,
		"start":    true,
		"host":     *lbl.lb.DNSName,
		"port":     *lbl.listener.Port,
		"meta": common.MapStr{
			"elb_listener": lbl.toMap(),
			"cloud":        lbl.toCloudMap(),
		},
	}

	require.Equal(t, expectedStartEvent, events.Get()[0])

	fetcher.setLbls([]*lbListener{})

	// Let run twice to ensure that duplicates don't cause an issue
	provider.watcher.once()
	provider.watcher.once()
	events.WaitForNumEvents(t, 2, time.Second)

	require.Equal(t, 2, events.Len())

	expectedStopEvent := bus.Event{
		"stop":     true,
		"id":       lbl.arn(),
		"provider": uuid,
	}

	require.Equal(t, expectedStopEvent, events.Get()[1])

	// Test that in an error situation nothing changes.
	preErrorEventCount := events.Len()
	fetcher.setError(errors.New("oops"))

	// Let run twice to ensure that duplicates don't cause an issue
	provider.watcher.once()
	provider.watcher.once()

	assert.Equal(t, preErrorEventCount, events.Len())
}
