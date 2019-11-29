// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package stackdriver

import (
	"context"
	"time"

	"github.com/pkg/errors"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"

	"github.com/elastic/beats/libbeat/common"

	"google.golang.org/api/option"

	"github.com/elastic/beats/x-pack/metricbeat/module/googlecloud"

	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/metricbeat/mb"
)

const (
	METRICSET_NAME = "stackdriver"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet(googlecloud.MODULE_NAME, METRICSET_NAME, New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	config config
}

type config struct {
	Metrics             []string `config:"stackdriver.metrics" validate:"required"`
	Zone                string   `config:"zone" validate:"required"`
	ProjectID           string   `config:"project_id" validate:"required"`
	IncludeLabels       bool     `config:"include_labels"`
	ServiceName         string   `config:"stackdriver.service"  validate:"required"`
	CredentialsFilePath string   `config:"credentials_file_path"`

	opt option.ClientOption
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The gcp '%s' metricset is beta.", METRICSET_NAME)

	m := &MetricSet{BaseMetricSet: base}

	if err := base.Module().UnpackConfig(&m.config); err != nil {
		return nil, err
	}

	m.config.opt = option.WithCredentialsFile(m.config.CredentialsFilePath)

	return m, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(ctx context.Context, reporter mb.ReporterV2) error {
	err := validatePeriodForGCP(m.Module().Config().Period)
	if err != nil {
		return errors.Wrap(err, "invalid time period for GCP")
	}

	//TODO Add credentials check?

	reqs, err := newStackdriverMetricsRequester(m.config, m.Module().Config().Period, m.Logger())
	if err != nil {
		return errors.Wrapf(err, "error trying to do create a request client to GCP project '%s' in zone '%s'", m.config.ProjectID, m.config.Zone)
	}

	responses, err := reqs.Metrics(ctx, m.config.Metrics)
	if err != nil {
		return errors.Wrapf(err, "error trying to get metrics for project '%s' and zone '%s'", m.config.ProjectID, m.config.Zone)
	}

	events, err := m.eventMapping(ctx, responses)
	if err != nil {
		return err
	}

	for _, event := range events {
		reporter.Event(event)
	}

	return nil
}

func (m *MetricSet) eventMapping(ctx context.Context, tss []*monitoringpb.TimeSeries) ([]mb.Event, error) {
	e := newIncomingFieldExtractor(m.Logger())

	var gcpService = googlecloud.NewStackdriverMetadataServiceForTimeSeries(nil)
	var err error
	// Override default metadata service labels are ON and service is known
	if m.config.IncludeLabels {
		if gcpService, err = NewMetadataServiceForConfig(ctx, m.config); err != nil {
			return nil, errors.Wrap(err, "error trying to create metadata service")
		}
	}

	tsGrouped, err := m.timeSeriesGrouped(ctx, gcpService, tss, e)
	if err != nil {
		return nil, errors.Wrap(err, "error trying to group time series data")
	}

	//Create single events for each group of data that matches some common patterns like labels and timestamp
	events := make([]mb.Event, 0)
	for _, groupedEvents := range tsGrouped {
		event := mb.Event{
			Timestamp:  groupedEvents[0].Timestamp,
			RootFields: groupedEvents[0].ECS,
			MetricSetFields: common.MapStr{
				"labels": groupedEvents[0].Labels,
			},
		}

		for _, singleEvent := range groupedEvents {
			event.MetricSetFields.Put(singleEvent.Key, singleEvent.Value)
		}

		events = append(events, event)
	}

	return events, nil
}

// TODO Validate time period?
func validatePeriodForGCP(d time.Duration) (err error) {
	return nil
}
