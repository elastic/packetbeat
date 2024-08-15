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

package cluster

import (
	"context"
	"fmt"
	"strings"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/vsphere"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func init() {
	mb.Registry.MustAddMetricSet("vsphere", "cluster", New,
		mb.WithHostParser(vsphere.HostParser),
		mb.DefaultMetricSet(),
	)
}

// MetricSet type defines all fields of the MetricSet.
type MetricSet struct {
	*vsphere.MetricSet
}

type assetNames struct {
	outputNtNames []string
	outputDsNames []string
	outputHsNames []string
}

// New creates a new instance of the MetricSet.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	ms, err := vsphere.NewMetricSet(base)
	if err != nil {
		return nil, err
	}
	return &MetricSet{ms}, nil
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

	// Create a view of Cluster objects
	mgr := view.NewManager(c)

	v, err := mgr.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"ClusterComputeResource"}, true)
	if err != nil {
		return fmt.Errorf("error in CreateContainerView: %w", err)
	}

	defer func() {
		if err := v.Destroy(ctx); err != nil {
			m.Logger().Debug(fmt.Errorf("error trying to destroy view from vshphere: %w", err))
		}
	}()

	// Retrieve summary property for all Clusters
	var clt []mo.ClusterComputeResource
	if err = v.Retrieve(ctx, []string{"ClusterComputeResource"}, []string{}, &clt); err != nil {
		return fmt.Errorf("error in Retrieve: %w", err)
	}

	pc := property.DefaultCollector(c)

	for _, cl := range clt {

		assetNames, err := getAssetNames(ctx, pc, &cl)
		if err != nil {
			m.Logger().Errorf("Failed to retrieve object from host: %w", err)
		}

		reporter.Event(mb.Event{
			MetricSetFields: m.eventMapping(cl, assetNames),
		})
	}

	return nil

}

func getAssetNames(ctx context.Context, pc *property.Collector, cl *mo.ClusterComputeResource) (*assetNames, error) {
	referenceList := make([]types.ManagedObjectReference, 0, len(cl.Datastore)+len(cl.Host))
	referenceList = append(referenceList, cl.Datastore...)
	referenceList = append(referenceList, cl.Host...)

	var objects []mo.ManagedEntity
	if len(referenceList) > 0 {
		if err := pc.Retrieve(ctx, referenceList, []string{"name"}, &objects); err != nil {
			return nil, err
		}
	}

	outputDsNames := make([]string, 0, len(cl.Datastore))
	outputHsNames := make([]string, 0, len(cl.Host))
	for _, ob := range objects {
		name := strings.ReplaceAll(ob.Name, ".", "_")
		switch ob.Reference().Type {
		case "Datastore":
			outputDsNames = append(outputDsNames, name)
		case "HostSystem":
			outputHsNames = append(outputHsNames, name)
		}
	}

	// calling network explicitly because of mo.Network's ManagedEntityObject.Name does not store Network name
	// instead mo.Network.Name contains correct value of Network name
	var netObjects []mo.Network
	if len(cl.Network) > 0 {
		if err := pc.Retrieve(ctx, cl.Network, []string{"name"}, &netObjects); err != nil {
			return nil, err
		}
	}

	outputNtNames := make([]string, 0, len(cl.Network))
	for _, ob := range netObjects {
		name := strings.ReplaceAll(ob.Name, ".", "_")
		outputNtNames = append(outputNtNames, name)
	}

	return &assetNames{
		outputNtNames: outputNtNames,
		outputDsNames: outputDsNames,
		outputHsNames: outputHsNames,
	}, nil
}
