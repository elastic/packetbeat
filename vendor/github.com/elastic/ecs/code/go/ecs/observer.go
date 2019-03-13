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

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

// An observer is defined as a special network, security, or application device
// used to detect, observe, or create network, security, or application-related
// events and metrics.
// This could be a custom hardware appliance or a server that has been
// configured to run special network, security, or application software.
// Examples include firewalls, intrusion detection/prevention systems, network
// monitoring sensors, web application firewalls, data loss prevention systems,
// and APM servers. The observer.* fields shall be populated with details of
// the system, if any, that detects, observes and/or creates a network,
// security, or application event or metric. Message queues and ETL components
// used in processing events or metrics are not considered observers in ECS.
type Observer struct {
	// MAC address of the observer
	MAC string `ecs:"mac"`

	// IP address of the observer.
	IP string `ecs:"ip"`

	// Hostname of the observer.
	Hostname string `ecs:"hostname"`

	// observer vendor information.
	Vendor string `ecs:"vendor"`

	// Observer version.
	Version string `ecs:"version"`

	// Observer serial number.
	SerialNumber string `ecs:"serial_number"`

	// The type of the observer the data is coming from.
	// There is no predefined list of observer types. Some examples are
	// `forwarder`, `firewall`, `ids`, `ips`, `proxy`, `poller`, `sensor`, `APM
	// server`.
	Type string `ecs:"type"`
}
