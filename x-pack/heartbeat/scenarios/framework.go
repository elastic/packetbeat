// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package scenarios

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/elastic/beats/v7/heartbeat/monitors/plugin"
	"github.com/elastic/beats/v7/heartbeat/monitors/stdfields"
	"github.com/elastic/beats/v7/heartbeat/monitors/wrappers/monitorstate"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"

	"github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/elastic-agent-libs/monitoring"

	hbconfig "github.com/elastic/beats/v7/heartbeat/config"
	"github.com/elastic/beats/v7/heartbeat/monitors"
	"github.com/elastic/beats/v7/heartbeat/scheduler"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/publisher/pipeline"
	beatversion "github.com/elastic/beats/v7/libbeat/version"
)

type ScenarioRun func(t *testing.T) (config mapstr.M, close func(), err error)

type Scenario struct {
	Name      string
	Type      string
	Runner    ScenarioRun
	Tags      []string
	Location  *hbconfig.LocationWithID
	ESEnabled bool
}

type Twist func(Scenario) Scenario

func MakeTwist(name string, fn Twist) Twist {
	return func(s Scenario) Scenario {
		newS := s.clone()
		newS.Name = fmt.Sprintf("%s>Twist(%s)", s.Name, name)
		return fn(newS)
	}
}

func (s Scenario) clone() Scenario {
	copy := s
	if s.Location != nil {
		locationCopy := *s.Location
		copy.Location = &locationCopy
	}
	return copy
}

func (s Scenario) Run(t *testing.T, twist Twist, callback func(t *testing.T, mtr *MonitorTestRun, err error)) {
	runS := s
	if twist != nil {
		runS = twist(s.clone())
	}

	cfgMap, rClose, err := runS.Runner(t)
	if rClose != nil {
		defer rClose()
	}
	if err != nil {
		callback(t, nil, err)
		return
	}

	stateLoader := monitorstate.NilStateLoader
	if s.ESEnabled {
		stateLoader = monitorstate.IntegESLoader(t, "heartbeat-*,synthetics-*", TestLocationDefault)
	}

	t.Run(runS.Name, func(t *testing.T) {
		t.Parallel()
		mtr, err := runMonitorOnce(t, cfgMap, runS.Location, stateLoader)
		mtr.Wait()
		callback(t, mtr, err)
		mtr.Close()
	})
}

type ScenarioDB struct {
	All      []Scenario
	ByTag    map[string][]Scenario
	initOnce *sync.Once
}

func (sdb *ScenarioDB) Init() {
	var prunedList []Scenario
	browserCapable := os.Getenv("ELASTIC_SYNTHETICS_CAPABLE") == "true"
	icmpCapable := os.Getenv("ELASTIC_ICMP_CAPABLE") == "true"
	sdb.initOnce.Do(func() {
		for _, s := range sdb.All {
			if s.Type == "browser" && !browserCapable {
				continue
			}
			if s.Type == "icmp" && !icmpCapable {
				continue
			}
			prunedList = append(prunedList, s)

			for _, t := range s.Tags {
				sdb.ByTag[t] = append(sdb.ByTag[t], s)
			}
		}
	})
	sdb.All = prunedList
}

func (sdb *ScenarioDB) Add(s ...Scenario) {
	sdb.All = append(sdb.All, s...)
}

func (sdb *ScenarioDB) RunAll(t *testing.T, callback func(*testing.T, *MonitorTestRun, error)) {
	sdb.RunAllWithATwist(t, nil, callback)
}

func (sdb *ScenarioDB) RunAllWithATwist(t *testing.T, twist Twist, callback func(*testing.T, *MonitorTestRun, error)) {
	sdb.Init()
	for _, s := range sdb.All {
		s.Run(t, twist, callback)
	}
}

func (sdb *ScenarioDB) RunTag(t *testing.T, tagName string, callback func(*testing.T, *MonitorTestRun, error)) {
	sdb.RunTagWithATwist(t, tagName, nil, callback)
}

func (sdb *ScenarioDB) RunTagWithATwist(t *testing.T, tagName string, twist Twist, callback func(*testing.T, *MonitorTestRun, error)) {
	sdb.Init()
	if len(sdb.ByTag[tagName]) < 1 {
		require.Failf(t, "no scenarios have tags matching %s", tagName)
	}
	for _, s := range sdb.ByTag[tagName] {
		s.Run(t, twist, callback)
	}
}

type MonitorTestRun struct {
	StdFields stdfields.StdMonitorFields
	Config    mapstr.M
	Monitor   *monitors.Monitor
	Events    func() []*beat.Event
	Wait      func()
	Close     func()
}

func runMonitorOnce(t *testing.T, monitorConfig mapstr.M, location *hbconfig.LocationWithID, stateLoader monitorstate.StateLoader) (mtr *MonitorTestRun, err error) {
	mtr = &MonitorTestRun{
		Config:    monitorConfig,
		StdFields: stdfields.StdMonitorFields{},
	}

	// make a pipeline
	pipe := &monitors.MockPipeline{}
	// pass it to the factory
	f, sched, closeFactory := setupFactoryAndSched(location, stateLoader)
	conf, err := config.NewConfigFrom(monitorConfig)
	require.NoError(t, err)
	err = conf.Unpack(&mtr.StdFields)
	require.NoError(t, err)

	mIface, err := f.Create(pipe, conf)
	require.NoError(t, err)
	mtr.Monitor = mIface.(*monitors.Monitor)
	require.NotNil(t, mtr.Monitor, "could not convert to monitor %v", mIface)
	mtr.Events = pipe.PublishedEvents

	// start the monitor
	mtr.Monitor.Start()
	mtr.Wait = func() {
		// wait for the monitor to stop
		sched.WaitForRunOnce()
		// stop the monitor itself
		mtr.Monitor.Stop()
		closeFactory()
	}
	mtr.Close = closeFactory
	return mtr, err
}

func setupFactoryAndSched(location *hbconfig.LocationWithID, stateLoader monitorstate.StateLoader) (factory *monitors.RunnerFactory, sched *scheduler.Scheduler, close func()) {
	id, _ := uuid.NewV4()
	eid, _ := uuid.NewV4()
	info := beat.Info{
		Beat:            "heartbeat",
		IndexPrefix:     "heartbeat",
		Version:         beatversion.GetDefaultVersion(),
		ElasticLicensed: true,
		Name:            "heartbeat",
		Hostname:        "localhost",
		ID:              id,
		EphemeralID:     eid,
		FirstStart:      time.Now(),
		StartTime:       time.Now(),
		Monitoring: struct {
			DefaultUsername string
		}{
			DefaultUsername: "test",
		},
	}

	sched = scheduler.Create(
		1,
		monitoring.NewRegistry(),
		time.Local,
		nil,
		true,
	)

	if stateLoader == nil {
		stateLoader = monitorstate.NilStateLoader
	}

	return monitors.NewFactory(monitors.FactoryParams{
			BeatInfo:    info,
			AddTask:     sched.Add,
			StateLoader: stateLoader,
			PluginsReg:  plugin.GlobalPluginsReg,
			PipelineClientFactory: func(pipeline beat.Pipeline) (pipeline.ISyncClient, error) {
				c, _ := pipeline.Connect()
				return monitors.SyncPipelineClientAdaptor{C: c}, nil
			},
			BeatLocation: location,
		}),
		sched,
		sched.Stop
}
