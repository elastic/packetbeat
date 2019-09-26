// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package httpjson

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"golang.org/x/sync/errgroup"

	"github.com/elastic/beats/filebeat/channel"
	"github.com/elastic/beats/filebeat/input"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

var (
	once sync.Once
	url  string
)

func testSetup(t *testing.T) {
	t.Helper()
	once.Do(func() {
		logp.TestingSetup()
	})
}

func runTest(t *testing.T, m map[string]interface{}, run func(input *httpjsonInput, out *stubOutleter, t *testing.T)) {
	// Setup httpbin environment
	testSetup(t)
	// Create test http server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			req, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
			var m interface{}
			err = json.Unmarshal(req, &m)
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(req)
			}
		} else {
			message := map[string]interface{}{
				"hello": "world",
				"embedded": map[string]string{
					"hello": "world",
				},
			}
			b, _ := json.Marshal(message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(b)
		}
	}))
	defer ts.Close()
	m["url"] = ts.URL
	cfg := common.MustNewConfigFrom(m)
	// Simulate input.Context from Filebeat input runner.
	inputCtx := newInputContext()
	defer close(inputCtx.Done)

	// Stub outlet for receiving events generated by the input.
	eventOutlet := newStubOutlet()
	defer eventOutlet.Close()

	connector := channel.ConnectorFunc(func(_ *common.Config, _ beat.ClientConfig) (channel.Outleter, error) {
		return eventOutlet, nil
	})

	in, err := NewInput(cfg, connector, inputCtx)
	if err != nil {
		t.Fatal(err)
	}
	input := in.(*httpjsonInput)
	defer input.Stop()

	run(input, eventOutlet, t)
}

func newInputContext() input.Context {
	return input.Context{
		Done: make(chan struct{}),
	}
}

type stubOutleter struct {
	sync.Mutex
	cond   *sync.Cond
	done   bool
	Events []beat.Event
}

func newStubOutlet() *stubOutleter {
	o := &stubOutleter{}
	o.cond = sync.NewCond(o)
	return o
}

func (o *stubOutleter) waitForEvents(numEvents int) ([]beat.Event, bool) {
	o.Lock()
	defer o.Unlock()

	for len(o.Events) < numEvents && !o.done {
		o.cond.Wait()
	}

	size := numEvents
	if size >= len(o.Events) {
		size = len(o.Events)
	}

	out := make([]beat.Event, size)
	copy(out, o.Events)
	return out, len(out) == numEvents
}

func (o *stubOutleter) Close() error {
	o.Lock()
	defer o.Unlock()
	o.done = true
	return nil
}

func (o *stubOutleter) Done() <-chan struct{} { return nil }

func (o *stubOutleter) OnEvent(event beat.Event) bool {
	o.Lock()
	defer o.Unlock()
	o.Events = append(o.Events, event)
	o.cond.Broadcast()
	return !o.done
}

// --- Test Cases

func TestGET(t *testing.T) {
	m := map[string]interface{}{
		"http_method": "GET",
		"interval":    0,
	}
	runTest(t, m, func(input *httpjsonInput, out *stubOutleter, t *testing.T) {
		group, _ := errgroup.WithContext(context.Background())
		group.Go(input.run)

		events, ok := out.waitForEvents(1)
		if !ok {
			t.Fatalf("Expected 1 events, but got %d.", len(events))
		}
		input.Stop()

		if err := group.Wait(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestPOST(t *testing.T) {
	m := map[string]interface{}{
		"http_method":       "POST",
		"http_request_body": map[string]interface{}{"test": "abc", "testNested": map[string]interface{}{"testNested1": 123}},
		"interval":          0,
	}
	runTest(t, m, func(input *httpjsonInput, out *stubOutleter, t *testing.T) {
		group, _ := errgroup.WithContext(context.Background())
		group.Go(input.run)

		events, ok := out.waitForEvents(1)
		if !ok {
			t.Fatalf("Expected 1 events, but got %d.", len(events))
		}
		input.Stop()

		if err := group.Wait(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestRunStop(t *testing.T) {
	m := map[string]interface{}{
		"http_method": "GET",
		"interval":    0,
	}
	runTest(t, m, func(input *httpjsonInput, out *stubOutleter, t *testing.T) {
		input.Run()
		input.Stop()
		input.Run()
		input.Stop()
	})
}
