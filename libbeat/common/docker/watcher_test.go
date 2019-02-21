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

package docker

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/elastic/beats/libbeat/logp"
)

type MockClient struct {
	// containers to return on ContainerList call
	containers [][]types.Container
	// event list to send on Events call
	events []interface{}

	done chan interface{}
}

func (m *MockClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	res := m.containers[0]
	m.containers = m.containers[1:]
	return res, nil
}

func (m *MockClient) Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error) {
	eventsC := make(chan events.Message)
	errorsC := make(chan error)

	go func() {
		for _, event := range m.events {
			switch e := event.(type) {
			case events.Message:
				eventsC <- e
			case error:
				errorsC <- e
			}
		}
		close(m.done)
	}()

	return eventsC, errorsC
}

func (m *MockClient) ContainerInspect(ctx context.Context, container string) (types.ContainerJSON, error) {
	return types.ContainerJSON{}, errors.New("unimplemented")
}

func TestWatcherInitialization(t *testing.T) {
	watcher := runWatcher(t, true, nil,
		[][]types.Container{
			[]types.Container{
				types.Container{
					ID:              "0332dbd79e20",
					Names:           []string{"/containername", "othername"},
					Image:           "busybox",
					Labels:          map[string]string{"foo": "bar"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
				types.Container{
					ID:              "6ac6ee8df5d4",
					Names:           []string{"/other"},
					Image:           "nginx",
					Labels:          map[string]string{},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
		},
		nil)

	assert.Equal(t, map[string]*Container{
		"0332dbd79e20": &Container{
			ID:     "0332dbd79e20",
			Name:   "containername",
			Image:  "busybox",
			Labels: map[string]string{"foo": "bar"},
		},
		"6ac6ee8df5d4": &Container{
			ID:     "6ac6ee8df5d4",
			Name:   "other",
			Image:  "nginx",
			Labels: map[string]string{},
		},
	}, watcher.Containers())
}

func TestWatcherInitializationShortID(t *testing.T) {
	watcher := runWatcherShortID(t, true, nil,
		[][]types.Container{
			[]types.Container{
				types.Container{
					ID:              "1234567890123",
					Names:           []string{"/containername", "othername"},
					Image:           "busybox",
					Labels:          map[string]string{"foo": "bar"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
				types.Container{
					ID:              "2345678901234",
					Names:           []string{"/other"},
					Image:           "nginx",
					Labels:          map[string]string{},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
		},
		nil, true)

	assert.Equal(t, map[string]*Container{
		"1234567890123": &Container{
			ID:     "1234567890123",
			Name:   "containername",
			Image:  "busybox",
			Labels: map[string]string{"foo": "bar"},
		},
		"2345678901234": &Container{
			ID:     "2345678901234",
			Name:   "other",
			Image:  "nginx",
			Labels: map[string]string{},
		},
	}, watcher.Containers())

	assert.Equal(t, &Container{
		ID:     "1234567890123",
		Name:   "containername",
		Image:  "busybox",
		Labels: map[string]string{"foo": "bar"},
	}, watcher.Container("123456789012"))
}

func TestWatcherAddEvents(t *testing.T) {
	watcher := runWatcher(t, true, nil,
		[][]types.Container{
			[]types.Container{
				types.Container{
					ID:              "0332dbd79e20",
					Names:           []string{"/containername", "othername"},
					Image:           "busybox",
					Labels:          map[string]string{"foo": "bar"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
			[]types.Container{
				types.Container{
					ID:              "6ac6ee8df5d4",
					Names:           []string{"/other"},
					Image:           "nginx",
					Labels:          map[string]string{"label": "value"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
		},
		[]interface{}{
			events.Message{
				Action: "start",
				Actor: events.Actor{
					ID: "6ac6ee8df5d4",
					Attributes: map[string]string{
						"name":  "other",
						"image": "nginx",
						"label": "value",
					},
				},
			},
		},
	)

	assert.Equal(t, map[string]*Container{
		"0332dbd79e20": &Container{
			ID:     "0332dbd79e20",
			Name:   "containername",
			Image:  "busybox",
			Labels: map[string]string{"foo": "bar"},
		},
		"6ac6ee8df5d4": &Container{
			ID:     "6ac6ee8df5d4",
			Name:   "other",
			Image:  "nginx",
			Labels: map[string]string{"label": "value"},
		},
	}, watcher.Containers())
}

func TestWatcherAddEventsShortID(t *testing.T) {
	watcher := runWatcherShortID(t, true, nil,
		[][]types.Container{
			[]types.Container{
				types.Container{
					ID:              "1234567890123",
					Names:           []string{"/containername", "othername"},
					Image:           "busybox",
					Labels:          map[string]string{"foo": "bar"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
			[]types.Container{
				types.Container{
					ID:              "2345678901234",
					Names:           []string{"/other"},
					Image:           "nginx",
					Labels:          map[string]string{"label": "value"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
		},
		[]interface{}{
			events.Message{
				Action: "start",
				Actor: events.Actor{
					ID: "2345678901234",
					Attributes: map[string]string{
						"name":  "other",
						"image": "nginx",
						"label": "value",
					},
				},
			},
		},
		true,
	)

	assert.Equal(t, map[string]*Container{
		"1234567890123": &Container{
			ID:     "1234567890123",
			Name:   "containername",
			Image:  "busybox",
			Labels: map[string]string{"foo": "bar"},
		},
		"2345678901234": &Container{
			ID:     "2345678901234",
			Name:   "other",
			Image:  "nginx",
			Labels: map[string]string{"label": "value"},
		},
	}, watcher.Containers())
}

func TestWatcherUpdateEvent(t *testing.T) {
	watcher := runWatcher(t, true, nil,
		[][]types.Container{
			[]types.Container{
				types.Container{
					ID:              "0332dbd79e20",
					Names:           []string{"/containername", "othername"},
					Image:           "busybox",
					Labels:          map[string]string{"label": "foo"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
			[]types.Container{
				types.Container{
					ID:              "0332dbd79e20",
					Names:           []string{"/containername", "othername"},
					Image:           "busybox",
					Labels:          map[string]string{"label": "bar"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
		},
		[]interface{}{
			events.Message{
				Action: "update",
				Actor: events.Actor{
					ID: "0332dbd79e20",
					Attributes: map[string]string{
						"name":  "containername",
						"image": "busybox",
						"label": "bar",
					},
				},
			},
		},
	)

	assert.Equal(t, map[string]*Container{
		"0332dbd79e20": &Container{
			ID:     "0332dbd79e20",
			Name:   "containername",
			Image:  "busybox",
			Labels: map[string]string{"label": "bar"},
		},
	}, watcher.Containers())
	assert.Equal(t, 0, len(watcher.deleted))
}

func TestWatcherUpdateEventShortID(t *testing.T) {
	watcher := runWatcherShortID(t, true, nil,
		[][]types.Container{
			[]types.Container{
				types.Container{
					ID:              "1234567890123",
					Names:           []string{"/containername", "othername"},
					Image:           "busybox",
					Labels:          map[string]string{"label": "foo"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
			[]types.Container{
				types.Container{
					ID:              "1234567890123",
					Names:           []string{"/containername", "othername"},
					Image:           "busybox",
					Labels:          map[string]string{"label": "bar"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
		},
		[]interface{}{
			events.Message{
				Action: "update",
				Actor: events.Actor{
					ID: "1234567890123",
					Attributes: map[string]string{
						"name":  "containername",
						"image": "busybox",
						"label": "bar",
					},
				},
			},
		},
		true,
	)

	assert.Equal(t, map[string]*Container{
		"1234567890123": &Container{
			ID:     "1234567890123",
			Name:   "containername",
			Image:  "busybox",
			Labels: map[string]string{"label": "bar"},
		},
	}, watcher.Containers())
	assert.Equal(t, 0, len(watcher.deleted))
}

func TestWatcherDie(t *testing.T) {
	clock := newTestClock()
	watcher := runWatcher(t, false, clock,
		[][]types.Container{
			[]types.Container{
				types.Container{
					ID:              "0332dbd79e20",
					Names:           []string{"/containername", "othername"},
					Image:           "busybox",
					Labels:          map[string]string{"label": "foo"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
		},
		[]interface{}{
			events.Message{
				Action: "die",
				Actor: events.Actor{
					ID: "0332dbd79e20",
				},
			},
		},
	)
	defer watcher.Stop()

	// Check it doesn't get removed while we request meta for the container
	for i := 0; i < 18; i++ {
		watcher.Container("0332dbd79e20")
		assert.Equal(t, 1, len(watcher.Containers()))
		clock.adjust(1 * time.Second)
	}

	// Checks after 10s for the watcher containers to be updated
	l := watcher.bus.Subscribe()
	defer l.Stop()
	clock.adjust(11 * time.Second)
	e := <-l.Events()
	assert.Equal(t, true, e["delete"])
	assert.Equal(t, 0, len(watcher.Containers()))
}

func TestWatcherDieShortID(t *testing.T) {
	clock := newTestClock()
	watcher := runWatcherShortID(t, false, clock,
		[][]types.Container{
			[]types.Container{
				types.Container{
					ID:              "0332dbd79e20aaa",
					Names:           []string{"/containername", "othername"},
					Image:           "busybox",
					Labels:          map[string]string{"label": "foo"},
					NetworkSettings: &types.SummaryNetworkSettings{},
				},
			},
		},
		[]interface{}{
			events.Message{
				Action: "die",
				Actor: events.Actor{
					ID: "0332dbd79e20aaa",
				},
			},
		},
		true,
	)
	defer watcher.Stop()

	// Check it doesn't get removed while we request meta for the container
	for i := 0; i < 18; i++ {
		watcher.Container("0332dbd79e20")
		assert.Equal(t, 1, len(watcher.Containers()))
		clock.adjust(1 * time.Second)
	}

	// Checks after 10s for the watcher containers to be updated
	l := watcher.bus.Subscribe()
	defer l.Stop()
	clock.adjust(11 * time.Second)
	e := <-l.Events()
	assert.Equal(t, true, e["delete"])
	assert.Equal(t, 0, len(watcher.Containers()))
}

func runWatcher(t *testing.T, kill bool, clock clock, containers [][]types.Container, events []interface{}) *watcher {
	return runWatcherShortID(t, kill, clock, containers, events, false)
}

func runWatcherShortID(t *testing.T, kill bool, clock clock, containers [][]types.Container, events []interface{}, enable bool) *watcher {
	logp.TestingSetup()

	client := &MockClient{
		containers: containers,
		events:     events,
		done:       make(chan interface{}),
	}

	w, err := NewWatcherWithClient(client, 10*time.Second, enable)
	if err != nil {
		t.Fatal(err)
	}
	watcher, ok := w.(*watcher)
	if !ok {
		t.Fatal("'watcher' was supposed to be pointer to the watcher structure")
	}

	watcher.clock = clock

	err = watcher.Start()
	if err != nil {
		t.Fatal(err)
	}

	<-client.done
	if kill {
		watcher.Stop()
		watcher.stopped.Wait()
	}

	return watcher
}

type testClock struct {
	sync.Mutex
	now    time.Time
	sleeps []testClockAfter
}

type testClockAfter struct {
	alarm time.Time
	c     chan time.Time
}

func newTestClock() *testClock {
	return &testClock{
		now: time.Now(),
	}
}

func (c *testClock) Now() time.Time {
	c.Lock()
	defer c.Unlock()
	return c.now
}

func (c *testClock) After(d time.Duration) <-chan time.Time {
	c.Lock()
	defer c.Unlock()
	s := testClockAfter{
		alarm: c.now.Add(d),
		c:     make(chan time.Time),
	}
	c.sleeps = append(c.sleeps, s)
	return s.c
}

func (c *testClock) adjust(d time.Duration) {
	c.Lock()
	defer c.Unlock()

	c.now = c.now.Add(d)
	var sleeps []testClockAfter
	for _, s := range c.sleeps {
		if c.now.After(s.alarm) {
			s.c <- c.now
		} else {
			sleeps = append(sleeps, s)
		}
	}
	c.sleeps = sleeps
}
