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

package openmetricscollector

import (
	"math"
	"strconv"

	p "github.com/elastic/beats/v7/metricbeat/helper/openmetrics"
	"github.com/prometheus/prometheus/pkg/textparse"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/metricbeat/helper/labelhash"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

// OpenMetricEvent stores a set of one or more metrics with the same labels
type OpenMetricEvent struct {
	Data      common.MapStr
	Labels    common.MapStr
	Help      string
	Type      textparse.MetricType
	Unit      string
	Created   int64
	Exemplars common.MapStr
}

// LabelsHash returns a repeatable string that is unique for the set of labels in this event
func (p *OpenMetricEvent) LabelsHash() string {
	return labelhash.LabelHash(p.Labels)
}

// DefaultOpenMetricEventsGeneratorFactory returns the default OpenMetrics events generator
func DefaultOpenMetricsEventsGeneratorFactory(ms mb.BaseMetricSet) (OpenMetricsEventsGenerator, error) {
	return &openmetricEventGenerator{}, nil
}

type openmetricEventGenerator struct{}

func (p *openmetricEventGenerator) Start() {}
func (p *openmetricEventGenerator) Stop()  {}

// DefaultopenmetricEventsGenerator stores all OpenMetrics metrics using
// only double field type in Elasticsearch.
func (p *openmetricEventGenerator) GenerateOpenMetricsEvents(mf *p.OpenMetricFamily) []OpenMetricEvent {
	var events []OpenMetricEvent

	name := *mf.Name
	metrics := mf.Metric
	// info := *mf.Info
	help := ""
	unit := ""
	if mf.Help != nil {
		help = *mf.Help
	}
	if mf.Unit != nil {
		unit = *mf.Unit
	}

	for _, metric := range metrics {
		labels := common.MapStr{}
		mn := metric.GetName()

		if len(metric.Label) != 0 {
			for _, label := range metric.Label {
				if label.Name != "" && label.Value != "" {
					labels[label.Name] = label.Value
				}
			}
		}

		exemplars := common.MapStr{}
		if metric.Exemplar != nil {
			exemplars = common.MapStr{*mn: metric.Exemplar.Value}
			if metric.Exemplar.HasTs {
				exemplars.Put("timestamp", metric.Exemplar.Ts)
			}
			for _, label := range metric.Exemplar.Labels {
				if label.Name != "" && label.Value != "" {
					exemplars.Put("labels."+label.Name, label.Value)
				}
			}
		}

		created := int64(0)
		if metric.CreatedMs != nil {
			created = *metric.CreatedMs
		}

		// TODO: where is timestamp used?
		//timestamp := int64(0)
		//if metric.TimestampMs != nil {
		//	timestamp = *metric.TimestampMs
		//}

		counter := metric.GetCounter()
		if counter != nil {
			if !math.IsNaN(counter.GetValue()) && !math.IsInf(counter.GetValue(), 0) {
				events = append(events, OpenMetricEvent{
					Created: created,
					Type:    textparse.MetricTypeCounter,
					Help:    help,
					Unit:    unit,
					Data: common.MapStr{
						"metrics": common.MapStr{
							*mn: counter.GetValue(),
						},
					},
					Labels:    labels,
					Exemplars: exemplars,
				})
			}
		}

		gauge := metric.GetGauge()
		if gauge != nil {
			if !math.IsNaN(gauge.GetValue()) && !math.IsInf(gauge.GetValue(), 0) {
				events = append(events, OpenMetricEvent{
					Type: textparse.MetricTypeGauge,
					Help: help,
					Unit: unit,
					Data: common.MapStr{
						"metrics": common.MapStr{
							name: gauge.GetValue(),
						},
					},
					Labels: labels,
				})
			}
		}

		info := metric.GetInfo()
		if info != nil {
			if info.HasValidValue() {
				events = append(events, OpenMetricEvent{
					Type: textparse.MetricTypeInfo,
					Data: common.MapStr{
						"metrics": common.MapStr{
							name: info.GetValue(),
						},
					},
					Labels: labels,
				})
			}
		}

		stateset := metric.GetStateset()
		if stateset != nil {
			if stateset.HasValidValue() {
				events = append(events, OpenMetricEvent{
					Type: textparse.MetricTypeStateset,
					Data: common.MapStr{
						"metrics": common.MapStr{
							name: stateset.GetValue(),
						},
					},
					Labels: labels,
				})
			}
		}

		summary := metric.GetSummary()
		if summary != nil {
			if !math.IsNaN(summary.GetSampleSum()) && !math.IsInf(summary.GetSampleSum(), 0) {
				events = append(events, OpenMetricEvent{
					Created: created,
					Type:    textparse.MetricTypeSummary,
					Help:    help,
					Unit:    unit,
					Data: common.MapStr{
						"metrics": common.MapStr{
							name + "_sum":   summary.GetSampleSum(),
							name + "_count": summary.GetSampleCount(),
						},
					},
					Labels: labels,
				})
			}

			for _, quantile := range summary.GetQuantile() {
				if math.IsNaN(quantile.GetValue()) || math.IsInf(quantile.GetValue(), 0) {
					continue
				}

				quantileLabels := labels.Clone()
				quantileLabels["quantile"] = strconv.FormatFloat(quantile.GetQuantile(), 'f', -1, 64)
				events = append(events, OpenMetricEvent{
					Data: common.MapStr{
						"metrics": common.MapStr{
							name: quantile.GetValue(),
						},
					},
					Labels: quantileLabels,
				})
			}
		}

		histogram := metric.GetHistogram()
		if histogram != nil {
			if !math.IsNaN(histogram.GetSampleSum()) && !math.IsInf(histogram.GetSampleSum(), 0) {
				var sum = "_sum"
				var count = "_count"
				var typ = textparse.MetricTypeHistogram
				if histogram.IsGaugeHistogram {
					sum = "_gsum"
					count = "_gcount"
					typ = textparse.MetricTypeGaugeHistogram
				}

				events = append(events, OpenMetricEvent{
					Created: created,
					Type:    typ,
					Help:    help,
					Unit:    unit,
					Data: common.MapStr{
						"metrics": common.MapStr{
							name + sum:   histogram.GetSampleSum(),
							name + count: histogram.GetSampleCount(),
						},
					},
					Labels: labels,
				})
			}

			for _, bucket := range histogram.GetBucket() {
				if bucket.GetCumulativeCount() == uint64(math.NaN()) || bucket.GetCumulativeCount() == uint64(math.Inf(0)) {
					continue
				}

				if bucket.Exemplar != nil {
					exemplars = common.MapStr{name: bucket.Exemplar.Value}
					if bucket.Exemplar.HasTs {
						exemplars.Put("timestamp", bucket.Exemplar.Ts)
					}
					for _, label := range bucket.Exemplar.Labels {
						if label.Name != "" && label.Value != "" {
							exemplars.Put("labels."+label.Name, label.Value)
						}
					}
				}

				bucketLabels := labels.Clone()
				bucketLabels["le"] = strconv.FormatFloat(bucket.GetUpperBound(), 'f', -1, 64)

				events = append(events, OpenMetricEvent{
					Data: common.MapStr{
						"metrics": common.MapStr{
							name + "_bucket": bucket.GetCumulativeCount(),
						},
					},
					Labels:    bucketLabels,
					Exemplars: exemplars,
				})
			}
		}

		unknown := metric.GetUnknown()
		if unknown != nil {
			if !math.IsNaN(unknown.GetValue()) && !math.IsInf(unknown.GetValue(), 0) {
				events = append(events, OpenMetricEvent{
					Type: textparse.MetricTypeUnknown,
					Help: help,
					Unit: unit,
					Data: common.MapStr{
						"metrics": common.MapStr{
							name: unknown.GetValue(),
						},
					},
					Labels: labels,
				})
			}
		}
	}
	return events
}
