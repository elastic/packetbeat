package node_stats

import (
	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/module/elasticsearch"
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	mb.Registry.MustAddMetricSet("elasticsearch", "node_stats", New,
		mb.WithHostParser(elasticsearch.HostParser),
		mb.DefaultMetricSet(),
		mb.WithNamespace("elasticsearch.node.stats"),
	)
}

// MetricSet type defines all fields of the MetricSet
type MetricSet struct {
	*elasticsearch.MetricSet
}

// New create a new instance of the MetricSet
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The elasticsearch node_stats metricset is beta")

	// Get the stats from the local node
	ms, err := elasticsearch.NewMetricSet(base, "/_nodes/_local/stats")
	if err != nil {
		return nil, err
	}
	return &MetricSet{MetricSet: ms}, nil
}

// Fetch methods implements the data gathering and data conversion to the right format
func (m *MetricSet) Fetch(r mb.ReporterV2) {
	content, err := m.HTTP.FetchContent()
	if err != nil {
		r.Error(err)
		return
	}

	if m.MetricSet.XPack {
		eventsMappingXPack(r, m, content)
	} else {
		eventsMapping(r, content)
	}
}
