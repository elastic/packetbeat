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

package licenser

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs/elasticsearch"
)

// OSSLicense default license to use.
var (
	OSSLicense = &License{
		UUID:   uuid.NewV4(),
		Type:   OSS,
		Mode:   OSS,
		Status: Active,
		Features: features{
			Graph:      graph{},
			Logstash:   logstash{},
			ML:         ml{},
			Monitoring: monitoring{},
			Rollup:     rollup{},
			Security:   security{},
			Watcher:    watcher{},
		},
	}
)

// Watcher allows a type to receive a new event when a new license is received.
type Watcher interface {
	OnNewLicense(license License)
	OnManagerStopped()
}

// Fetcher interface implements the mechanism to retrieve a License. Currently we only
// support license coming from the '/_xpack' rest api.
type Fetcher interface {
	Fetch() (*License, error)
}

// Errors returned by the manager.
var (
	ErrWatcherAlreadyExist = errors.New("watcher already exist")
	ErrWatcherDoesntExist  = errors.New("watcher doesn't exist")

	ErrManagerStopped = errors.New("license manager is stopped")
	ErrNoLicenseFound = errors.New("no license found")
)

// Backoff values when the remote cluster is not responding.
var (
	maxBackoff  = time.Duration(60)
	initBackoff = time.Duration(5)
	jitterCap   = 1000 // 1000 milliseconds
)

// Manager keeps tracks of license management, it uses a fetcher usually the ElasticFetcher to
// retrieve a licence from a specific cluster.
//
// Starting the manager will start a go routine to periodically query the license fetcher.
// if an error occur on the fetcher we will retry until we successfully
// receive a new license. During that period we start a grace counter, we assume the license is
// still valid during the grace period, when this period expire we will keep retrying but the previous
// license will be invalidated and we will fallback to the OSS license.
//
// Retrieving the current license:
// - Call the `Get()` on the manager instance.
// - Or register a `Watcher` with the manager to receive the new license and acts on it, you will
// also receive an event when the Manager is stopped.
//
//
// Notes:
// - When the manager is started no license is set by default.
// - When a license is invalidated, we fallback to the OSS License and the watchers get notified.
// - Adding a watcher will automatically send the current license to the newly added watcher if
//   available.
type Manager struct {
	done chan struct{}
	sync.RWMutex
	wg          sync.WaitGroup
	fetcher     Fetcher
	duration    time.Duration
	gracePeriod time.Duration
	license     *License
	watchers    map[Watcher]Watcher
	log         *logp.Logger
}

// New takes an elasticsearch client and wraps it into a fetcher, the fetch will handle the JSON
// and response code from the cluster.
func New(client *elasticsearch.Client, duration time.Duration, gracePeriod time.Duration) *Manager {
	fetcher := NewElasticFetcher(client)
	return NewWithFetcher(fetcher, duration, gracePeriod)
}

// NewWithFetcher takes a fetcher and return a license manager.
func NewWithFetcher(fetcher Fetcher, duration time.Duration, gracePeriod time.Duration) *Manager {
	m := &Manager{
		fetcher:     fetcher,
		duration:    duration,
		log:         logp.NewLogger("license-manager"),
		done:        make(chan struct{}),
		gracePeriod: gracePeriod,
		watchers:    make(map[Watcher]Watcher),
	}

	return m
}

// AddWatcher register a new watcher to receive events when the license is retrieved or when the manager
// is closed.
func (m *Manager) AddWatcher(watcher Watcher) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.watchers[watcher]; ok {
		return ErrWatcherAlreadyExist
	}

	m.watchers[watcher] = watcher

	// when we register a new watchers send the current license unless we did not retrieve it.
	if m.license != nil {
		watcher.OnNewLicense(*m.license)
	}
	return nil
}

// RemoveWatcher removes the watcher if it exist or return an error.
func (m *Manager) RemoveWatcher(watcher Watcher) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.watchers[watcher]; ok {
		delete(m.watchers, watcher)
		return nil
	}
	return ErrWatcherDoesntExist
}

// Get return the current active license, it can return an error if the manager is stopped or when
// there is no license in the manager, Instead of querying the Manager it is easier to register a
// watcher to listen to license change.
func (m *Manager) Get() (*License, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-m.done:
		return nil, ErrManagerStopped
	default:
		if m.license == nil {
			return nil, ErrNoLicenseFound
		}
		return m.license, nil
	}
}

// Start starts the License manager, the manager will start a go routine to periodically
// retrieve the license from the fetcher.
func (m *Manager) Start() {
	// First update should be in sync at startup to ensure a
	// consistent state.
	m.log.Info("license manager started, no license found.")
	m.wg.Add(1)
	go m.worker()
}

// Stop terminates the license manager, the go routine will be stopped and the cached license will
// be removed and no more checks can be done on the manager.
func (m *Manager) Stop() {
	select {
	case <-m.done:
		m.log.Error("license manager already stopped")
	default:
	}

	defer m.log.Info("license manager stopped")
	defer m.notify(func(w Watcher) {
		w.OnManagerStopped()
	})

	// stop the periodic check license and wait for it to complete
	close(m.done)
	m.wg.Wait()

	// invalidate current license
	m.Lock()
	defer m.Unlock()
	m.license = nil
}

func (m *Manager) notify(op func(Watcher)) {
	m.RLock()
	defer m.RUnlock()

	if len(m.watchers) == 0 {
		m.log.Debugf("no watchers configured")
		return
	}

	m.log.Debugf("notifying %d watchers", len(m.watchers))
	for _, w := range m.watchers {
		op(w)
	}
}

func (m *Manager) worker() {
	defer m.wg.Done()
	m.log.Debug("starting periodic license check")
	defer m.log.Debug("periodic license check is stopped")

	jitter := rand.Intn(jitterCap)

	// Add some jitter to space requests from a large fleet of beats.
	select {
	case <-time.After(time.Duration(jitter) * time.Millisecond):
	}

	// eager initial check.
	m.update()

	// periodically checks license.
	for {
		select {
		case <-m.done:
			return
		case <-time.After(m.duration):
			m.log.Debug("license is too old, updating, grace period: %s", m.gracePeriod)
			m.update()
		}
	}
}

func (m *Manager) update() {
	backoff := common.NewBackoff(m.done, initBackoff, maxBackoff)
	startedAt := time.Now()
	for {
		select {
		case <-m.done:
			return
		default:
			license, err := m.fetcher.Fetch()
			if err != nil {
				m.log.Info("cannot retrieve license, retrying later, error: %s", err)

				// check if the license is still in the grace period.
				// permit some operations if the license could not be checked
				// right away. This is to smooth any networks problems.
				if grace := time.Now().Sub(startedAt); grace > m.gracePeriod {
					m.log.Info("grace period expired, invalidating license")
					m.invalidate()
				} else {
					m.log.Debugf("license is too old, grace time remaining: %s", m.gracePeriod-grace)
				}

				backoff.Wait()
				continue
			}

			// we have a valid license, notify watchers and sleep until next check.
			m.log.Info(
				"valid license retrieved, license mode: %s, type: %s, status: %s",
				license.Get(),
				license.Type,
				license.Status,
			)
			m.saveAndNotify(license)
			return
		}
	}
}

func (m *Manager) saveAndNotify(license *License) {
	if !m.save(license) {
		return
	}

	l := *license
	m.notify(func(w Watcher) {
		w.OnNewLicense(l)
	})
}

func (m *Manager) save(license *License) bool {
	m.Lock()
	defer m.Unlock()

	// License didn't change no need to notify watchers.
	if m.license != nil && m.license.EqualTo(license) {
		return false
	}
	defer m.log.Debug("license information updated")

	m.license = license
	return true
}

func (m *Manager) invalidate() {
	defer m.log.Debug("invalidate cached license, fallback to OSS")
	m.saveAndNotify(OSSLicense)
}
