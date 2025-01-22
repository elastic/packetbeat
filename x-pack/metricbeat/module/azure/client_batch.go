// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package azure

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/logp"
)

// BatchClient represents the azure batch client which will make use of the azure sdk go metrics related clients
type BatchClient struct {
	AzureMonitorService    Service
	Config                 Config
	ResourceConfigurations ConcurrentResourceConfig
	Log                    *logp.Logger
	Resources              []Resource
	MetricRegistry         *MetricRegistry
}

// Resource definitions grouping criteria
type ResDefGroupingCriteria struct {
	Namespace      string
	SubscriptionID string
	Location       string
	Names          string
	Aggregations   string
	TimeGrain      string
	Dimensions     string
}

// concurrentMapResourceMetrics function type will map the configuration options to Batch Client metrics (depending on the metricset)
type concurrentMapResourceMetrics func(client *BatchClient, resources []*armresources.GenericResourceExpanded, resourceConfig ResourceConfig, wg *sync.WaitGroup)

// NewBatchClient instantiates the Azure monitoring batch client
func NewBatchClient(config Config) (*BatchClient, error) {
	azureMonitorService, err := NewService(config)
	if err != nil {
		return nil, err
	}

	logger := logp.NewLogger("azure monitor client")

	client := &BatchClient{
		AzureMonitorService: azureMonitorService,
		Config:              config,
		Log:                 logger,
		MetricRegistry:      NewMetricRegistry(logger),
	}
	client.ResourceConfigurations.MetricDefinitions = MetricDefinitions{
		Update:  true,
		Metrics: make(map[string][]Metric),
	}
	client.ResourceConfigurations.RefreshInterval = config.RefreshListInterval
	client.ResourceConfigurations.MetricDefinitionsChan = nil
	client.ResourceConfigurations.ErrorChan = nil

	return client, nil
}

// InitResources function will retrieve and validate the resources configured by the users and then map the information configured to client metrics.
// the mapMetric function sent in this case will handle the mapping part as different metric and aggregation options work for different metricsets
func (client *BatchClient) InitResources(fn concurrentMapResourceMetrics) error {
	if len(client.Config.Resources) == 0 {
		return fmt.Errorf("no resource options defined")
	}

	// check if refresh interval has been set and if it has expired
	if !client.ResourceConfigurations.Expired() {
		client.Log.Infof("MetricDefinitions are not expired. Writing metrics to MetricDefinitionsChan")
		client.ResourceConfigurations.MetricDefinitionsChan = make(chan []Metric)
		client.ResourceConfigurations.ErrorChan = make(chan error, 1)
		// MetricDefinitions do not need update
		client.ResourceConfigurations.MetricDefinitions.Update = false
		go func() {
			defer close(client.ResourceConfigurations.MetricDefinitionsChan)
			defer close(client.ResourceConfigurations.ErrorChan)
			for _, metrics := range client.ResourceConfigurations.MetricDefinitions.Metrics {
				client.ResourceConfigurations.MetricDefinitionsChan <- metrics
			}
		}()
		return nil
	}
	// MetricDefinitions need update
	client.ResourceConfigurations.MetricDefinitions.Update = true

	// Initialize a WaitGroup to track all goroutines
	var wg sync.WaitGroup
	//reset client resources
	client.Resources = []Resource{}
	for _, resourceConfig := range client.Config.Resources {
		// retrieve azure resources information
		resourceList, err := client.AzureMonitorService.GetResourceDefinitions(resourceConfig.Id, resourceConfig.Group, resourceConfig.Type, resourceConfig.Query)
		if err != nil {
			err = fmt.Errorf("failed to retrieve resources: %w", err)
			// Should we return here or continue?
			return err
		}

		if len(resourceList) == 0 {
			err = fmt.Errorf("failed to retrieve resources: No resources returned using the configuration options resource ID %s, resource group %s, resource type %s, resource query %s",
				resourceConfig.Id, resourceConfig.Group, resourceConfig.Type, resourceConfig.Query)
			client.Log.Error(err)
			continue
		}
		// create the channels if they are not already created by a previous itteration
		if client.ResourceConfigurations.MetricDefinitionsChan == nil && client.ResourceConfigurations.ErrorChan == nil {
			client.ResourceConfigurations.MetricDefinitionsChan = make(chan []Metric)
			client.ResourceConfigurations.ErrorChan = make(chan error, 1)
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

		// Collects and stores metrics definitions for the cloud resources.
		wg.Add(1)
		fn(client, resourceList, resourceConfig, &wg)
		client.Log.Infof("Finished collection with %d metric definitions", len(resourceList))
	}
	go func() {
		wg.Wait() // Wait for all the resource collection goroutines to finish
		// Once all the goroutines are done, close the channels
		client.Log.Infof("All collections finished. Closing channels ")
		close(client.ResourceConfigurations.MetricDefinitionsChan)
		close(client.ResourceConfigurations.ErrorChan)
	}()
	return nil
}

// GetMetricsInBatch will query the batch API for each group
func (client *BatchClient) GetMetricsInBatch(groupedMetrics map[ResDefGroupingCriteria][]Metric, referenceTime time.Time, reporter mb.ReporterV2) []Metric {
	var result []Metric
	for criteria, metricsDefinitions := range groupedMetrics {
		// Same end time for all metrics in the same batch.
		interval := client.Config.Period

		// // Fetch in the range [{-2 x INTERVAL},{-1 x INTERVAL}) with a delay of {INTERVAL}.
		endTime := referenceTime
		// startTime := endTime.Add(interval * (-1))
		timespanDuration := max(asDuration(criteria.TimeGrain), interval)
		startTime := endTime.Add(timespanDuration * -1)
		// Limit batch size to 50 resources (if you have more, you can split the batch)
		filter := ""
		if len(metricsDefinitions[0].Dimensions) > 0 {
			var filterList []string
			for _, dim := range metricsDefinitions[0].Dimensions {
				filterList = append(filterList, dim.Name+" eq '"+dim.Value+"'")
			}
			filter = strings.Join(filterList, " AND ")
		}
		for i := 0; i < len(metricsDefinitions); i += BatchApiResourcesLimit {
			end := i + BatchApiResourcesLimit
			if end > len(metricsDefinitions) {
				end = len(metricsDefinitions)
			}

			// Slice the metrics to form the batch request
			batchMetrics := metricsDefinitions[i:end]
			// Make the batch API call (adjust parameters as needed)
			response, err := client.AzureMonitorService.QueryResources(
				getResourceIDs(batchMetrics), // Get the resource IDs from the batch
				criteria.SubscriptionID,
				criteria.Namespace,
				criteria.TimeGrain,
				startTime.Format("2006-01-02T15:04:05.000Z07:00"),
				endTime.Format("2006-01-02T15:04:05.000Z07:00"),
				strings.Split(criteria.Names, ","),
				strings.ToLower(batchMetrics[0].Aggregations),
				filter,
				criteria.Location,
			)
			if err != nil {
				err = fmt.Errorf("error while listing metric values by resource ID %s and namespace  %s: %w", metricsDefinitions[0].ResourceSubId, metricsDefinitions[0].Namespace, err)
				client.Log.Error(err)
				reporter.Error(err)
				continue
			}

			// Process the response as needed
			for i, v := range response {
				client.MetricRegistry.Update(metricsDefinitions[i], MetricCollectionInfo{
					timeGrain: *response[i].Interval,
					timestamp: referenceTime,
				})
				values := mapBatchMetricValues(client, v)
				metricsDefinitions[i].Values = append(metricsDefinitions[i].Values, values...)
				if metricsDefinitions[i].TimeGrain == "" {
					metricsDefinitions[i].TimeGrain = *response[i].Interval
				}

			}

			result = append(result, metricsDefinitions...)
		}
	}

	return result
}

func (client *BatchClient) GroupAndStoreMetrics(metricsDefinitions []Metric, referenceTime time.Time, store map[ResDefGroupingCriteria]*MetricStore) {
	for _, metric := range metricsDefinitions {

		criteria := ResDefGroupingCriteria{
			Namespace:      metric.Namespace,
			SubscriptionID: metric.SubscriptionId,
			Location:       metric.Location,
			Names:          strings.Join(metric.Names, ","),
			TimeGrain:      metric.TimeGrain,
			Dimensions:     getDimensionKey(metric.Dimensions),
		}

		//
		// Before fetching the metric values, check if the metric
		// has been collected within the time grain.
		//
		// Why do we need this?
		//
		// Some metricsets contains metrics with long time grains (e.g. 1 hour).
		//
		// If we collect the metric values every 5 minutes, we will end up fetching
		// the same data over and over again for all metrics with a time grain
		// larger than 5 minutes.
		//
		// The registry keeps track of the last timestamp the metricset collected
		// the metric values and the time grain used.
		//
		// By comparing the last collection time with the current time, and
		// the time grain of the metric, we can determine if the metric needs
		// to be collected again, or if we can skip it.
		//
		if !client.MetricRegistry.NeedsUpdate(referenceTime, metric) {
			continue
		}
		if _, exists := store[criteria]; !exists {
			store[criteria] = &MetricStore{}
		}
		store[criteria].AddMetric(metric)
	}
}

// CreateMetric function will create a client metric based on the resource and metrics configured
func (client *BatchClient) CreateMetric(resourceId string, subResourceId string, namespace string, location string, subscriptionId string, metrics []string, aggregations string, dimensions []Dimension, timeGrain string) Metric {
	if subResourceId == "" {
		subResourceId = resourceId
	}
	met := Metric{
		ResourceId:     resourceId,
		ResourceSubId:  subResourceId,
		Namespace:      namespace,
		Names:          metrics,
		Dimensions:     dimensions,
		Aggregations:   aggregations,
		TimeGrain:      timeGrain,
		Location:       location,
		SubscriptionId: subscriptionId,
	}
	if prevMetrics, ok := client.ResourceConfigurations.MetricDefinitions.Metrics[resourceId]; ok {
		for _, prevMet := range prevMetrics {
			if len(prevMet.Values) != 0 && matchMetrics(prevMet, met) {
				met.Values = prevMet.Values
			}
		}
	}

	return met
}

// MapMetricByPrimaryAggregation will map the primary aggregation of the metric definition to the client metric
func (client *BatchClient) MapMetricByPrimaryAggregation(metrics []armmonitor.MetricDefinition, resourceId string, location string, subscriptionId string, subResourceId string, namespace string, dim []Dimension, timeGrain string) []Metric {
	clientMetrics := make([]Metric, 0)
	metricGroups := make(map[string][]armmonitor.MetricDefinition)

	for _, met := range metrics {
		metricGroups[string(*met.PrimaryAggregationType)] = append(metricGroups[string(*met.PrimaryAggregationType)], met)
	}

	for key, metricGroup := range metricGroups {
		var metricNames []string
		for _, metricName := range metricGroup {
			metricNames = append(metricNames, *metricName.Name.Value)
		}
		clientMetrics = append(clientMetrics, client.CreateMetric(resourceId, subResourceId, namespace, location, subscriptionId, metricNames, key, dim, timeGrain))
	}

	return clientMetrics
}

// GetVMForMetadata func will retrieve the VM details in order to fill in the cloud metadata
// and also update the client resources
func (client *BatchClient) GetVMForMetadata(resource *Resource, referencePoint KeyValuePoint) VmResource {
	var (
		vm           VmResource
		resourceName = resource.Name
		resourceId   = resource.Id
	)

	// Search the dimensions for the "VMName" dimension. This dimension is present for VM Scale Sets.
	if dimensionValue, ok := getDimension("VMName", referencePoint.Dimensions); ok {
		instanceId := getInstanceId(dimensionValue)
		if instanceId != "" {
			resourceId += fmt.Sprintf("/virtualMachines/%s", instanceId)
			resourceName = dimensionValue
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
func (client *BatchClient) GetResourceForMetaData(grouped Metric) Resource {
	for _, res := range client.Resources {
		if res.Id == grouped.ResourceId {
			return res
		}
	}
	return Resource{}
}

func (client *BatchClient) LookupResource(resourceId string) Resource {
	for _, res := range client.Resources {
		if res.Id == resourceId {
			return res
		}
	}
	return Resource{}
}

// AddVmToResource will add the vm details to the resource
func (client *BatchClient) AddVmToResource(resourceId string, vm VmResource) {
	if len(vm.Id) > 0 && len(vm.Name) > 0 {
		for i, res := range client.Resources {
			if res.Id == resourceId {
				client.Resources[i].Vms = append(client.Resources[i].Vms, vm)
			}
		}
	}
}

// mapToEvents maps the metric values to events and reports them to Elasticsearch.
func (client *BatchClient) MapToEvents(metrics []Metric, reporter mb.ReporterV2) error {

	// Map the metric values into a list of key/value points.
	//
	// This makes it easier to group the metrics by timestamp
	// and dimensions.
	points := mapToKeyValuePoints(metrics)

	// Group the points by timestamp and other fields we consider
	// as dimensions for the whole event.
	//
	// Metrics have their own dimensions, and this is fine at the
	// metric level.
	//
	// We identified a set of field we consider as dimensions
	// at the event level. The event level dimensions define
	// the time series when TSDB is enabled.
	groupedPoints := make(map[string][]KeyValuePoint)
	for _, point := range points {
		groupingKey := fmt.Sprintf(
			"%s,%s,%s,%s,%s,%s",
			point.Timestamp,
			point.Namespace,
			point.ResourceId,
			point.ResourceSubId,
			point.Dimensions,
			point.TimeGrain,
		)

		groupedPoints[groupingKey] = append(groupedPoints[groupingKey], point)
	}

	// Create an event for each group of points and send
	// it to Elasticsearch.
	for _, _points := range groupedPoints {
		if len(_points) == 0 {
			// This should never happen, but I don't feel like
			// writing points[0] without checking the length first.
			continue
		}

		// We assume that all points have the same timestamp and
		// dimensions because they were grouped by the same key.
		//
		// We use the reference point to get the resource ID and
		// all other information common to all points.
		referencePoint := _points[0]

		// Look up the full cloud resource information in the cache.
		resource := client.LookupResource(referencePoint.ResourceId)

		// Build the event using all the information we have.
		event, err := buildEventFrom(referencePoint, _points, resource, client.Config.DefaultResourceType)
		if err != nil {
			return err
		}

		//
		// Enrich the event with cloud metadata.
		//
		if client.Config.AddCloudMetadata {
			vm := client.GetVMForMetadata(&resource, referencePoint)
			addCloudVMMetadata(&event, vm, resource.Subscription)
		}

		//
		// Report the event to Elasticsearch.
		//
		reporter.Event(event)
	}

	return nil
}
