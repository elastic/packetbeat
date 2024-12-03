package wmi

import (
	"fmt"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"

	wmiquery "github.com/microsoft/wmi/pkg/base/query"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("windows", "wmi", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	config WmibeatConfig
}

const WMIDefaultNamespace = "root\\cimv2"

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The windows wmi metricset is beta.")

	config := NewDefaultConfig()
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	m := &MetricSet{
		BaseMetricSet: base,
		config:        config,
	}

	// Compiling the Query once
	for i := range m.config.Queries {
		q := m.config.Queries[i]
		m.config.Queries[i].Query = wmiquery.NewWmiQueryWithSelectList(q.Class, q.Fields, q.Where...)
	}

	return m, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {

	var err error

	sm := wmi.NewWmiSessionManager()
	defer sm.Dispose()

	session, err := sm.GetSession(m.config.Namespace, m.config.Host, "", m.config.User, m.config.Password)

	if err != nil {
		return fmt.Errorf("could not initialize session %v", err)
	}
	_, err = session.Connect()
	if err != nil {
		return fmt.Errorf("could not connect session %v", err)
	}
	defer session.Dispose()

	for _, queryConfig := range m.config.Queries {

		rows, err := session.QueryInstances(queryConfig.Query.String())
		if err != nil {
			logp.Warn("Could not execute query %v", err)
			continue
		}

		for _, instance := range rows {
			event := mb.Event{
				MetricSetFields: mapstr.M{
					"class":     queryConfig.Class,
					"namespace": m.config.Namespace,
				},
			}

			if m.config.IncludeQueries {
				event.MetricSetFields.Put(".query", queryConfig.Query)
			}

			// Get only the required properites
			properties := queryConfig.Fields

			// With special array, we retrive all properties
			if len(queryConfig.Fields) == 1 && queryConfig.Fields[0] == "*" {
				properties = instance.GetClass().GetPropertiesNames()
			}

			for _, fieldName := range properties {
				fieldValue, err := instance.GetProperty(fieldName)
				if err != nil {
					logp.Err("Unable to get propery by name: %v", err)
					continue
				}

				if !m.config.IncludeNull && fieldValue == nil {
					continue
				}

				event.MetricSetFields.Put(fieldName, fieldValue)
			}
			report.Event(event)
		}
	}
	return nil
}
