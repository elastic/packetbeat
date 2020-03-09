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

/*
Package beater provides the implementation of the libbeat Beater interface for
Metricbeat and functions for running Metricbeat Modules on their own.

Metricbeat collects metrics from operating systems and services. The code for
gathering metrics from a particular service is organized into a logical grouping
called a Module. Each Module has one or more MetricSet implementations which
do the work of collecting a particular set of metrics from the service.

The public interfaces used in implementing Modules and MetricSets are defined in
the github.com/elastic/beats/v7/metricbeat/mb package.

Event Format

Each event generated by Metricbeat has the same general structure. The example
event below was generated by a MetricSet named "cpu" in the "system" Module.

	{
		"@timestamp": "2016-05-23T08:05:34.853Z",
		"beat": {
			"hostname": "host.example.com",
			"name": "host.example.com"
		},
		"metricset": {
			"host": "localhost",
			"module": "system",
			"name": "cpu",
			"rtt": 115
		},
		"system": {
			"cpu": {
				"idle": {
					"pct": 0.852,
					"ticks": 44421033
				},
				"iowait": {
					"pct": 0,
					"ticks": 159735
				},
				"irq": {
					"pct": 0,
					"ticks": 0
				},
				"nice": {
					"pct": 0,
					"ticks": 0
				},
				"softirq": {
					"pct": 0,
					"ticks": 14070
				},
				"steal": {
					"pct": 0,
					"ticks": 0
				},
				"system": {
					"pct": 0.0408,
					"ticks": 305704
				},
				"user": {
					"pct": 0.1071,
					"ticks": 841974
				}
			}
		},
		"type": "metricsets"
	}

All events are stored in one index called metricbeat by default. Each
MetricSet's data format is potentially unique so the MetricSet data is added to
event as a dictionary under a key that is unique to the MetricSet. The key
is constructed from the Module name and MetricSet name to ensure uniqueness.
All documents are stored under the same type called "metricsets".
*/
package beater
