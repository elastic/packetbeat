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

package mb

import (
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// EventModifier is a function that can modifies an Event. This is typically
// used to apply transformations to an Event as it is converted to a
// beat.Event. An example is AddMetricSetInfo.
type EventModifier func(module, metricset string, event *Event)

// Event contains the data generated by a MetricSet.
type Event struct {
	RootFields      mapstr.M // Fields that will be added to the root of the event.
	ModuleFields    mapstr.M // Fields that will be namespaced under [module].
	MetricSetFields mapstr.M // Fields that will be namespaced under [module].[metricset].

	Index     string        // Index name prefix. If set overwrites the default prefix.
	ID        string        // ID of event. If set, overwrites the default ID.
	Namespace string        // Fully qualified namespace to use for MetricSetFields.
	Timestamp time.Time     // Timestamp when the event data was collected.
	Error     error         // Error that occurred while collecting the event data.
	Host      string        // Host from which the data was collected.
	Service   string        // Service type
	Took      time.Duration // Amount of time it took to collect the event data.
	Period    time.Duration // Period that is set to retrieve the events

	DisableTimeSeries bool // true if the event doesn't contain timeseries data
}

// BeatEvent returns a new beat.Event containing the data this Event. It does
// mutate the underlying data in the Event.
func (e *Event) BeatEvent(module, metricSet string, modifiers ...EventModifier) beat.Event {
	if e.RootFields == nil {
		e.RootFields = mapstr.M{}
	}

	for _, modify := range modifiers {
		modify(module, metricSet, e)
	}

	b := beat.Event{
		Timestamp:  e.Timestamp,
		Fields:     e.RootFields,
		TimeSeries: !e.DisableTimeSeries,
	}

	if len(e.ModuleFields) > 0 {
		b.Fields.Put(module, e.ModuleFields)
		e.ModuleFields = nil
	}

	// If service is not set, falls back to the module name
	if e.Service == "" {
		e.Service = module
	}
	e.RootFields.Put("service.type", e.Service)

	if len(e.MetricSetFields) > 0 {
		switch e.Namespace {
		case ".":
			// Add fields to root.
			b.Fields.DeepUpdate(e.MetricSetFields)
		case "":
			b.Fields.Put(module+"."+metricSet, e.MetricSetFields)
		default:
			b.Fields.Put(e.Namespace, e.MetricSetFields)
		}

		e.MetricSetFields = nil
	}

	// Set index prefix to overwrite default
	if e.Index != "" {
		b.Meta = mapstr.M{"index": e.Index}
	}

	if e.ID != "" {
		b.SetID(e.ID)
	}

	if e.Error != nil {
		b.Fields["error"] = mapstr.M{
			"message": e.Error.Error(),
		}
	}

	return b
}

// AddMetricSetInfo is an EventModifier that adds information about the
// MetricSet that generated the event. It will always add the metricset and
// module names. And it will add the host, period (in milliseconds), and
// duration (round-trip time in nanoseconds) values if they are non-zero
// values.
//
//	{
//	  "event": {
//	    "dataset": "apache.status",
//	    "duration": 115,
//	    "module": "apache"
//	  },
//	  "service": {
//	    "address": "127.0.0.1",
//	  },
//	  "metricset": {
//	    "name": "status",
//	    "period": 10000
//	  }
//	}
func AddMetricSetInfo(module, metricset string, event *Event) {
	if event.Namespace == "" {
		event.Namespace = fmt.Sprintf("%s.%s", module, metricset)
	}

	e := mapstr.M{
		"event": mapstr.M{
			"dataset": event.Namespace,
			"module":  module,
		},
		// TODO: This should only be sent if migration layer is enabled
		"metricset": mapstr.M{
			"name": metricset,
		},
	}
	if event.Host != "" {
		e.Put("service.address", event.Host)
	}
	if event.Took > 0 {
		e.Put("event.duration", event.Took/time.Nanosecond)
	}
	if event.Period > 0 {
		e.Put("metricset.period", event.Period/time.Millisecond)
	}

	if event.RootFields == nil {
		event.RootFields = e
	} else {
		event.RootFields.DeepUpdate(e)
	}
}

// TransformMapStrToEvent transforms a mapstr.M produced by MetricSet
// (like any MetricSet that does not natively produce a mb.Event). It accounts
// for the special key names and routes the data stored under those keys to the
// correct location in the event.
func TransformMapStrToEvent(module string, m mapstr.M, err error) Event {
	var (
		event = Event{RootFields: mapstr.M{}, Error: err}
	)

	for k, v := range m {
		switch k {
		case TimestampKey:
			switch ts := v.(type) {
			case time.Time:
				delete(m, TimestampKey)
				event.Timestamp = ts
			case common.Time:
				delete(m, TimestampKey)
				event.Timestamp = time.Time(ts)
			}
		case ModuleDataKey:
			delete(m, ModuleDataKey)
			event.ModuleFields, _ = tryToMapStr(v)
		case RTTKey:
			delete(m, RTTKey)
			if took, ok := v.(time.Duration); ok {
				event.Took = took
			}
		case NamespaceKey:
			delete(m, NamespaceKey)
			if ns, ok := v.(string); ok {
				// The _namespace value does not include the module name and
				// it is required in the mb.Event.Namespace value.
				event.Namespace = module + "." + ns
			}
		}
	}

	event.MetricSetFields = m
	return event
}

func tryToMapStr(v interface{}) (mapstr.M, bool) {
	switch m := v.(type) {
	case mapstr.M:
		return m, true
	case map[string]interface{}:
		return mapstr.M(m), true
	default:
		return nil, false
	}
}

// PartialMetricsError indicates that metrics are only partially filled.
// This will be removed once we fix the underlying problem.
// See https://github.com/elastic/beats/issues/40542 for more details.
type PartialMetricsError struct {
	Err error
}

func (p PartialMetricsError) Error() string {
	if p.Err == nil {
		return ""
	}
	return p.Err.Error()
}

func (p PartialMetricsError) Unwrap() error {
	return p.Err
}
