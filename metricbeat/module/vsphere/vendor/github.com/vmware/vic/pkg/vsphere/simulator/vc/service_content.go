// Copyright 2016 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vc

import "github.com/vmware/govmomi/vim25/types"

var ServiceContent = types.ServiceContent{
	DynamicData:       types.DynamicData{},
	RootFolder:        types.ManagedObjectReference{Type: "Folder", Value: "group-d1"},
	PropertyCollector: types.ManagedObjectReference{Type: "PropertyCollector", Value: "propertyCollector"},
	ViewManager:       &types.ManagedObjectReference{Type: "ViewManager", Value: "ViewManager"},
	About: types.AboutInfo{
		DynamicData:           types.DynamicData{},
		Name:                  "VMware vCenter Server",
		FullName:              "VMware vCenter Server 6.0.0 build-3634794",
		Vendor:                "VMware, Inc.",
		Version:               "6.0.0",
		Build:                 "3634794",
		LocaleVersion:         "INTL",
		LocaleBuild:           "000",
		OsType:                "linux-x64",
		ProductLineId:         "vpx",
		ApiType:               "VirtualCenter",
		ApiVersion:            "6.0",
		InstanceUuid:          "8e7597d6-f720-4b5c-b9fc-3412faf07d99",
		LicenseProductName:    "VMware VirtualCenter Server",
		LicenseProductVersion: "6.0",
	},
	Setting:                   &types.ManagedObjectReference{Type: "OptionManager", Value: "VpxSettings"},
	UserDirectory:             &types.ManagedObjectReference{Type: "UserDirectory", Value: "UserDirectory"},
	SessionManager:            &types.ManagedObjectReference{Type: "SessionManager", Value: "SessionManager"},
	AuthorizationManager:      &types.ManagedObjectReference{Type: "AuthorizationManager", Value: "AuthorizationManager"},
	ServiceManager:            &types.ManagedObjectReference{Type: "ServiceManager", Value: "ServiceMgr"},
	PerfManager:               &types.ManagedObjectReference{Type: "PerformanceManager", Value: "PerfMgr"},
	ScheduledTaskManager:      &types.ManagedObjectReference{Type: "ScheduledTaskManager", Value: "ScheduledTaskManager"},
	AlarmManager:              &types.ManagedObjectReference{Type: "AlarmManager", Value: "AlarmManager"},
	EventManager:              &types.ManagedObjectReference{Type: "EventManager", Value: "EventManager"},
	TaskManager:               &types.ManagedObjectReference{Type: "TaskManager", Value: "TaskManager"},
	ExtensionManager:          &types.ManagedObjectReference{Type: "ExtensionManager", Value: "ExtensionManager"},
	CustomizationSpecManager:  &types.ManagedObjectReference{Type: "CustomizationSpecManager", Value: "CustomizationSpecManager"},
	CustomFieldsManager:       &types.ManagedObjectReference{Type: "CustomFieldsManager", Value: "CustomFieldsManager"},
	AccountManager:            (*types.ManagedObjectReference)(nil),
	DiagnosticManager:         &types.ManagedObjectReference{Type: "DiagnosticManager", Value: "DiagMgr"},
	LicenseManager:            &types.ManagedObjectReference{Type: "LicenseManager", Value: "LicenseManager"},
	SearchIndex:               &types.ManagedObjectReference{Type: "SearchIndex", Value: "SearchIndex"},
	FileManager:               &types.ManagedObjectReference{Type: "FileManager", Value: "FileManager"},
	DatastoreNamespaceManager: &types.ManagedObjectReference{Type: "DatastoreNamespaceManager", Value: "DatastoreNamespaceManager"},
	VirtualDiskManager:        &types.ManagedObjectReference{Type: "VirtualDiskManager", Value: "virtualDiskManager"},
	VirtualizationManager:     (*types.ManagedObjectReference)(nil),
	SnmpSystem:                &types.ManagedObjectReference{Type: "HostSnmpSystem", Value: "SnmpSystem"},
	VmProvisioningChecker:     &types.ManagedObjectReference{Type: "VirtualMachineProvisioningChecker", Value: "ProvChecker"},
	VmCompatibilityChecker:    &types.ManagedObjectReference{Type: "VirtualMachineCompatibilityChecker", Value: "CompatChecker"},
	OvfManager:                &types.ManagedObjectReference{Type: "OvfManager", Value: "OvfManager"},
	IpPoolManager:             &types.ManagedObjectReference{Type: "IpPoolManager", Value: "IpPoolManager"},
	DvSwitchManager:           &types.ManagedObjectReference{Type: "DistributedVirtualSwitchManager", Value: "DVSManager"},
	HostProfileManager:        &types.ManagedObjectReference{Type: "HostProfileManager", Value: "HostProfileManager"},
	ClusterProfileManager:     &types.ManagedObjectReference{Type: "ClusterProfileManager", Value: "ClusterProfileManager"},
	ComplianceManager:         &types.ManagedObjectReference{Type: "ProfileComplianceManager", Value: "MoComplianceManager"},
	LocalizationManager:       &types.ManagedObjectReference{Type: "LocalizationManager", Value: "LocalizationManager"},
	StorageResourceManager:    &types.ManagedObjectReference{Type: "StorageResourceManager", Value: "StorageResourceManager"},
	GuestOperationsManager:    &types.ManagedObjectReference{Type: "GuestOperationsManager", Value: "guestOperationsManager"},
	OverheadMemoryManager:     &types.ManagedObjectReference{Type: "OverheadMemoryManager", Value: "OverheadMemoryManger"},
	CertificateManager:        &types.ManagedObjectReference{Type: "CertificateManager", Value: "certificateManager"},
	IoFilterManager:           &types.ManagedObjectReference{Type: "IoFilterManager", Value: "IoFilterManager"},
}
