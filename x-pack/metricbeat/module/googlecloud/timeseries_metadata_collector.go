// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package googlecloud

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"

	"github.com/elastic/beats/libbeat/common"
)

// NewStackdriverCollectorInputData returns a ready to use MetadataCollectorInputData to be sent to Metadata collectors
func NewStackdriverCollectorInputData(ts *monitoringpb.TimeSeries, projectID, zone string) *MetadataCollectorInputData {
	return &MetadataCollectorInputData{
		TimeSeries: ts,
		ProjectID:  projectID,
		Zone:       zone,
	}
}

// NewStackdriverMetadataServiceForTimeSeries apart from having a long name takes a time series object to return the
// Stackdriver canonical Metadata extractor
func NewStackdriverMetadataServiceForTimeSeries(ts *monitoringpb.TimeSeries) MetadataService {
	return &StackdriverTimeSeriesMetadataCollector{
		timeSeries: ts,
	}
}

// StackdriverTimeSeriesMetadataCollector is the implementation of MetadataCollecter to collect metrics from Stackdriver
// common TimeSeries objects
type StackdriverTimeSeriesMetadataCollector struct {
	timeSeries *monitoringpb.TimeSeries
}

// Metadata parses a Timeseries object to return its metadata divided into "unkonwn" (first object) and ECS (second
// object https://www.elastic.co/guide/en/ecs/master/index.html)
func (s *StackdriverTimeSeriesMetadataCollector) Metadata(ctx context.Context, in *monitoringpb.TimeSeries) (common.MapStr, common.MapStr, error) {
	m := common.MapStr{}

	var availabilityZone, accountID string

	if in.Resource != nil && in.Resource.Labels != nil {
		availabilityZone = in.Resource.Labels[JSON_PATH_ECS_AVAILABILITY_ZONE]
		accountID = in.Resource.Labels[JSON_PATH_ECS_ACCOUNT_ID]
	}

	ecs := common.MapStr{
		ECS_CLOUD: common.MapStr{
			ECS_CLOUD_AVAILABILITY_ZONE: availabilityZone,
			ECS_CLOUD_ACCOUNT: common.MapStr{
				ECS_CLOUD_ACCOUNT_ID: accountID,
			},
			ECS_CLOUD_PROVIDER: "googlecloud",
		},
	}

	//Remove keys from resource that refers to ECS fields

	if s.timeSeries != nil || s.timeSeries.Metric != nil {
		metrics := make(map[string]interface{})
		// common.Mapstr seems to not work as expected when deleting keys so I have to iterate over all results to add
		// the ones I want
		for k, v := range s.timeSeries.Metric.Labels {
			if k == JSON_PATH_ECS_INSTANCE_NAME {
				continue
			}

			metrics[k] = v
		}

		//Do not write metrics labels if it's finally empty
		for k, v := range metrics {
			m.Put(LABEL_METRICS+"."+k, v)
		}
	}

	if s.timeSeries.Resource != nil {
		resources := make(map[string]interface{})
		// common.Mapstr seems to not work as expected when deleting keys so I have to iterate over all results to add
		// the ones I want
		for k, v := range s.timeSeries.Resource.Labels {
			if k == JSON_PATH_ECS_AVAILABILITY_ZONE || k == JSON_PATH_ECS_INSTANCE_ID || k == JSON_PATH_ECS_ACCOUNT_ID {
				continue
			}

			resources[k] = v
		}

		//Do not write resources labels if it's finally empty
		for k, v := range resources {
			m.Put(LABEL_RESOURCE+"."+k, v)
		}
	}

	if s.timeSeries.Metadata != nil {
		m.Put(LABEL_SYSTEM, s.timeSeries.Metadata.SystemLabels)
		m.Put(LABEL_USER_METADATA, s.timeSeries.Metadata.UserLabels)
	}

	return m, ecs, nil
}

// ID returns a unique generated ID for an event when no service is implemented to get a "better" ID.`El trickerionEl trickerion
func (s *StackdriverTimeSeriesMetadataCollector) ID(ctx context.Context, in *MetadataCollectorInputData) (string, error) {
	m := common.MapStr{
		KEY_TIMESTAMP: in.Timestamp.UnixNano(),
	}

	if s.timeSeries == nil {
		return "", errors.New("no data found on the time series")
	}

	if s.timeSeries.Metric != nil {
		if s.timeSeries.Metric.Type != "" {
			m.Put("metric.type", s.timeSeries.Metric.Type)
		}

		if s.timeSeries.Metric.Labels != nil {
			m.Put("metric.labels", s.timeSeries.Metric.Labels)
		}
	}

	if s.timeSeries.Resource != nil {
		if s.timeSeries.Resource.Type != "" {
			m.Put("resource.type", s.timeSeries.Resource.Type)
		}

		if s.timeSeries.Resource.Labels != nil {
			m.Put("resource.labels", s.timeSeries.Resource.Labels)
		}
	}

	if s.timeSeries.Metadata != nil {
		if s.timeSeries.Metadata.SystemLabels != nil {
			m.Put("metadata.system.labels", s.timeSeries.Metadata.SystemLabels)
		}
		if s.timeSeries.Metadata.UserLabels != nil {
			m.Put("metadata.user.labels", s.timeSeries.Metadata.UserLabels)
		}
	}

	return m.String(), nil
}

func (s *StackdriverTimeSeriesMetadataCollector) getTimestamp(p *monitoringpb.Point) (t time.Time, err error) {
	// Don't add point intervals that can't be "stated" at some timestamp.
	if p != nil && p.Interval != nil {
		if t, err = ptypes.Timestamp(p.Interval.StartTime); err != nil {
			return time.Time{}, errors.Errorf("error trying to parse timestamp '%#v' from metric\n", p.Interval.StartTime)
		}
	}

	return time.Time{}, errors.New("error trying to extract the timestamp from the point data")
}
