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

package shard

import (
	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/elasticsearch"
)

func init() {
	mb.Registry.MustAddMetricSet(elasticsearch.ModuleName, "shard", New,
		mb.WithHostParser(elasticsearch.HostParser),
		mb.DefaultMetricSet(),
		mb.WithNamespace("elasticsearch.shard"),
	)
}

const (
	statePath = "/_cluster/state/version,nodes,master_node,routing_table"
)

// MetricSet type defines all fields of the MetricSet
type MetricSet struct {
	*elasticsearch.MetricSet
}

// New create a new instance of the MetricSet
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	// Get the stats from the local node
	ms, err := elasticsearch.NewMetricSet(base, statePath)
	if err != nil {
		return nil, err
	}
	return &MetricSet{MetricSet: ms}, nil
}

// Fetch methods implements the data gathering and data conversion to the right format
func (m *MetricSet) Fetch(r mb.ReporterV2) error {
	// If we're talking to a set of ES nodes directly, only collect stats from the master node so
	// we don't collect the same stats from every node and end up duplicating them.
	if m.HostsMode == elasticsearch.HostsModeNode {
		isMaster, err := elasticsearch.IsMaster(m.HTTP, m.HostData().SanitizedURI+statePath)
		if err != nil {
			return errors.Wrap(err, "error determining if connected Elasticsearch node is master")
		}

		// Not master, no event sent
		if !isMaster {
			m.Logger().Debug("trying to fetch shard stats from a non-master node")
			return nil
		}
	}

	content, err := m.HTTP.FetchContent()
	if err != nil {
		return err
	}

	if m.XPack {
		err = eventsMappingXPack(r, m, content)
		if err != nil {
			// Since this is an x-pack code path, we log the error but don't
			// return it. Otherwise it would get reported into `metricbeat-*`
			// indices.
			m.Logger().Error(err)
			return nil
		}
	} else {
		return eventsMapping(r, content)
	}

	return nil
}
