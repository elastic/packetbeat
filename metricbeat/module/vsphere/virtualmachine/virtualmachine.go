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

package virtualmachine

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/vsphere"
	"github.com/elastic/elastic-agent-libs/mapstr"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func init() {
	mb.Registry.MustAddMetricSet("vsphere", "virtualmachine", New,
		mb.WithHostParser(vsphere.HostParser),
		mb.DefaultMetricSet(),
	)
}

// MetricSet type defines all fields of the MetricSet.
type MetricSet struct {
	*vsphere.MetricSet
	GetCustomFields bool
}

<<<<<<< HEAD
=======
type VMData struct {
	VM             mo.VirtualMachine
	HostID         string
	HostName       string
	NetworkNames   []string
	DatastoreNames []string
	CustomFields   mapstr.M
	Snapshots      []VMSnapshotData
}

type VMSnapshotData struct {
	ID          int32                          `json:"id"`
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	CreateTime  common.Time                    `json:"createtime"`
	State       types.VirtualMachinePowerState `json:"state"`
}

>>>>>>> 3f44bd1f9b ([Metricbeat][vSphere] New metrics support and minor changes to existing metricsets (#40766))
// New creates a new instance of the MetricSet.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	ms, err := vsphere.NewMetricSet(base)
	if err != nil {
		return nil, err
	}

	config := struct {
		GetCustomFields bool `config:"get_custom_fields"`
	}{
		GetCustomFields: false,
	}

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}
	return &MetricSet{
		MetricSet:       ms,
		GetCustomFields: config.GetCustomFields,
	}, nil
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

	// Get custom fields (attributes) names if get_custom_fields is true.
	customFieldsMap := make(map[int32]string)
	if m.GetCustomFields {
		var err error
		customFieldsMap, err = setCustomFieldsMap(ctx, c)
		if err != nil {
			return fmt.Errorf("error in setCustomFieldsMap: %w", err)
		}
	}

	// Create view of VirtualMachine objects
	mgr := view.NewManager(c)

	v, err := mgr.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		return fmt.Errorf("error in CreateContainerView: %w", err)
	}

	defer func() {
		if err := v.Destroy(ctx); err != nil {
			m.Logger().Debug(fmt.Errorf("error trying to destroy view from vshphere: %w", err))
		}
	}()

	// Retrieve summary property for all machines
	var vmt []mo.VirtualMachine
<<<<<<< HEAD
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vmt)
=======
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary", "datastore", "snapshot"}, &vmt)
>>>>>>> 3f44bd1f9b ([Metricbeat][vSphere] New metrics support and minor changes to existing metricsets (#40766))
	if err != nil {
		return fmt.Errorf("error in Retrieve: %w", err)
	}

	for _, vm := range vmt {
		usedMemory := int64(vm.Summary.QuickStats.GuestMemoryUsage) * 1024 * 1024
		usedCPU := vm.Summary.QuickStats.OverallCpuUsage
		event := mapstr.M{
			"name": vm.Summary.Config.Name,
			"os":   vm.Summary.Config.GuestFullName,
			"cpu": mapstr.M{
				"used": mapstr.M{
					"mhz": usedCPU,
				},
			},
			"memory": mapstr.M{
				"used": mapstr.M{
					"guest": mapstr.M{
						"bytes": usedMemory,
					},
					"host": mapstr.M{
						"bytes": int64(vm.Summary.QuickStats.HostMemoryUsage) * 1024 * 1024,
					},
				},
			},
		}

		totalCPU := vm.Summary.Config.CpuReservation
		if totalCPU > 0 {
			freeCPU := totalCPU - usedCPU
			// Avoid negative values if reported used CPU is slightly over total configured.
			if freeCPU < 0 {
				freeCPU = 0
			}
			event.Put("cpu.total.mhz", totalCPU)
			event.Put("cpu.free.mhz", freeCPU)
		}

		totalMemory := int64(vm.Summary.Config.MemorySizeMB) * 1024 * 1024
		if totalMemory > 0 {
			freeMemory := totalMemory - usedMemory
			// Avoid negative values if reported used memory is slightly over total configured.
			if freeMemory < 0 {
				freeMemory = 0
			}
			event.Put("memory.total.guest.bytes", totalMemory)
			event.Put("memory.free.guest.bytes", freeMemory)
		}

		if host := vm.Summary.Runtime.Host; host != nil {
			event["host.id"] = host.Value
			hostSystem, err := getHostSystem(ctx, c, host.Reference())
			if err == nil {
				event["host.hostname"] = hostSystem.Summary.Config.Name
			} else {
				m.Logger().Debug(err.Error())
			}
		} else {
			m.Logger().Debug("'Host', 'Runtime' or 'Summary' data not found. This is either a parsing error " +
				"from vsphere library, an error trying to reach host/guest or incomplete information returned " +
				"from host/guest")
		}

		// Get custom fields (attributes) values if get_custom_fields is true.
		if m.GetCustomFields && vm.Summary.CustomValue != nil {
			customFields := getCustomFields(vm.Summary.CustomValue, customFieldsMap)

			if len(customFields) > 0 {
				event["custom_fields"] = customFields
			}
		} else {
			m.Logger().Debug("custom fields not activated or custom values not found/parse in Summary data. This " +
				"is either a parsing error from vsphere library, an error trying to reach host/guest or incomplete " +
				"information returned from host/guest")
		}

		if vm.Summary.Vm != nil {
			networkNames, err := getNetworkNames(ctx, c, vm.Summary.Vm.Reference())
			if err != nil {
				m.Logger().Debug(err.Error())
			} else {
				if len(networkNames) > 0 {
					event["network_names"] = networkNames
				}
			}
		}

		reporter.Event(mb.Event{
			MetricSetFields: event,
		})
	}

	return nil
}

func getCustomFields(customFields []types.BaseCustomFieldValue, customFieldsMap map[int32]string) mapstr.M {
	outputFields := mapstr.M{}
	for _, v := range customFields {
		customFieldString := v.(*types.CustomFieldStringValue)
		key, ok := customFieldsMap[v.GetCustomFieldValue().Key]
		if ok {
			// If key has '.', is replaced with '_' to be compatible with ES2.x.
			fmtKey := strings.Replace(key, ".", "_", -1)
			outputFields.Put(fmtKey, customFieldString.Value)
		}
	}

	return outputFields
}

func getNetworkNames(ctx context.Context, c *vim25.Client, ref types.ManagedObjectReference) ([]string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var outputNetworkNames []string

	pc := property.DefaultCollector(c)

	var vm mo.VirtualMachine
	err := pc.RetrieveOne(ctx, ref, []string{"network"}, &vm)
	if err != nil {
		return nil, fmt.Errorf("error retrieving virtual machine information: %v", err)
	}

	if len(vm.Network) == 0 {
		return nil, errors.New("no networks found")
	}

	var networkRefs []types.ManagedObjectReference
	for _, obj := range vm.Network {
		if obj.Type == "Network" {
			networkRefs = append(networkRefs, obj)
		}
	}

	// If only "Distributed port group" was found, for example.
	if len(networkRefs) == 0 {
		return nil, errors.New("no networks found")
	}

	var nets []mo.Network
	err = pc.Retrieve(ctx, networkRefs, []string{"name"}, &nets)
	if err != nil {
		return nil, fmt.Errorf("error retrieving network from virtual machine: %v", err)
	}

	for _, net := range nets {
		name := strings.Replace(net.Name, ".", "_", -1)
		outputNetworkNames = append(outputNetworkNames, name)
	}

	return outputNetworkNames, nil
}

func setCustomFieldsMap(ctx context.Context, client *vim25.Client) (map[int32]string, error) {
	customFieldsMap := make(map[int32]string)

	customFieldsManager, err := object.GetCustomFieldsManager(client)

	if err != nil {
		return nil, fmt.Errorf("failed to get custom fields manager: %w", err)
	}
	field, err := customFieldsManager.Field(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get custom fields: %w", err)
	}

	for _, def := range field {
		customFieldsMap[def.Key] = def.Name
	}

	return customFieldsMap, nil
}

func getHostSystem(ctx context.Context, c *vim25.Client, ref types.ManagedObjectReference) (*mo.HostSystem, error) {
	pc := property.DefaultCollector(c)

	var hs mo.HostSystem
	err := pc.RetrieveOne(ctx, ref, []string{"summary"}, &hs)
	if err != nil {
		return nil, fmt.Errorf("error retrieving host information: %v", err)
	}
	return &hs, nil
}
<<<<<<< HEAD
=======

func fetchSnapshots(snapshotTree []types.VirtualMachineSnapshotTree) []VMSnapshotData {
	snapshots := make([]VMSnapshotData, 0, len(snapshotTree))
	for _, snapshot := range snapshotTree {
		snapshots = append(snapshots, VMSnapshotData{
			ID:          snapshot.Id,
			Name:        snapshot.Name,
			Description: snapshot.Description,
			CreateTime:  common.Time(snapshot.CreateTime),
			State:       snapshot.State,
		})

		// Recursively add child snapshots
		if len(snapshot.ChildSnapshotList) > 0 {
			snapshots = append(snapshots, fetchSnapshots(snapshot.ChildSnapshotList)...)
		}
	}
	return snapshots
}
>>>>>>> 3f44bd1f9b ([Metricbeat][vSphere] New metrics support and minor changes to existing metricsets (#40766))
