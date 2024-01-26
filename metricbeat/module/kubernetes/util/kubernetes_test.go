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

package util

import (
	"fmt"
	kubernetes2 "github.com/elastic/beats/v7/libbeat/autodiscover/providers/kubernetes"
	"github.com/elastic/elastic-agent-autodiscover/kubernetes/metadata"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/mapstr"
	"k8s.io/apimachinery/pkg/api/meta"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"testing"
	"time"


	"github.com/elastic/elastic-agent-autodiscover/kubernetes"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/stretchr/testify/require"
	k8sfake "k8s.io/client-go/kubernetes/fake"
)

func TestWatchOptions(t *testing.T) {
	log := logp.NewLogger("test")

	client := k8sfake.NewSimpleClientset()
	config := &kubernetesConfig{
		Namespace:  "test-ns",
		SyncPeriod: time.Minute,
		Node:       "test-node",
	}

	options, err := getWatchOptions(config, false, client, log)
	require.NoError(t, err)
	require.Equal(t, options.SyncTimeout, config.SyncPeriod)
	require.Equal(t, options.Namespace, config.Namespace)
	require.NotEqual(t, options.Node, config.Node)

	options, err = getWatchOptions(config, true, client, log)
	require.NoError(t, err)
	require.Equal(t, options.SyncTimeout, config.SyncPeriod)
	require.Equal(t, options.Namespace, config.Namespace)
	require.Equal(t, options.Node, config.Node)
}

func TestCreateWatcher(t *testing.T) {
	resourceWatchers := NewWatchers()

	client := k8sfake.NewSimpleClientset()
	config := &kubernetesConfig{
		Namespace:  "test-ns",
		SyncPeriod: time.Minute,
		Node:       "test-node",
	}
	log := logp.NewLogger("test")

	options, err := getWatchOptions(config, false, client, log)
	require.NoError(t, err)

	created, err := createWatcher(NamespaceResource, &kubernetes.Node{}, *options, client, resourceWatchers)
	require.True(t, created)
	require.NoError(t, err)

	resourceWatchers.lock.Lock()
	require.Equal(t, 1, len(resourceWatchers.watchersMap))
	require.NotNil(t, resourceWatchers.watchersMap[NamespaceResource])
	require.NotNil(t, resourceWatchers.watchersMap[NamespaceResource].watcher)
	resourceWatchers.lock.Unlock()

	created, err = createWatcher(NamespaceResource, &kubernetes.Namespace{}, *options, client, resourceWatchers)
	require.False(t, created)
	require.NoError(t, err)

	resourceWatchers.lock.Lock()
	require.Equal(t, 1, len(resourceWatchers.watchersMap))
	require.NotNil(t, resourceWatchers.watchersMap[NamespaceResource])
	require.NotNil(t, resourceWatchers.watchersMap[NamespaceResource].watcher)
	resourceWatchers.lock.Unlock()

	created, err = createWatcher(DeploymentResource, &kubernetes.Deployment{}, *options, client, resourceWatchers)
	require.True(t, created)
	require.NoError(t, err)

	resourceWatchers.lock.Lock()
	require.Equal(t, 2, len(resourceWatchers.watchersMap))
	require.NotNil(t, resourceWatchers.watchersMap[DeploymentResource])
	require.NotNil(t, resourceWatchers.watchersMap[NamespaceResource])
	resourceWatchers.lock.Unlock()
}

func TestAddToMetricsetsUsing(t *testing.T) {
	resourceWatchers := NewWatchers()

	client := k8sfake.NewSimpleClientset()
	config := &kubernetesConfig{
		Namespace:  "test-ns",
		SyncPeriod: time.Minute,
		Node:       "test-node",
	}
	log := logp.NewLogger("test")

	options, err := getWatchOptions(config, false, client, log)
	require.NoError(t, err)

	// Create the new entry with watcher and nil string array first
	created, err := createWatcher(DeploymentResource, &kubernetes.Deployment{}, *options, client, resourceWatchers)
	require.True(t, created)
	require.NoError(t, err)

	resourceWatchers.lock.Lock()
	require.NotNil(t, resourceWatchers.watchersMap[DeploymentResource].watcher)
	require.Nil(t, resourceWatchers.watchersMap[DeploymentResource].metricsetsUsing)
	resourceWatchers.lock.Unlock()

	metricsetDeployment := "state_deployment"
	addToMetricsetsUsing(DeploymentResource, metricsetDeployment, resourceWatchers)
	resourceWatchers.lock.Lock()
	require.NotNil(t, resourceWatchers.watchersMap[DeploymentResource].metricsetsUsing)
	require.Equal(t, []string{metricsetDeployment}, resourceWatchers.watchersMap[DeploymentResource].metricsetsUsing)
	resourceWatchers.lock.Unlock()

	metricsetContainer := "container"
	addToMetricsetsUsing(DeploymentResource, metricsetContainer, resourceWatchers)
	resourceWatchers.lock.Lock()
	require.Equal(t, []string{metricsetDeployment, metricsetContainer}, resourceWatchers.watchersMap[DeploymentResource].metricsetsUsing)
	resourceWatchers.lock.Unlock()
}

func TestRemoveFromMetricsetsUsing(t *testing.T) {
	resourceWatchers := NewWatchers()

	client := k8sfake.NewSimpleClientset()
	config := &kubernetesConfig{
		Namespace:  "test-ns",
		SyncPeriod: time.Minute,
		Node:       "test-node",
	}
	log := logp.NewLogger("test")

	options, err := getWatchOptions(config, false, client, log)
	require.NoError(t, err)

	// Create the new entry with watcher and nil string array first
	created, err := createWatcher(DeploymentResource, &kubernetes.Deployment{}, *options, client, resourceWatchers)
	require.True(t, created)
	require.NoError(t, err)

	metricsetDeployment := "state_deployment"
	metricsetPod := "state_pod"
	addToMetricsetsUsing(DeploymentResource, metricsetDeployment, resourceWatchers)
	addToMetricsetsUsing(DeploymentResource, metricsetPod, resourceWatchers)

	resourceWatchers.lock.Lock()
	defer resourceWatchers.lock.Unlock()

	removed, size := removeFromMetricsetsUsing(DeploymentResource, metricsetDeployment, resourceWatchers)
	require.True(t, removed)
	require.Equal(t, 1, size)

	removed, size = removeFromMetricsetsUsing(DeploymentResource, metricsetDeployment, resourceWatchers)
	require.False(t, removed)
	require.Equal(t, 1, size)

	removed, size = removeFromMetricsetsUsing(DeploymentResource, metricsetPod, resourceWatchers)
	require.True(t, removed)
	require.Equal(t, 0, size)
}

func TestCreateAllWatchers(t *testing.T) {
	resourceWatchers := NewWatchers()

	client := k8sfake.NewSimpleClientset()
	config := &kubernetesConfig{
		Namespace:  "test-ns",
		SyncPeriod: time.Minute,
		Node:       "test-node",
		AddResourceMetadata: &metadata.AddResourceMetadataConfig{
			CronJob:    false,
			Deployment: true,
		},
	}
	log := logp.NewLogger("test")

	// Start watchers based on a resource that does not exist should cause an error
	err := createAllWatchers(client, "does-not-exist", "does-not-exist", false, config, log, resourceWatchers)
	require.Error(t, err)
	resourceWatchers.lock.Lock()
	require.Equal(t, 0, len(resourceWatchers.watchersMap))
	resourceWatchers.lock.Unlock()

	// Start watcher for a resource that requires other resources, should start all the watchers
	metricsetPod := "pod"
	extras := getExtraWatchers(PodResource, config)
	err = createAllWatchers(client, metricsetPod, PodResource, false, config, log, resourceWatchers)
	require.NoError(t, err)

	// Check that all the required watchers are in the map
	resourceWatchers.lock.Lock()
	// we add 1 to the expected result to represent the resource itself
	require.Equal(t, len(extras)+1, len(resourceWatchers.watchersMap))
	for _, extra := range extras {
		require.NotNil(t, resourceWatchers.watchersMap[extra])
	}
	resourceWatchers.lock.Unlock()
}

func TestCreateMetaGen(t *testing.T) {
	resourceWatchers := NewWatchers()

	commonMetaConfig := metadata.Config{}
	commonConfig, err := conf.NewConfigFrom(&commonMetaConfig)
	require.NoError(t, err)

	log := logp.NewLogger("test")
	config := &kubernetesConfig{
		Namespace:  "test-ns",
		SyncPeriod: time.Minute,
		Node:       "test-node",
		AddResourceMetadata: &metadata.AddResourceMetadataConfig{
			CronJob:    false,
			Deployment: true,
		},
	}
	client := k8sfake.NewSimpleClientset()

	_, err = createMetadataGen(client, commonConfig, config, DeploymentResource, resourceWatchers)
	// At this point, no watchers were created
	require.Error(t, err)

	// Create the watchers necessary for the metadata generator
	metricsetDeployment := "state_deployment"
	err = createAllWatchers(client, metricsetDeployment, DeploymentResource, false, config, log, resourceWatchers)
	require.NoError(t, err)

	// Create the generators, this time without error
	_, err = createMetadataGen(client, commonConfig, config, DeploymentResource, resourceWatchers)
	require.NoError(t, err)
}

func TestCreateMetaGenSpecific(t *testing.T) {
	resourceWatchers := NewWatchers()

	commonMetaConfig := metadata.Config{}
	commonConfig, err := conf.NewConfigFrom(&commonMetaConfig)
	require.NoError(t, err)

	log := logp.NewLogger("test")
	config := &kubernetesConfig{
		Namespace:  "test-ns",
		SyncPeriod: time.Minute,
		Node:       "test-node",
		AddResourceMetadata: &metadata.AddResourceMetadataConfig{
			CronJob:    false,
			Deployment: true,
		},
	}
	client := k8sfake.NewSimpleClientset()

	// For pod:
	metricsetPod := "pod"

	_, err = createMetadataGenSpecific(client, commonConfig, config, PodResource, resourceWatchers)
	// At this point, no watchers were created
	require.Error(t, err)

	// Create the pod resource + the extras
	err = createAllWatchers(client, metricsetPod, PodResource, false, config, log, resourceWatchers)
	require.NoError(t, err)

	_, err = createMetadataGenSpecific(client, commonConfig, config, PodResource, resourceWatchers)
	require.NoError(t, err)

	// For service:
	_, err = createMetadataGenSpecific(client, commonConfig, config, ServiceResource, resourceWatchers)
	// At this point, no watchers were created
	require.Error(t, err)

	// Create the service resource + the extras
	metricsetService := "state_service"
	err = createAllWatchers(client, metricsetService, ServiceResource, false, config, log, resourceWatchers)
	require.NoError(t, err)

	_, err = createMetadataGenSpecific(client, commonConfig, config, ServiceResource, resourceWatchers)
	require.NoError(t, err)
}

func TestBuildMetadataEnricher_Start_Stop(t *testing.T) {
	resourceWatchers := NewWatchers()

	metricsetNamespace := "state_namespace"
	metricsetDeployment := "state_deployment"

	resourceWatchers.lock.Lock()
	resourceWatchers.watchersMap[NamespaceResource] = &watcherData{
		watcher:         &mockWatcher{},
		started:         false,
		metricsetsUsing: []string{metricsetNamespace, metricsetDeployment},
	}
	resourceWatchers.watchersMap[DeploymentResource] = &watcherData{
		watcher:         &mockWatcher{},
		started:         true,
		metricsetsUsing: []string{metricsetDeployment},
	}
	resourceWatchers.lock.Unlock()

	funcs := mockFuncs{}
	config := &kubernetesConfig{
		Namespace:  "test-ns",
		SyncPeriod: time.Minute,
		Node:       "test-node",
		AddResourceMetadata: &metadata.AddResourceMetadataConfig{
			CronJob:    false,
			Deployment: false,
		},
	}

	enricherNamespace := buildMetadataEnricher(metricsetNamespace, NamespaceResource, resourceWatchers, config,
		funcs.update, funcs.delete, funcs.index)
	resourceWatchers.lock.Lock()
	watcher := resourceWatchers.watchersMap[NamespaceResource]
	require.False(t, watcher.started)
	resourceWatchers.lock.Unlock()

	enricherNamespace.Start(resourceWatchers)
	resourceWatchers.lock.Lock()
	watcher = resourceWatchers.watchersMap[NamespaceResource]
	require.True(t, watcher.started)
	resourceWatchers.lock.Unlock()

	// Stopping should not stop the watcher because it is still being used by deployment metricset
	enricherNamespace.Stop(resourceWatchers)
	resourceWatchers.lock.Lock()
	watcher = resourceWatchers.watchersMap[NamespaceResource]
	require.True(t, watcher.started)
	require.Equal(t, []string{metricsetDeployment}, watcher.metricsetsUsing)
	resourceWatchers.lock.Unlock()

	// Stopping the deployment watcher should stop now both watchers
	enricherDeployment := buildMetadataEnricher(metricsetDeployment, DeploymentResource, resourceWatchers, config,
		funcs.update, funcs.delete, funcs.index)
	enricherDeployment.Stop(resourceWatchers)

	resourceWatchers.lock.Lock()
	watcher = resourceWatchers.watchersMap[NamespaceResource]

	require.False(t, watcher.started)
	require.Equal(t, []string{}, watcher.metricsetsUsing)

	watcher = resourceWatchers.watchersMap[DeploymentResource]
	require.False(t, watcher.started)
	require.Equal(t, []string{}, watcher.metricsetsUsing)

	resourceWatchers.lock.Unlock()

}

func TestBuildMetadataEnricher_Start_Stop_SameResources(t *testing.T) {
	resourceWatchers := NewWatchers()

	metricsetPod := "pod"
	metricsetStatePod := "state_pod"

	resourceWatchers.lock.Lock()
	resourceWatchers.watchersMap[PodResource] = &watcherData{
		watcher:         &mockWatcher{},
		started:         false,
		metricsetsUsing: []string{metricsetStatePod, metricsetPod},
	}
	resourceWatchers.lock.Unlock()

	funcs := mockFuncs{}
	config := &kubernetesConfig{
		Namespace:  "test-ns",
		SyncPeriod: time.Minute,
		Node:       "test-node",
		AddResourceMetadata: &metadata.AddResourceMetadataConfig{
			CronJob:    false,
			Deployment: false,
		},
	}

	enricherPod := buildMetadataEnricher(metricsetPod, PodResource, resourceWatchers, config,
		funcs.update, funcs.delete, funcs.index)
	resourceWatchers.lock.Lock()
	watcher := resourceWatchers.watchersMap[PodResource]
	require.False(t, watcher.started)
	resourceWatchers.lock.Unlock()

	enricherPod.Start(resourceWatchers)
	resourceWatchers.lock.Lock()
	watcher = resourceWatchers.watchersMap[PodResource]
	require.True(t, watcher.started)
	resourceWatchers.lock.Unlock()

	// Stopping should not stop the watcher because it is still being used by state_pod metricset
	enricherPod.Stop(resourceWatchers)
	resourceWatchers.lock.Lock()
	watcher = resourceWatchers.watchersMap[PodResource]
	require.True(t, watcher.started)
	require.Equal(t, []string{metricsetStatePod}, watcher.metricsetsUsing)
	resourceWatchers.lock.Unlock()

	// Stopping the state_pod watcher should stop pod watcher
	enricherStatePod := buildMetadataEnricher(metricsetStatePod, PodResource, resourceWatchers, config,
		funcs.update, funcs.delete, funcs.index)
	enricherStatePod.Stop(resourceWatchers)

	resourceWatchers.lock.Lock()
	watcher = resourceWatchers.watchersMap[PodResource]
	require.False(t, watcher.started)
	require.Equal(t, []string{}, watcher.metricsetsUsing)
	resourceWatchers.lock.Unlock()
}

/*
func TestBuildMetadataEnricher_EventHandler(t *testing.T) {
	resourceWatchers := NewWatchers()

	resourceWatchers.lock.Lock()
	resourceWatchers.watchersMap[PodResource] = &watcherData{
		watcher:         &mockWatcher{},
		started:         false,
		metricsetsUsing: []string{PodResource},
	}
	resourceWatchers.lock.Unlock()

	funcs := mockFuncs{}
	resource := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			UID:  types.UID("mockuid"),
			Name: "enrich",
			Labels: map[string]string{
				"label": "value",
			},
			Namespace: "default",
		},
	}

	config := &kubernetesConfig{
		Namespace:  "test-ns",
		SyncPeriod: time.Minute,
		Node:       "test-node",
		AddResourceMetadata: &metadata.AddResourceMetadataConfig{
			CronJob:    false,
			Deployment: false,
		},
	}

	enricher := buildMetadataEnricher(PodResource, resourceWatchers, config, funcs.update, funcs.delete, funcs.index)
	resourceWatchers.lock.Lock()
	wData := resourceWatchers.watchersMap[PodResource]
	mockW := wData.watcher.(*mockWatcher)
	require.NotNil(t, mockW.handler)
	resourceWatchers.lock.Unlock()

	enricher.Start(resourceWatchers)
	resourceWatchers.lock.Lock()
	watcher := resourceWatchers.watchersMap[PodResource]
	require.True(t, watcher.started)
	resourceWatchers.lock.Unlock()

	resourceWatchers.lock.Lock()
	wData = resourceWatchers.watchersMap[PodResource]
	mockW = wData.watcher.(*mockWatcher)
	mockW.handler.OnAdd(resource)
	resourceWatchers.lock.Unlock()

	require.Equal(t, resource, funcs.updated)

	// Test enricher
	events := []mapstr.M{
		{"name": "unknown"},
		{"name": "enrich"},
	}
	enricher.Enrich(events)

	require.Equal(t, []mapstr.M{
		{"name": "unknown"},
		{
			"name":    "enrich",
			"_module": mapstr.M{"label": "value", "pod": mapstr.M{"name": "enrich", "uid": "mockuid"}},
			"meta":    mapstr.M{"orchestrator": mapstr.M{"cluster": mapstr.M{"name": "gke-4242"}}},
		},
	}, events)

	// Enrich a pod (metadata goes in root level)
	events = []mapstr.M{
		{"name": "unknown"},
		{"name": "enrich"},
	}
	enricher.isPod = true
	enricher.Enrich(events)

	require.Equal(t, []mapstr.M{
		{"name": "unknown"},
		{
			"name":    "enrich",
			"uid":     "mockuid",
			"_module": mapstr.M{"label": "value"},
			"meta":    mapstr.M{"orchestrator": mapstr.M{"cluster": mapstr.M{"name": "gke-4242"}}},
		},
	}, events)

	// Emit delete event
	resourceWatchers.lock.Lock()
	wData = resourceWatchers.watchersMap[PodResource]
	mockW = wData.watcher.(*mockWatcher)
	mockW.handler.OnDelete(resource)
	resourceWatchers.lock.Unlock()

	require.Equal(t, resource, funcs.deleted)

	events = []mapstr.M{
		{"name": "unknown"},
		{"name": "enrich"},
	}
	enricher.Enrich(events)

	require.Equal(t, []mapstr.M{
		{"name": "unknown"},
		{"name": "enrich"},
	}, events)

	enricher.Stop(resourceWatchers)
	resourceWatchers.lock.Lock()
	watcher = resourceWatchers.watchersMap[PodResource]
	require.False(t, watcher.started)
	resourceWatchers.lock.Unlock()
}
*/

type mockFuncs struct {
	updated kubernetes.Resource
	deleted kubernetes.Resource
	indexed mapstr.M
}

func (f *mockFuncs) update(m map[string]mapstr.M, obj kubernetes.Resource) {
	accessor, _ := meta.Accessor(obj)
	f.updated = obj
	meta := mapstr.M{
		"kubernetes": mapstr.M{
			"pod": mapstr.M{
				"name": accessor.GetName(),
				"uid":  string(accessor.GetUID()),
			},
		},
	}
	logger := logp.NewLogger("kubernetes")
	for k, v := range accessor.GetLabels() {
		kubernetes2.ShouldPut(meta, fmt.Sprintf("kubernetes.%v", k), v, logger)
	}
	kubernetes2.ShouldPut(meta, "orchestrator.cluster.name", "gke-4242", logger)
	m[accessor.GetName()] = meta
}

func (f *mockFuncs) delete(m map[string]mapstr.M, obj kubernetes.Resource) {
	accessor, _ := meta.Accessor(obj)
	f.deleted = obj
	delete(m, accessor.GetName())
}

func (f *mockFuncs) index(m mapstr.M) string {
	f.indexed = m
	return m["name"].(string)
}

type mockWatcher struct {
	handler kubernetes.ResourceEventHandler
}

func (m *mockWatcher) Start() error {
	return nil
}

func (m *mockWatcher) Stop() {

}

func (m *mockWatcher) AddEventHandler(r kubernetes.ResourceEventHandler) {
	m.handler = r
}

func (m *mockWatcher) Store() cache.Store {
	return nil
}

func (m *mockWatcher) Client() k8s.Interface {
	return nil
}

func (m *mockWatcher) CachedObject() runtime.Object {
	return nil
}
