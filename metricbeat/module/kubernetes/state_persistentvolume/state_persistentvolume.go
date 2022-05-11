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

package state_persistentvolume

import (
	"fmt"

	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/metricbeat/module/kubernetes/util"

	p "github.com/elastic/beats/v7/metricbeat/helper/prometheus"
	"github.com/elastic/beats/v7/metricbeat/mb"
	k8smod "github.com/elastic/beats/v7/metricbeat/module/kubernetes"
)

func init() {
	mb.Registry.MustAddMetricSet("kubernetes", "state_persistentvolume",
		NewPersistentVolumeMetricSet,
		mb.WithHostParser(p.HostParser))
}

// PersistentVolumeMetricSet is a prometheus based MetricSet that looks for
// mb.ModuleDataKey prefixed fields and puts then at the module level
type PersistentVolumeMetricSet struct {
	mb.BaseMetricSet
	prometheus p.Prometheus
	mapping    *p.MetricsMapping
	mod        k8smod.Module
	enricher   util.Enricher
}

// NewPersistentVolumeMetricSet returns a prometheus based metricset for Persistent Volumes
func NewPersistentVolumeMetricSet(base mb.BaseMetricSet) (mb.MetricSet, error) {
	prometheus, err := p.NewPrometheusClient(base)
	if err != nil {
		return nil, err
	}
	mod, ok := base.Module().(k8smod.Module)
	if !ok {
		return nil, fmt.Errorf("must be child of kubernetes module")
	}
	return &PersistentVolumeMetricSet{
		BaseMetricSet: base,
		prometheus:    prometheus,
		mod:           mod,
		enricher:      util.NewResourceMetadataEnricher(base, &kubernetes.PersistentVolume{}, false),
		mapping: &p.MetricsMapping{
			Metrics: map[string]p.MetricMap{
				"kube_persistentvolume_capacity_bytes": p.Metric("capacity.bytes"),
				"kube_persistentvolume_status_phase":   p.LabelMetric("phase", "phase"),
				"kube_persistentvolume_labels": p.ExtendedInfoMetric(
					p.Configuration{
						StoreNonMappedLabels:     true,
						NonMappedLabelsPlacement: mb.ModuleDataKey + ".labels",
						MetricProcessingOptions:  []p.MetricOption{p.OpLabelKeyPrefixRemover("label_")},
					}),
				"kube_persistentvolume_info": p.InfoMetric(),
			},
			Labels: map[string]p.LabelMap{
				"persistentvolume": p.KeyLabel("name"),
				"storageclass":     p.Label("storage_class"),
			},
		},
	}, nil
}

// Fetch prometheus metrics and treats those prefixed by mb.ModuleDataKey as
// module rooted fields at the event that gets reported
func (m *PersistentVolumeMetricSet) Fetch(reporter mb.ReporterV2) {
	m.enricher.Start()

	families, err := m.mod.GetStateMetricsFamilies(m.prometheus)
	if err != nil {
		m.Logger().Error(err)
		reporter.Error(err)
		return
	}
	events, err := m.prometheus.ProcessMetrics(families, m.mapping)
	if err != nil {
		m.Logger().Error(err)
		reporter.Error(err)
		return
	}

	m.enricher.Enrich(events)
	for _, event := range events {

		e, err := util.CreateEvent(event, "kubernetes.persistentvolume")
		if err != nil {
			m.Logger().Error(err)
		}

		if reported := reporter.Event(e); !reported {
			m.Logger().Debug("error trying to emit event")
			return
		}
	}

	return
}

// Close stops this metricset
func (m *PersistentVolumeMetricSet) Close() error {
	m.enricher.Stop()
	return nil
}
