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

package kubernetes

import (
<<<<<<< HEAD
=======
	"fmt"
	httpnet "net/http"
	"strings"
>>>>>>> 2491f08379 ([metricbeat][Kubernetes] Update token on the first HTTP 401 when making requests to kubelet API (#40636))
	"sync"
	"time"

	"github.com/mitchellh/hashstructure"
	"github.com/pkg/errors"
	dto "github.com/prometheus/client_model/go"

	"github.com/elastic/beats/v7/metricbeat/helper"
	p "github.com/elastic/beats/v7/metricbeat/helper/prometheus"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/module/kubernetes/util"
)

func init() {
	// Register the ModuleFactory function for the "kubernetes" module.
	if err := mb.Registry.AddModule("kubernetes", ModuleBuilder()); err != nil {
		panic(err)
	}
}

type Module interface {
	mb.Module
	GetStateMetricsFamilies(prometheus p.Prometheus) ([]*dto.MetricFamily, error)
	GetKubeletStats(http *helper.HTTP) ([]byte, error)
	GetMetricsRepo() *util.MetricsRepo
}

type familiesCache struct {
	sharedFamilies     []*dto.MetricFamily
	lastFetchErr       error
	lastFetchTimestamp time.Time
}

type kubeStateMetricsCache struct {
	cacheMap map[uint64]*familiesCache
	lock     sync.Mutex
}

func (c *kubeStateMetricsCache) getCacheMapEntry(hash uint64) *familiesCache {
	if _, ok := c.cacheMap[hash]; !ok {
		c.cacheMap[hash] = &familiesCache{}
	}
	return c.cacheMap[hash]
}

type statsCache struct {
	sharedStats        []byte
	lastFetchErr       error
	lastFetchTimestamp time.Time
}

type kubeletStatsCache struct {
	cacheMap map[uint64]*statsCache
	lock     sync.Mutex
}

func (c *kubeletStatsCache) getCacheMapEntry(hash uint64) *statsCache {
	if _, ok := c.cacheMap[hash]; !ok {
		c.cacheMap[hash] = &statsCache{}
	}
	return c.cacheMap[hash]
}

type module struct {
	mb.BaseModule

	kubeStateMetricsCache *kubeStateMetricsCache
	kubeletStatsCache     *kubeletStatsCache
	metricsRepo           *util.MetricsRepo
	cacheHash             uint64
}

func ModuleBuilder() func(base mb.BaseModule) (mb.Module, error) {
	kubeStateMetricsCache := &kubeStateMetricsCache{
		cacheMap: make(map[uint64]*familiesCache),
	}
	kubeletStatsCache := &kubeletStatsCache{
		cacheMap: make(map[uint64]*statsCache),
	}
	metricsRepo := util.NewMetricsRepo()
	return func(base mb.BaseModule) (mb.Module, error) {
		hash, err := generateCacheHash(base.Config().Hosts)
		if err != nil {
			return nil, errors.Wrap(err, "error generating cache hash for kubeStateMetricsCache")
		}

		m := module{
			BaseModule:            base,
			kubeStateMetricsCache: kubeStateMetricsCache,
			kubeletStatsCache:     kubeletStatsCache,
			metricsRepo:           metricsRepo,
			cacheHash:             hash,
		}
		return &m, nil
	}
}

func (m *module) GetStateMetricsFamilies(prometheus p.Prometheus) ([]*dto.MetricFamily, error) {
	m.kubeStateMetricsCache.lock.Lock()
	defer m.kubeStateMetricsCache.lock.Unlock()

	now := time.Now()
	// NOTE: These entries will never be removed, this can be a leak if
	// metricbeat is used to monitor clusters dynamically created.
	// (https://github.com/elastic/beats/pull/25640#discussion_r633395213)
	familiesCache := m.kubeStateMetricsCache.getCacheMapEntry(m.cacheHash)

	if familiesCache.lastFetchTimestamp.IsZero() || now.Sub(familiesCache.lastFetchTimestamp) > m.Config().Period {
		familiesCache.sharedFamilies, familiesCache.lastFetchErr = prometheus.GetFamilies()
		familiesCache.lastFetchTimestamp = now
	}

	return familiesCache.sharedFamilies, familiesCache.lastFetchErr
}

func (m *module) GetKubeletStats(http *helper.HTTP) ([]byte, error) {
	m.kubeletStatsCache.lock.Lock()
	defer m.kubeletStatsCache.lock.Unlock()

	now := time.Now()

	// NOTE: These entries will never be removed, this can be a leak if
	// metricbeat is used to monitor clusters dynamically created.
	// (https://github.com/elastic/beats/pull/25640#discussion_r633395213)
	statsCache := m.kubeletStatsCache.getCacheMapEntry(m.cacheHash)

	// Check if the last time we tried to make a request to the Kubelet API ended in a 401 Unauthorized error.
	// If this is the case, we should not keep making requests.
	errorUnauthorisedMsg := fmt.Sprintf("HTTP error %d", httpnet.StatusUnauthorized)
	if statsCache.lastFetchErr != nil && strings.Contains(statsCache.lastFetchErr.Error(), errorUnauthorisedMsg) {
		return statsCache.sharedStats, statsCache.lastFetchErr
	}

	// If this is the first request, or it has passed more time than config.period, we should
	// make a request to the Kubelet API again to get the last metrics' values.
	if statsCache.lastFetchTimestamp.IsZero() || now.Sub(statsCache.lastFetchTimestamp) > m.Config().Period {
		statsCache.sharedStats, statsCache.lastFetchErr = http.FetchContent()

		// If we got an unauthorized error from our HTTP request, it is possible the token has expired.
		// We should update the Authorization header in that case. We only try this for the first time
		// we get HTTP 401 to avoid getting in a loop in case the cause of the error is something different.
		if statsCache.lastFetchErr != nil && strings.Contains(statsCache.lastFetchErr.Error(), errorUnauthorisedMsg) {
			if _, err := http.RefreshAuthorizationHeader(); err == nil {
				statsCache.sharedStats, statsCache.lastFetchErr = http.FetchContent()
			}
		}

		statsCache.lastFetchTimestamp = now
	}

	return statsCache.sharedStats, statsCache.lastFetchErr
}

func generateCacheHash(host []string) (uint64, error) {
	id, err := hashstructure.Hash(host, nil)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *module) GetMetricsRepo() *util.MetricsRepo {
	return m.metricsRepo
}
