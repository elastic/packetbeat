// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package device_health

import (
	"fmt"
	"reflect"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"

	meraki "github.com/meraki/dashboard-api-go/v3/sdk"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("meraki", "device_health", New)
}

type config struct {
	BaseURL       string   `config:"apiBaseURL"`
	ApiKey        string   `config:"apiKey"`
	DebugMode     string   `config:"apiDebugMode"`
	Organizations []string `config:"organizations"`
	// todo: device filtering?
}

func defaultConfig() *config {
	return &config{
		BaseURL:   "https://api.meraki.com",
		DebugMode: "false",
	}
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	logger        *logp.Logger
	client        *meraki.Client
	organizations []string
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The meraki device_health metricset is beta.")

	logger := logp.NewLogger(base.FullyQualifiedName())

	config := defaultConfig()
	if err := base.Module().UnpackConfig(config); err != nil {
		return nil, err
	}

	logger.Debugf("loaded config: %v", config)
	client, err := meraki.NewClientWithOptions(config.BaseURL, config.ApiKey, config.DebugMode, "Metricbeat Elastic")
	if err != nil {
		logger.Error("creating Meraki dashboard API client failed: %w", err)
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		logger:        logger,
		client:        client,
		organizations: config.Organizations,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(reporter mb.ReporterV2) error {
	for _, org := range m.organizations {
		// some metrics require a 'timespan' parameter; we match this to our
		// collection interval to only collect new metric values
		collectionPeriod := m.BaseMetricSet.Module().Config().Period

		// First we get the list of all devices for this org (and their metadata).
		// Devices are uniquely identified by their serial number, which are used to
		// associate the metrics we collect later with the devices returned here.
		devices, err := getDevices(m.client, org)
		if err != nil {
			return fmt.Errorf("getDevices failed; %w", err)
		}

		// Now we continue to populate the device data structure with health
		// attributes/statuses/metrics etc in the following functions...
		err = getDeviceStatuses(m.client, org, devices)
		if err != nil {
			return fmt.Errorf("getDeviceStatuses failed; %w", err)
		}

		err = getDevicePerformanceScores(m.client, devices)
		if err != nil {
			return fmt.Errorf("getDevicePerformanceScores failed; %w", err)
		}

		err = getDeviceChannelUtilization(m.client, devices, collectionPeriod)
		if err != nil {
			return fmt.Errorf("getDeviceChannelUtilization failed; %w", err)
		}

		err = getDeviceLicenses(m.client, org, devices)
		if err != nil {
			return fmt.Errorf("getDeviceLicenses failed; %w", err)
		}

		err = getDeviceUplinks(m.client, org, devices, collectionPeriod)
		if err != nil {
			return fmt.Errorf("getDeviceUplinks failed; %w", err)
		}

		err = getDeviceSwitchports(m.client, org, devices, collectionPeriod)
		if err != nil {
			return fmt.Errorf("getDeviceSwitchports failed; %w", err)
		}

		// Once we have collected _all_ the data and associated it with the correct device
		// we can report the various device health metrics. These functions are split up
		// in this way primarily to allow better seperation of the code, but also because
		// each function here corresponds to a distinct set of reported metric events
		// i.e. there is one event per device, one event per uplink (but multiple uplinks per device),
		// one event per switchport (but multiple switchports per device), etc.
		reportDeviceMetrics(reporter, org, devices)
		reportUplinkMetrics(reporter, org, devices)
		reportSwitchportMetrics(reporter, org, devices)
	}

	return nil
}

func reportMetricsForOrganization(reporter mb.ReporterV2, organizationID string, metrics ...[]mapstr.M) {
	for _, metricSlice := range metrics {
		for _, metric := range metricSlice {
			event := mb.Event{ModuleFields: mapstr.M{"organization_id": organizationID}}
			if ts, ok := metric["@timestamp"]; ok {
				t, err := time.Parse(time.RFC3339, ts.(string))
				if err == nil {
					// if the timestamp parsing fails, we just fall back to the event time
					// (and leave the additional timestamp in the event for posterity)
					event.Timestamp = t
					delete(metric, "@timestamp")
				}
			}

			for k, v := range metric {
				if !isEmpty(v) {
					event.ModuleFields.Put(k, v)
				}
			}

			reporter.Event(event)
		}
	}
}

func isEmpty(value interface{}) bool {
	// we make use of the fact that all the dashboard API responses utilize
	// pointers for non-string types to filter out empty values from metric events.

	if value == nil {
		return true
	}

	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Ptr {
		return reflect.ValueOf(value).IsNil()
	}

	if t.Kind() == reflect.Slice || t.Kind() == reflect.String {
		return reflect.ValueOf(value).Len() == 0
	}

	return false
}
