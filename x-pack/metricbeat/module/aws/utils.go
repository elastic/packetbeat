// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package aws

import (
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/cloudwatchiface"
	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
)

// GetStartTimeEndTime function uses durationString to create startTime and endTime for queries.
func GetStartTimeEndTime(durationString string) (startTime time.Time, endTime time.Time, err error) {
	endTime = time.Now()
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return
	}
	startTime = endTime.Add(duration)
	return startTime, endTime, nil
}

// GetListMetricsOutput function gets listMetrics results from cloudwatch per namespace for each region.
// ListMetrics Cloudwatch API is used to list the specified metrics. The returned metrics can be used with GetMetricData
// to obtain statistical data.
func GetListMetricsOutput(namespace string, regionName string, svcCloudwatch cloudwatchiface.CloudWatchAPI) ([]cloudwatch.Metric, error) {
	listMetricsInput := &cloudwatch.ListMetricsInput{Namespace: &namespace}
	reqListMetrics := svcCloudwatch.ListMetricsRequest(listMetricsInput)

	// List metrics of a given namespace for each region
	listMetricsOutput, err := reqListMetrics.Send()
	if err != nil {
		err = errors.Wrap(err, "ListMetricsRequest failed, skipping region "+regionName)
		return nil, err
	}

	if listMetricsOutput.Metrics == nil || len(listMetricsOutput.Metrics) == 0 {
		// No metrics in this region
		return nil, nil
	}
	return listMetricsOutput.Metrics, nil
}

func getMetricDataPerRegion(metricDataQueries []cloudwatch.MetricDataQuery, nextToken *string, svc cloudwatchiface.CloudWatchAPI, startTime time.Time, endTime time.Time) (*cloudwatch.GetMetricDataOutput, error) {
	getMetricDataInput := &cloudwatch.GetMetricDataInput{
		NextToken:         nextToken,
		StartTime:         &startTime,
		EndTime:           &endTime,
		MetricDataQueries: metricDataQueries,
	}

	reqGetMetricData := svc.GetMetricDataRequest(getMetricDataInput)
	getMetricDataOutput, err := reqGetMetricData.Send()
	if err != nil {
		err = errors.Wrap(err, "Error GetMetricDataInput")
		return nil, err
	}
	return getMetricDataOutput, nil
}

// GetMetricDataResults function uses MetricDataQueries to get metric data output.
func GetMetricDataResults(metricDataQueries []cloudwatch.MetricDataQuery, svc cloudwatchiface.CloudWatchAPI, startTime time.Time, endTime time.Time) ([]cloudwatch.MetricDataResult, error) {
	init := true
	getMetricDataOutput := &cloudwatch.GetMetricDataOutput{NextToken: nil}
	for init || getMetricDataOutput.NextToken != nil {
		init = false
		output, err := getMetricDataPerRegion(metricDataQueries, getMetricDataOutput.NextToken, svc, startTime, endTime)
		if err != nil {
			err = errors.Wrap(err, "getMetricDataPerRegion failed")
			return getMetricDataOutput.MetricDataResults, err
		}
		getMetricDataOutput.MetricDataResults = append(getMetricDataOutput.MetricDataResults, output.MetricDataResults...)
	}
	return getMetricDataOutput.MetricDataResults, nil
}

// EventMapping maps data in input to a predefined schema.
func EventMapping(input map[string]interface{}, schema s.Schema) (common.MapStr, error) {
	return schema.Apply(input, s.FailOnRequired)
}

// InArray checks if input val exists in array and if it exists, return the position.
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func FindTimestamp(getMetricDataResults []cloudwatch.MetricDataResult) time.Time {
	timestamp := time.Time{}
	for _, output := range getMetricDataResults {
		// When there are outputs with one timestamp, use this timestamp.
		if output.Timestamps != nil && len(output.Timestamps) == 1 {
			// Use the first timestamp from Timestamps field to collect the latest data.
			timestamp = output.Timestamps[0]
			break
		}
	}

	// When there is no output with one timestamp, use the latest timestamp from timestamp list.
	if timestamp.IsZero() {
		for _, output := range getMetricDataResults {
			// When there are outputs with one timestamp, use this timestamp
			if output.Timestamps != nil && len(output.Timestamps) > 1 {
				// Example Timestamps: [2019-03-11 17:36:00 +0000 UTC,2019-03-11 17:31:00 +0000 UTC]
				timestamp = output.Timestamps[0]
				break
			}
		}
	}

	return timestamp
}
