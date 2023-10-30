// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package azure

import (
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/logp"
)

func NewMetricRegistry() *MetricRegistry {
	return &MetricRegistry{
		collectionsInfo: make(map[string]MetricCollectionInfo),
	}
}

type MetricRegistry struct {
	collectionsInfo map[string]MetricCollectionInfo
}

func (m *MetricRegistry) Register(metric Metric, info MetricCollectionInfo) {
	m.collectionsInfo[m.buildMetricKey(metric)] = info
}

func (m *MetricRegistry) NeedsUpdate(metric Metric) bool {
	metricKey := m.buildMetricKey(metric)

	if info, exists := m.collectionsInfo[metricKey]; exists {
		duration := convertTimegrainToDuration(info.timeGrain)
		if info.timestamp.After(time.Now().Add(duration * (-1))) {
			return false
		}
	}

	return true
}

func (m *MetricRegistry) buildMetricKey(metric Metric) string {
	keyComponents := []string{
		metric.Namespace,
		metric.ResourceId,
	}
	keyComponents = append(keyComponents, metric.Names...)

	return strings.Join(keyComponents, ",")
}

type MetricCollectionInfo struct {
	timestamp time.Time
	timeGrain string
}

// Client represents the azure client which will make use of the azure sdk go metrics related clients
type Client struct {
	AzureMonitorService    Service
	Config                 Config
	ResourceConfigurations ResourceConfiguration
	Log                    *logp.Logger
	Resources              []Resource
	MetricRegistry         *MetricRegistry
}

// mapResourceMetrics function type will map the configuration options to client metrics (depending on the metricset)
type mapResourceMetrics func(client *Client, resources []*armresources.GenericResourceExpanded, resourceConfig ResourceConfig) ([]Metric, error)

// NewClient instantiates the Azure monitoring client
func NewClient(config Config) (*Client, error) {
	azureMonitorService, err := NewService(config)
	if err != nil {
		return nil, err
	}

	client := &Client{
		AzureMonitorService: azureMonitorService,
		Config:              config,
		Log:                 logp.NewLogger("azure monitor client"),
		MetricRegistry:      NewMetricRegistry(),
	}

	client.ResourceConfigurations.RefreshInterval = config.RefreshListInterval

	return client, nil
}

// InitResources function will retrieve and validate the resources configured by the users and then map the information configured to client metrics.
// the mapMetric function sent in this case will handle the mapping part as different metric and aggregation options work for different metricsets
func (client *Client) InitResources(fn mapResourceMetrics) error {
	if len(client.Config.Resources) == 0 {
		return fmt.Errorf("no resource options defined")
	}
	// check if refresh interval has been set and if it has expired
	if !client.ResourceConfigurations.Expired() {
		return nil
	}
	var metrics []Metric
	//reset client resources
	client.Resources = []Resource{}
	for _, resource := range client.Config.Resources {
		// retrieve azure resources information
		resourceList, err := client.AzureMonitorService.GetResourceDefinitions(resource.Id, resource.Group, resource.Type, resource.Query)
		if err != nil {
			err = fmt.Errorf("failed to retrieve resources: %w", err)
			return err
		}
		if len(resourceList) == 0 {
			err = fmt.Errorf("failed to retrieve resources: No resources returned using the configuration options resource ID %s, resource group %s, resource type %s, resource query %s",
				resource.Id, resource.Group, resource.Type, resource.Query)
			client.Log.Error(err)
			continue
		}
		// Map resources to the client
		for _, resource := range resourceList {
			if !containsResource(*resource.ID, client.Resources) {
				client.Resources = append(client.Resources, Resource{
					Id:           *resource.ID,
					Name:         *resource.Name,
					Location:     *resource.Location,
					Type:         *resource.Type,
					Group:        getResourceGroupFromId(*resource.ID),
					Tags:         mapTags(resource.Tags),
					Subscription: client.Config.SubscriptionId})
			}
		}
		resourceMetrics, err := fn(client, resourceList, resource)
		if err != nil {
			return err
		}
		metrics = append(metrics, resourceMetrics...)
	}
	// users could add or remove resources while metricbeat is running so we could encounter the situation where resources are unavailable we log an error message (see above)
	// we also log a debug message when absolutely no resources are found
	if len(metrics) == 0 {
		client.Log.Debug("no resources were found based on all the configurations options entered")
	}
	client.ResourceConfigurations.Metrics = metrics

	return nil
}

// GetMetricValues returns the specified metric data points for the specified resource ID/namespace.
func (client *Client) GetMetricValues(metrics []Metric, report mb.ReporterV2) []Metric {
	var resultedMetrics []Metric

	// Same end time for all metrics in the same batch
	referenceTime := time.Now().UTC()

	// loop over the set of metrics
	for _, metric := range metrics {
		// select period to collect metrics, will double the interval value in order to retrieve any missing values
		//if timegrain is larger than intervalx2 then interval will be assigned the timegrain value
		interval := client.Config.Period
		//duration := convertTimegrainToDuration(metric.TimeGrain)
		//if t := convertTimegrainToDuration(metric.TimeGrain); t > interval*2 {
		//	interval = t
		//}

		// copy the reference time
		endTime := referenceTime

		// Fetch in the range [{-2xINTERVAL},{-INTERVAL}) with a delay of {INTERVAL}
		// It results in one data point {-2xINTERVAL} per call
		endTime = endTime.Add(interval * (-1))
		startTime := endTime.Add(interval * (-1))
		timespan := fmt.Sprintf("%s/%s", startTime.Format(time.RFC3339), endTime.Format(time.RFC3339))

		if !client.MetricRegistry.NeedsUpdate(metric) {
			continue
		}

		//metricKey := strings.Join(metric.Names, ",")
		//metricKey += metric.ResourceId
		//
		//if metricInfo, exists := client.MetricRegistry[metricKey]; exists {
		//	duration := convertTimegrainToDuration(metricInfo.timeGrain)
		//	if metricInfo.timestamp.After(referenceTime.Add(duration * (-1))) {
		//		continue
		//	}
		//}

		// Interval math for timegrain > period
		//if duration > client.Config.Period {
		//	inTimespan := false
		//
		//	var diffSec = int64(endTime.Second() - startTime.Second())
		//
		//	var diffMin = int64(endTime.Minute() - startTime.Minute())
		//	var diffMinDuration = time.Duration(diffMin) * time.Minute
		//
		//	var diffHour = int64(endTime.Hour() - startTime.Hour())
		//	var diffHourDuration = time.Duration(diffHour) * time.Hour
		//
		//	// If timegrain is unit 1 day, 1 hour or 1 min
		//	if duration == 24*time.Hour {
		//		startOfDay := endTime.Truncate(24 * time.Hour)
		//		if (startOfDay.Equal(startTime) || startOfDay.After(startTime)) && startOfDay.Before(endTime) {
		//			inTimespan = true
		//		}
		//
		//	} else if duration >= time.Hour {
		//		if diffMin < 0 && diffHourDuration > 0 && diffHourDuration%duration == 0 {
		//			inTimespan = true
		//		}
		//	} else {
		//		if diffSec < 0 && diffMinDuration%duration == 0 {
		//			inTimespan = true
		//		}
		//	}
		//
		//	// if the time grain mark is not within the sampling timespan,
		//	// remove that metric from the list in this batch and skip to the next one
		//	if !inTimespan {
		//		// Remove metric from list
		//		ind := 0
		//		for i, currentMetric := range client.ResourceConfigurations.Metrics {
		//			if matchMetrics(currentMetric, metric) {
		//				ind = i
		//				break
		//			}
		//		}
		//		client.ResourceConfigurations.Metrics = append(client.ResourceConfigurations.Metrics[:ind], client.ResourceConfigurations.Metrics[ind+1:]...)
		//		continue
		//	}
		//}

		// build the 'filter' parameter which will contain any dimensions configured
		var filter string
		if len(metric.Dimensions) > 0 {
			var filterList []string
			for _, dim := range metric.Dimensions {
				filterList = append(filterList, dim.Name+" eq '"+dim.Value+"'")
			}
			filter = strings.Join(filterList, " AND ")
		}

		resp, timeGrain, err := client.AzureMonitorService.GetMetricValues(
			metric.ResourceSubId,
			metric.Namespace,
			metric.TimeGrain,
			timespan,
			metric.Names,
			metric.Aggregations,
			filter,
		)
		if err != nil {
			err = fmt.Errorf("error while listing metric values by resource ID %s and namespace  %s: %w", metric.ResourceSubId, metric.Namespace, err)
			client.Log.Error(err)
			report.Error(err)
		} else {

			// Update the metric registry with the latest timestamp and time grain.
			client.MetricRegistry.Register(metric, MetricCollectionInfo{
				timeGrain: timeGrain,
				timestamp: referenceTime,
			})

			//client.MetricRegistry[metricKey] = MetricCollectionInfo{
			//	timeGrain: timeGrain,
			//	timestamp: referenceTime,
			//}

			for i, currentMetric := range client.ResourceConfigurations.Metrics {
				if matchMetrics(currentMetric, metric) {
					//current := mapMetricValues(resp, currentMetric.Values, endTime.Truncate(time.Minute).Add(interval*(-1)), endTime.Truncate(time.Minute))
					current := mapMetricValues(resp, currentMetric.Values)
					client.ResourceConfigurations.Metrics[i].Values = current
					if client.ResourceConfigurations.Metrics[i].TimeGrain == "" {
						client.ResourceConfigurations.Metrics[i].TimeGrain = timeGrain
					}
					resultedMetrics = append(resultedMetrics, client.ResourceConfigurations.Metrics[i])
				}
			}
		}
	}

	return resultedMetrics
}

// CreateMetric function will create a client metric based on the resource and metrics configured
func (client *Client) CreateMetric(resourceId string, subResourceId string, namespace string, metrics []string, aggregations string, dimensions []Dimension, timeGrain string) Metric {
	if subResourceId == "" {
		subResourceId = resourceId
	}
	met := Metric{
		ResourceId:    resourceId,
		ResourceSubId: subResourceId,
		Namespace:     namespace,
		Names:         metrics,
		Dimensions:    dimensions,
		Aggregations:  aggregations,
		TimeGrain:     timeGrain,
	}

	for _, prevMet := range client.ResourceConfigurations.Metrics {
		if len(prevMet.Values) != 0 && matchMetrics(prevMet, met) {
			met.Values = prevMet.Values
		}
	}

	return met
}

// MapMetricByPrimaryAggregation will map the primary aggregation of the metric definition to the client metric
func (client *Client) MapMetricByPrimaryAggregation(metrics []armmonitor.MetricDefinition, resourceId string, subResourceId string, namespace string, dim []Dimension, timeGrain string) []Metric {
	var clientMetrics []Metric

	metricGroups := make(map[string][]armmonitor.MetricDefinition)

	for _, met := range metrics {
		metricGroups[string(*met.PrimaryAggregationType)] = append(metricGroups[string(*met.PrimaryAggregationType)], met)
	}

	for key, metricGroup := range metricGroups {
		var metricNames []string
		for _, metricName := range metricGroup {
			metricNames = append(metricNames, *metricName.Name.Value)
		}
		clientMetrics = append(clientMetrics, client.CreateMetric(resourceId, subResourceId, namespace, metricNames, key, dim, timeGrain))
	}

	return clientMetrics
}

// GetVMForMetaData func will retrieve the vm details in order to fill in the cloud metadata and also update the client resources
func (client *Client) GetVMForMetaData(resource *Resource, metricValues []MetricValue) VmResource {
	var (
		vm           VmResource
		resourceName = resource.Name
		resourceId   = resource.Id
	)

	// check first if this is a vm scaleset and the instance name is stored in the dimension value
	if dimension, ok := getDimension("VMName", metricValues[0].dimensions); ok {
		instanceId := getInstanceId(dimension.Value)
		if instanceId != "" {
			resourceId += fmt.Sprintf("/virtualMachines/%s", instanceId)
			resourceName = dimension.Value
		}
	}

	// if vm has been already added to the resource then it should be returned
	if existingVM, ok := getVM(resourceName, resource.Vms); ok {
		return existingVM
	}

	// an additional call is necessary in order to retrieve the vm specific details
	expandedResource, err := client.AzureMonitorService.GetResourceDefinitionById(resourceId)
	if err != nil {
		client.Log.Error(err, "could not retrieve the resource details by resource ID %s", resourceId)
		return VmResource{}
	}

	vm.Name = *expandedResource.Name

	if expandedResource.Properties != nil {
		if properties, ok := expandedResource.Properties.(map[string]interface{}); ok {
			if hardware, ok := properties["hardwareProfile"]; ok {
				if vmSz, ok := hardware.(map[string]interface{})["vmSize"]; ok {
					vm.Size = vmSz.(string)
				}
				if vmID, ok := properties["vmId"]; ok {
					vm.Id = vmID.(string)
				}
			}
		}
	}

	if len(vm.Size) == 0 && expandedResource.SKU != nil && expandedResource.SKU.Name != nil {
		vm.Size = *expandedResource.SKU.Name
	}

	// the client resource and selected resources are being updated in order to avoid additional calls
	client.AddVmToResource(resource.Id, vm)

	resource.Vms = append(resource.Vms, vm)

	return vm
}

// GetResourceForMetaData will retrieve resource details for the selected metric configuration
func (client *Client) GetResourceForMetaData(grouped Metric) Resource {
	for _, res := range client.Resources {
		if res.Id == grouped.ResourceId {
			return res
		}
	}
	return Resource{}
}

func (client *Client) LookupResource(resourceId string) Resource {
	for _, res := range client.Resources {
		if res.Id == resourceId {
			return res
		}
	}
	return Resource{}
}

// AddVmToResource will add the vm details to the resource
func (client *Client) AddVmToResource(resourceId string, vm VmResource) {
	if len(vm.Id) > 0 && len(vm.Name) > 0 {
		for i, res := range client.Resources {
			if res.Id == resourceId {
				client.Resources[i].Vms = append(client.Resources[i].Vms, vm)
			}
		}
	}
}

// NewMockClient instantiates a new client with the mock azure service
func NewMockClient() *Client {
	azureMockService := new(MockService)
	client := &Client{
		AzureMonitorService: azureMockService,
		Config:              Config{},
		Log:                 logp.NewLogger("test azure monitor"),
	}
	return client
}
