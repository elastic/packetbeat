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

// +build windows

package perfmon

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"
)

// CounterConfig for perfmon counters.
type CounterConfig struct {
	InstanceLabel    string `config:"instance_label"    validate:"required"`
	InstanceName     string `config:"instance_name"`
	MeasurementLabel string `config:"measurement_label" validate:"required"`
	Query            string `config:"query"             validate:"required"`
	Format           string `config:"format"`
}

// Config for the windows perfmon metricset.
type Config struct {
	IgnoreNECounters bool            `config:"perfmon.ignore_non_existent_counters"`
	CounterConfig    []CounterConfig `config:"perfmon.counters" validate:"required"`
}

func init() {
	mb.Registry.MustAddMetricSet("windows", "perfmon", New)
}

type MetricSet struct {
	mb.BaseMetricSet
	reader *PerfmonReader
	log    *logp.Logger
}

// New create a new instance of the MetricSet.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The perfmon metricset is beta")

	var config Config
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	for _, value := range config.CounterConfig {
		form := strings.ToLower(value.Format)
		switch form {
		case "":
			value.Format = "float"
		case "float", "long":
		default:
			return nil, errors.Errorf("initialization failed: format '%s' "+
				"for counter '%s' is invalid (must be float or long)",
				value.Format, value.InstanceLabel)
		}

	}

	reader, err := NewPerfmonReader(config)
	if err != nil {
		return nil, errors.Wrap(err, "initialization of reader failed")
	}

	return &MetricSet{
		BaseMetricSet: base,
		reader:        reader,
		log:           logp.NewLogger("perfmon"),
	}, nil
}

func (m *MetricSet) Fetch(report mb.ReporterV2) {
	events, err := m.reader.Read()
	if err != nil {
		m.log.Debugw("Failed reading counters", "error", err)
		err = errors.Wrap(err, "failed reading counters")
		report.Error(err)
	}

	for _, event := range events {
		report.Event(event)
	}
}
