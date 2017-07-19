package testing

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/outputs/codec/json"
	"github.com/elastic/beats/libbeat/publisher/beat"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/mb/module"
)

var (
	// Use `go test -data` to update files.
	dataFlag = flag.Bool("data", false, "Write updated data.json files")
)

// WriteEvent fetches a single event writes the output to a ./_meta/data.json
// file.
func WriteEvent(f mb.EventFetcher, t testing.TB) error {
	if !*dataFlag {
		t.Skip("skip data generation tests")
	}

	event, err := f.Fetch()
	if err != nil {
		return err
	}

	fullEvent := CreateFullEvent(f, event)
	WriteEventToDataJSON(t, fullEvent)
	return nil
}

// WriteEvents fetches events and writes the first event to a ./_meta/data.json
// file.
func WriteEvents(f mb.EventsFetcher, t testing.TB) error {
	if !*dataFlag {
		t.Skip("skip data generation tests")
	}

	events, err := f.Fetch()
	if err != nil {
		return err
	}

	if len(events) == 0 {
		return fmt.Errorf("no events were generated")
	}

	fullEvent := CreateFullEvent(f, events[0])
	WriteEventToDataJSON(t, fullEvent)
	return nil
}

// CreateFullEvent builds a full event given the data generated by a MetricSet.
// This simulates the output of Metricbeat as if it were
// 2016-05-23T08:05:34.853Z and the hostname is host.example.com.
func CreateFullEvent(ms mb.MetricSet, metricSetData common.MapStr) beat.Event {
	startTime, err := time.Parse(time.RFC3339Nano, "2016-05-23T08:05:34.853Z")
	if err != nil {
		panic(err)
	}

	build := module.EventBuilder{
		ModuleName:    ms.Module().Name(),
		MetricSetName: ms.Name(),
		Host:          ms.Host(),
		StartTime:     startTime,
		FetchDuration: 115 * time.Microsecond,
		Event:         metricSetData,
	}

	fullEvent, err := build.Build()
	if err != nil {
		panic(err)
	}
	fullEvent.Fields["beat"] = common.MapStr{
		"name":     "host.example.com",
		"hostname": "host.example.com",
	}

	// Delete meta data as not needed for the event output here.
	delete(fullEvent.Fields, common.EventMetadataKey)

	return fullEvent
}

// WriteEventToDataJSON writes the given event as "pretty" JSON to
// a ./_meta/data.json file. If the -data CLI flag is unset or false then the
// method is a no-op.
func WriteEventToDataJSON(t testing.TB, fullEvent beat.Event) {
	if !*dataFlag {
		return
	}

	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// use json output codec to encode event to json
	output, err := json.New(true).Encode("noindex", &fullEvent)
	if err != nil {
		t.Fatal(err)
	}

	if err = ioutil.WriteFile(path+"/_meta/data.json", output, 0644); err != nil {
		t.Fatal(err)
	}
}
