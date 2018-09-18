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

package index_summary

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/mb/parse"
	"github.com/elastic/beats/metricbeat/module/elasticsearch"
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	mb.Registry.MustAddMetricSet(elasticsearch.ModuleName, "index_summary", New,
		mb.WithHostParser(hostParser),
		mb.WithNamespace("elasticsearch.index.summary"),
	)
}

const (
	statsPath = "/_stats"
)

var (
	hostParser = parse.URLHostParserBuilder{
		DefaultScheme: "http",
		PathConfigKey: "path",
	}.Build()
)

// MetricSet type defines all fields of the MetricSet
type MetricSet struct {
	*elasticsearch.MetricSet
}

// New create a new instance of the MetricSet
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The  " + base.FullyQualifiedName() + " metricset is beta")

	// Get the stats from the local node
	ms, err := elasticsearch.NewMetricSet(base, statsPath)
	if err != nil {
		return nil, err
	}
	return &MetricSet{MetricSet: ms}, nil
}

// Fetch gathers stats for each index from the _stats API
func (m *MetricSet) Fetch(r mb.ReporterV2) {
	isMaster, err := elasticsearch.IsMaster(m.HTTP, m.HostData().SanitizedURI+statsPath)
	if err != nil {
		err = errors.Wrap(err, "error determining if connected Elasticsearch node is master")
		r.Error(err)
		m.Log.Error(err)
		return
	}

	// Not master, no event sent
	if !isMaster {
		err = fmt.Errorf("trying to fetch index summary stats from a non-master node")
		r.Error(err)
		m.Log.Error(err)
		return
	}

	content, err := m.HTTP.FetchContent()
	if err != nil {
		r.Error(err)
		m.Log.Error(err)
		return
	}

	info, err := elasticsearch.GetInfo(m.HTTP, m.HostData().SanitizedURI+statsPath)
	if err != nil {
		err = errors.Wrap(err, "failed to get info from Elasticsearch")
		r.Error(err)
		m.Log.Error(err)
		return
	}

	if m.XPack {
		err = eventMappingXPack(r, m, *info, content)
	} else {
		err = eventMapping(r, *info, content)
	}

	if err != nil {
		m.Log.Error(err)
	}
}
