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

package cpu

import (
	"fmt"

	"github.com/elastic/beats/v7/metricbeat/module/kubernetes/util"

	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/metricbeat/helper/prometheus"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

// Metricset for apiserver is a prometheus based metricset
type metricset struct {
	mb.BaseMetricSet
	prometheusClient           prometheus.Prometheus
	prometheusMappings         *prometheus.MetricsMapping
	preSystemCpuUsage          float64
	preContainerCpuTotalUsage  map[string]float64
	preContainerCpuKernelUsage map[string]float64
	preContainerCpuUserUsage   map[string]float64
}

var _ mb.ReportingMetricSetV2Error = (*metricset)(nil)

// getMetricsetFactory as required by` mb.Registry.MustAddMetricSet`
func getMetricsetFactory(prometheusMappings *prometheus.MetricsMapping) mb.MetricSetFactory {
	return func(base mb.BaseMetricSet) (mb.MetricSet, error) {
		pc, err := prometheus.NewPrometheusClient(base)
		if err != nil {
			return nil, err
		}
		return &metricset{
			BaseMetricSet:              base,
			prometheusClient:           pc,
			prometheusMappings:         prometheusMappings,
			preSystemCpuUsage:          0.0,
			preContainerCpuTotalUsage:  map[string]float64{},
			preContainerCpuKernelUsage: map[string]float64{},
			preContainerCpuUserUsage:   map[string]float64{},
		}, nil
	}
}

// Fetch gathers information from the containerd and reports events with this information.
func (m *metricset) Fetch(reporter mb.ReporterV2) error {
	events, err := m.prometheusClient.GetProcessedMetrics(m.prometheusMappings)
	if err != nil {
		return errors.Wrap(err, "error getting metrics")
	}
	var systemTotalNs int64
	perContainerCpus := make(map[string]int)
	elToDel := -1
	for i, event := range events {
		systemTotalSeconds, err := event.GetValue("system.total")
		if err == nil {
			systemTotalNs = systemTotalSeconds.(int64) * 1e9
			elToDel = i
		}
		if _, err = event.GetValue("cpu"); err == nil {
			// calculate cpus used by each container
			setContCpus(event, perContainerCpus)
		}

	}
	// Remove event containing system.total
	if elToDel != -1 {
		events[elToDel] = events[len(events)-1] // Copy last element to index i.
		events[len(events)-1] = common.MapStr{} // Erase last element (write empty value).
		events = events[:len(events)-1]
	}
	//m.Logger().Infof("ContainerCpus are %+v", perContainerCpus)

	for _, event := range events {
		// setting ECS container.id
		containerFields := common.MapStr{}
		var cID string
		if containerID, ok := event["id"]; ok {
			cID = (containerID).(string)
			containerFields.Put("id", cID)
			event.Delete("id")
		}
		e, err := util.CreateEvent(event, "containerd.cpu")
		if err != nil {
			m.Logger().Error(err)
		}

		if len(containerFields) > 0 {
			if e.RootFields != nil {
				e.RootFields.DeepUpdate(common.MapStr{
					"container": containerFields,
				})
			} else {
				e.RootFields = common.MapStr{
					"container": containerFields,
				}
			}
		}
		contCpus, ok := perContainerCpus[cID]
		if !ok {
			contCpus = 1
		}
		// calculate system usage delta
		systemUsageDelta := float64(systemTotalNs) - m.preSystemCpuUsage

		// Calculate cpu total usage percentage
		cpuUsageTotal, err := event.GetValue("usage.total.ns")
		if err == nil {
			cpuUsageTotalPct := calcCpuTotalUsagePct(cpuUsageTotal.(float64), systemUsageDelta,
				float64(contCpus), cID, m.preContainerCpuTotalUsage)
			m.Logger().Infof("cpuUsageTotalPct for %+v is %+v", cID, cpuUsageTotalPct)
			e.MetricSetFields.Put("usage.total.pct", cpuUsageTotalPct)
			// Update values
			m.preContainerCpuTotalUsage[cID] = cpuUsageTotal.(float64)
		}

		// Calculate cpu kernel usage percentage
		cpuUsageKernel, err := event.GetValue("usage.kernel.ns")
		if err == nil {
			cpuUsageKernelPct := calcCpuKernelUsagePct(cpuUsageKernel.(float64), systemUsageDelta,
				float64(contCpus), cID, m.preContainerCpuKernelUsage)
			m.Logger().Infof("cpuUsageKernelPct for %+v is %+v", cID, cpuUsageKernelPct)
			e.MetricSetFields.Put("usage.kernel.pct", cpuUsageKernelPct)
			// Update values
			m.preContainerCpuKernelUsage[cID] = cpuUsageKernel.(float64)
		}

		// Calculate cpu user usage percentage
		cpuUsageUser, err := event.GetValue("usage.user.ns")
		if err == nil {
			cpuUsageUserPct := calcCpuUserUsagePct(cpuUsageUser.(float64), systemUsageDelta,
				float64(contCpus), cID, m.preContainerCpuUserUsage)
			m.Logger().Infof("cpuUsageUserPct for %+v is %+v", cID, cpuUsageUserPct)
			e.MetricSetFields.Put("usage.user.pct", cpuUsageUserPct)
			// Update values
			m.preContainerCpuUserUsage[cID] = cpuUsageUser.(float64)
		}

		if cpuId, err := event.GetValue("cpu"); err == nil {
			perCpuNs, err := event.GetValue("usage.percpu.ns")
			if err == nil {
				key := fmt.Sprintf("usage.cpu.%s.ns", cpuId)
				e.MetricSetFields.Put(key, perCpuNs)
				e.MetricSetFields.Delete("cpu")
				e.MetricSetFields.Delete("usage.percpu.ns")
			}
		}

		if reported := reporter.Event(e); !reported {
			return nil
		}
	}
	m.preSystemCpuUsage = float64(systemTotalNs)
	return nil
}

func setContCpus(event common.MapStr, perContainerCpus map[string]int) {
	val, err := event.GetValue("id")
	if err != nil {
		return
	}
	cid := val.(string)
	_, err = event.GetValue("usage.percpu.ns")
	if err != nil {
		return
	}
	if _, ok := perContainerCpus[cid]; ok {
		perContainerCpus[cid] += 1
	} else {
		perContainerCpus[cid] = 1
	}
	//if percpu.(float64) != 0.0 {
	//	if _, ok := perContainerCpus[cid]; ok {
	//		perContainerCpus[cid] += 1
	//	} else {
	//		perContainerCpus[cid] = 1
	//	}
	//}
}

func calcCpuTotalUsagePct(cpuUsageTotal, systemUsageDelta, contCpus float64,
	cid string, preContainerCpuTotalUsage map[string]float64) float64 {
	var contUsageDelta, cpuUsageTotalPct float64
	if cpuPreval, ok := preContainerCpuTotalUsage[cid]; ok {
		contUsageDelta = cpuUsageTotal - cpuPreval
	} else {
		contUsageDelta = cpuUsageTotal
	}
	if contUsageDelta == 0.0 || systemUsageDelta == 0.0 {
		cpuUsageTotalPct = 0.0
	} else {
		// normalize percentage with cpus used per container
		cpuUsageTotalPct = (contUsageDelta / systemUsageDelta) / contCpus
	}
	return cpuUsageTotalPct
}

func calcCpuKernelUsagePct(cpuUsageKernel, systemUsageDelta, contCpus float64,
	cid string, preContainerCpuKernelUsage map[string]float64) float64 {
	var contUsageDelta, cpuUsageKernelPct float64
	if cpuPreval, ok := preContainerCpuKernelUsage[cid]; ok {
		contUsageDelta = cpuUsageKernel - cpuPreval
	} else {
		contUsageDelta = cpuUsageKernel
	}
	if contUsageDelta == 0.0 || systemUsageDelta == 0.0 {
		cpuUsageKernelPct = 0.0
	} else {
		// normalize percentage with cpus used per container
		cpuUsageKernelPct = (contUsageDelta / systemUsageDelta) / contCpus
	}
	return cpuUsageKernelPct
}

func calcCpuUserUsagePct(cpuUsageUser, systemUsageDelta, contCpus float64,
	cid string, preContainerCpuUserUsage map[string]float64) float64 {
	var contUsageDelta, cpuUsageUserPct float64
	if cpuPreval, ok := preContainerCpuUserUsage[cid]; ok {
		contUsageDelta = cpuUsageUser - cpuPreval
	} else {
		contUsageDelta = cpuUsageUser
	}
	if contUsageDelta == 0.0 || systemUsageDelta == 0.0 {
		cpuUsageUserPct = 0.0
	} else {
		// normalize percentage with cpus used per container
		cpuUsageUserPct = (contUsageDelta / systemUsageDelta) / contCpus
	}
	return cpuUsageUserPct
}
