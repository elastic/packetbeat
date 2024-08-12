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

package datastore

import (
	"context"
	"fmt"
	"strings"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/vsphere"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func init() {
	mb.Registry.MustAddMetricSet("vsphere", "datastore", New,
		mb.WithHostParser(vsphere.HostParser),
		mb.DefaultMetricSet(),
	)
}

// MetricSet type defines all fields of the MetricSet.
type MetricSet struct {
	*vsphere.MetricSet
}

// New creates a new instance of the MetricSet.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	ms, err := vsphere.NewMetricSet(base)
	if err != nil {
		return nil, err
	}
	return &MetricSet{ms}, nil
}

type PerformanceMetrics struct {
	DsRead         int64
	DsWrite        int64
	DsIops         int64
	DsReadLatency  int64
	DsWriteLatency int64
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(ctx context.Context, reporter mb.ReporterV2) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, err := govmomi.NewClient(ctx, m.HostURL, m.Insecure)
	if err != nil {
		return fmt.Errorf("error in NewClient: %w", err)
	}
	defer func() {
		if err := client.Logout(ctx); err != nil {
			m.Logger().Debug(fmt.Errorf("error trying to logout from vshphere: %w", err))
		}
	}()

	c := client.Client

	// Create a view of Datastore objects
	mgr := view.NewManager(c)

	v, err := mgr.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"Datastore"}, true)
	if err != nil {
		return fmt.Errorf("error in CreateContainerView: %w", err)
	}

	defer func() {
		if err := v.Destroy(ctx); err != nil {
			m.Logger().Debug(fmt.Errorf("error trying to destroy view from vshphere: %w", err))
		}
	}()

	// Retrieve summary property for all datastores
	var dst []mo.Datastore
	if err = v.Retrieve(ctx, []string{"Datastore"}, []string{"summary", "host", "vm", "overallStatus"}, &dst); err != nil {
		return fmt.Errorf("error in Retrieve: %w", err)
	}

	// Create a performance manager
	perfManager := performance.NewManager(c)

	// Retrieve metric IDs for the specified metric names
	metrics, err := perfManager.CounterInfoByName(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve metrics: %w", err)
	}

	// Define metrics to be collected
	metricNames := []string{
		"datastore.read.average",
		"datastore.write.average",
		"datastore.datastoreIops.average",
		"datastore.totalReadLatency.average",
		"datastore.totalWriteLatency.average",
	}

	// Define reference of structure
	var metricsVar PerformanceMetrics

	// Map metric names to struture	fields
	metricMap := map[string]*int64{
		"datastore.read.average":              &metricsVar.DsRead,
		"datastore.write.average":             &metricsVar.DsWrite,
		"datastore.datastoreIops.average":     &metricsVar.DsIops,
		"datastore.totalReadLatency.average":  &metricsVar.DsReadLatency,
		"datastore.totalWriteLatency.average": &metricsVar.DsWriteLatency,
	}

	var spec types.PerfQuerySpec
	metricIDs := make([]types.PerfMetricId, 0, len(metricNames))

	for _, metricName := range metricNames {
		metric, exists := metrics[metricName]
		if !exists {
			m.Logger().Debug("Metric ", metricName, " not found")
			continue
		}

		metricIDs = append(metricIDs, types.PerfMetricId{
			CounterId: metric.Key,
		})
	}
	spec = types.PerfQuerySpec{
		MetricId:   metricIDs,
		MaxSample:  1,
		IntervalId: 20, // right now we are only grabbing real time metrics from the performance manager
	}

	pc := property.DefaultCollector(c)
	for _, ds := range dst {
		spec.Entity = ds.Reference()

		// Query performance data
		samples, err := perfManager.Query(ctx, []types.PerfQuerySpec{spec})
		if err != nil {
			m.Logger().Debug("Failed to query performance data: %v", err)
			continue
		}
		if len(samples) > 0 {
			results, err := perfManager.ToMetricSeries(ctx, samples)
			if err != nil {
				m.Logger().Debug("Failed to query performance data: %v", err)
			}

			for _, result := range results[0].Value {
				if len(result.Value) > 0 {
					if assignValue, exists := metricMap[result.Name]; exists {
						*assignValue = result.Value[0] // Assign the metric value to the variable
					}
				} else {
					m.Logger().Debug("Metric ", result.Name, ": No result found")
				}
			}
		} else {
			m.Logger().Debug("No samples returned from performance manager")
		}

		outputVmNames := []string{}
		if len(ds.Vm) > 0 {
			VmRefs := []types.ManagedObjectReference{}
			for _, obj := range ds.Vm {
				if obj.Type == "VirtualMachine" {
					VmRefs = append(VmRefs, obj)
				}
			}

			// Retrieve VM names
			var vms []mo.VirtualMachine
			if len(VmRefs) > 0 {
				err := pc.Retrieve(ctx, VmRefs, []string{"name"}, &vms)
				if err != nil {
					m.Logger().Debug("Error retrieving VMs: %v", err)
				}
			}

			for _, vm := range vms {
				name := strings.Replace(vm.Name, ".", "_", -1)
				outputVmNames = append(outputVmNames, name)
			}
		}

		outputHsNames := []string{}
		if len(ds.Host) > 0 {
			hsRefs := []types.ManagedObjectReference{}
			for _, obj := range ds.Host {
				if obj.Key.Type == "HostSystem" {
					hsRefs = append(hsRefs, obj.Key)
				}
			}

			// Retrieve Host names
			var hosts []mo.HostSystem
			if len(hsRefs) > 0 {
				err := pc.Retrieve(ctx, hsRefs, []string{"name"}, &hosts)
				if err != nil {
					m.Logger().Debug("Error retrieving VMs: %v", err)
				}
			}

			for _, host := range hosts {
				name := strings.Replace(host.Name, ".", "_", -1)
				outputHsNames = append(outputHsNames, name)
			}
		}

		reporter.Event(mb.Event{
			MetricSetFields: m.eventMapping(ds, &metricsVar, outputVmNames, outputHsNames),
		})
	}

	return nil
}
