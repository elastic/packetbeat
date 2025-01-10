// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build linux && (amd64 || arm64) && cgo

package process

import (
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/elastic/beats/v7/auditbeat/ab"
	"github.com/elastic/beats/v7/auditbeat/helper/tty"
	"github.com/elastic/beats/v7/libbeat/common/capabilities"
	"github.com/elastic/beats/v7/metricbeat/mb"
	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
	"github.com/elastic/beats/v7/x-pack/auditbeat/module/system"
	"github.com/elastic/elastic-agent-libs/mapstr"
	quark "github.com/elastic/go-quark"
	"github.com/stretchr/testify/require"
)

type backend int

const (
	Ebpf backend = iota
	Kprobe
)

func TestInitialSnapshotEbpf(t *testing.T) {
	testInitialSnapshot(t, Ebpf)
}

func TestInitialSnapshotKprobe(t *testing.T) {
	testInitialSnapshot(t, Kprobe)
}

func TestForkExecExitEbpf(t *testing.T) {
	testForkExecExit(t, Ebpf)
}

func TestForkExecExitKprobe(t *testing.T) {
	testForkExecExit(t, Kprobe)
}

func TestQuarkMetricSetEbpf(t *testing.T) {
	testQuarkMetricSet(t, Ebpf)
}

func TestQuarkMetricSetKprobe(t *testing.T) {
	testQuarkMetricSet(t, Kprobe)
}

// testInitialSnapshot see if quark is generating snapshot events
func testInitialSnapshot(t *testing.T, be backend) {
	qq := openQueue(t, be)
	defer qq.Close()

	// There should be events of kind quark.QUARK_EV_SNAPSHOT
	qevs := drainFor(t, qq, 5*time.Millisecond)
	var gotsnap bool
	for _, qev := range qevs {
		if qev.Events&quark.QUARK_EV_SNAPSHOT != 0 {
			gotsnap = true
		}
	}

	require.True(t, gotsnap)
}

// testForkExecExit tests if a spawned process shows up in quark
func testForkExecExit(t *testing.T, be backend) {
	qq := openQueue(t, be)
	defer qq.Close()

	// runNop will fork+exec+exit /bin/true
	cmd := runNop(t)
	qev := drainFirstOfPid(t, qq, cmd.Process.Pid)

	// We should get at least FORK|EXEC|EXIT in the aggregation
	require.Equal(t,
		qev.Events&(quark.QUARK_EV_FORK|quark.QUARK_EV_EXEC|quark.QUARK_EV_EXIT),
		quark.QUARK_EV_FORK|quark.QUARK_EV_EXEC|quark.QUARK_EV_EXIT)

	// This is virtually impossible to fail, but we're pedantic
	require.True(t, qev.Process.Proc.Valid)

	// We need these otherwise nothing works
	require.NotZero(t, qev.Process.Proc.MntInonum)
	require.NotZero(t, qev.Process.Proc.TimeBoot)
	require.NotZero(t, qev.Process.Proc.Ppid)

	// Must be /bin/true
	require.Equal(t, qev.Process.Filename, cmd.Path)
	require.Equal(t, qev.Process.Filename, cmd.Args[0])

	// Kprobe cwd path depth is limited
	if be != Kprobe {
		cwd, err := os.Getwd()
		require.NoError(t, err)

		require.Equal(t, cwd, qev.Process.Cwd)
	}

	// Check exit
	require.True(t, qev.Process.Exit.Valid)
	require.Zero(t, qev.Process.Exit.ExitCode)
	// Don't care about ExitTime, it's also not precise
}

// testQuarkMetricSet will start the module and check if the it
// generates the correct event for os.Getpid() (an existing process),
// and for a process we spawn ourselves via runNop().
func testQuarkMetricSet(t *testing.T, be backend) {
	config := getConfigForQuark(be)

	// Start the module, it will open its own queue
	f := mbtest.NewPushMetricSetV2WithRegistry(t, config, ab.Registry)
	ms, ok := f.(*QuarkMetricSet)
	require.True(t, ok)

	// Start our own queue in parallel, so we can compare some
	// members we have no other way of fetching
	qq := openQueue(t, be)
	defer qq.Close()

	// The queue is open, so we can spawn something and it should show up
	cmd := runNop(t)

	// Run the main loop, it should get a snapshot event of
	// and a fork+exec+exit of the nop we just ran
	// XXX sadly we can't control the holdTime of the queue used by beats,
	// so we need to wait a full second
	events := mbtest.RunPushMetricSetV2(1100*time.Millisecond, 0, ms)
	require.NotEmpty(t, events)

	// Lookup self from qq, we need Proc.TimeBoot and Suid/Sgid
	selfFromQQ, ok := qq.Lookup(os.Getpid())
	require.True(t, ok)
	// Make a fake event from self (Os.getpid())
	selfTarget := makeSelfEvent(t, selfFromQQ, be)
	// Lookup what we actually got from beats
	selfActual := firstEventOfPid(t, events, os.Getpid())
	// Compare
	checkEvent(t, selfTarget, selfActual)

	// Drain until we find the event generated by runNop(), we
	// need Proc.TimeBoot and Suid/Sgid
	spawnedFromQQ := drainFirstOfPid(t, qq, cmd.Process.Pid)
	// Make an event of the spawned cmd
	spawnedTarget := makeEventOfCmd(t, cmd, spawnedFromQQ, be)
	// Lookup what we actually got from beats
	spawnedActual := firstEventOfPid(t, events, cmd.Process.Pid)
	// Compare
	checkEvent(t, spawnedTarget, spawnedActual)
}

// checkEvent checks the equality of all the events from target in actual,
// not the other way around since actual is always larger
func checkEvent(t *testing.T, target mb.Event, actual mb.Event) {
	for tk, tv := range target.RootFields {
		av, err := actual.RootFields.GetValue(tk)
		require.NoError(t, err)
		require.Equal(t, tv, av)
	}
}

// openQueue opens a quark queue on a specific backend
func openQueue(t *testing.T, be backend) *quark.Queue {
	attr := quark.DefaultQueueAttr()
	attr.HoldTime = 25
	attr.Flags &= ^quark.QQ_ALL_BACKENDS
	if be == Ebpf {
		attr.Flags |= quark.QQ_EBPF
	} else if be == Kprobe {
		attr.Flags |= quark.QQ_KPROBE
	}
	qq, err := quark.OpenQueue(attr, 1)
	require.NoError(t, err)

	return qq
}

// runNop does fork+exec+exit /bin/true
func runNop(t *testing.T) *exec.Cmd {
	cmd := exec.Command("/bin/true")
	require.NotNil(t, cmd)
	err := cmd.Run()
	require.NoError(t, err)

	return cmd
}

// drainFor drains all events for `d`
func drainFor(t *testing.T, qq *quark.Queue, d time.Duration) []quark.Event {
	var allQevs []quark.Event

	start := time.Now()

	for {
		qevs, err := qq.GetEvents()
		require.NoError(t, err)
		for _, qev := range qevs {
			if !wantedEvent(qev) {
				continue
			}
			allQevs = append(allQevs, qev)
		}
		if time.Since(start) > d {
			break
		}
		// Intentionally placed at the end so that we always
		// get one more try after the last block
		if len(qevs) == 0 {
			qq.Block()
		}
	}

	return allQevs
}

// drainFirstOfPid returns the first event
func drainFirstOfPid(t *testing.T, qq *quark.Queue, pid int) quark.Event {
	start := time.Now()

	for {
		qevs, err := qq.GetEvents()
		require.NoError(t, err)
		for _, qev := range qevs {
			if !wantedEvent(qev) {
				continue
			}
			if qev.Process.Pid == uint32(pid) {
				return qev
			}
		}
		if time.Since(start) > time.Second {
			break
		}
		// Intentionally placed at the end so that we always
		// get one more try after the last block
		if len(qevs) == 0 {
			qq.Block()
		}
	}

	t.Fatalf("Can't find event of pid %d", pid)

	return quark.Event{} // NOTREACHED
}

// firstEventOfPid looks up the first event of `pid` in `events`
func firstEventOfPid(t *testing.T, events []mb.Event, pid int) mb.Event {
	for _, event := range events {
		pid2, err := event.RootFields.GetValue("process.pid")
		require.NoError(t, err)
		if pid2.(uint32) == uint32(pid) {
			return event
		}
	}

	t.Fatalf("Can't find event of pid %d", pid)

	return mb.Event{} // NOTREACHED
}

// makeSelfEvent builds what should be the event that quark will
// generate as a initial snapshot of the current process
func makeSelfEvent(t *testing.T, qp quark.Process, be backend) mb.Event {
	exe, err := os.Executable()
	require.NoError(t, err)

	interactive := tty.InteractiveFromTTY(tty.TTYDev{
		Major: qp.Proc.TtyMajor,
		Minor: qp.Proc.TtyMinor,
	})

	capEff, err := capabilities.FromPid(capabilities.Effective, os.Getpid())
	require.NoError(t, err)
	capPer, err := capabilities.FromPid(capabilities.Permitted, os.Getpid())
	require.NoError(t, err)

	self := mb.Event{
		RootFields: mapstr.M{
			"event.type":                            []string{"info"},
			"event.action":                          "existing_process",
			"event.category":                        []string{"process"},
			"event.kind":                            "event",
			"process.name":                          qp.Comm,
			"process.args":                          qp.Cmdline,
			"process.args_count":                    len(qp.Cmdline),
			"process.pid":                           uint32(os.Getpid()),
			"process.executable":                    exe,
			"process.parent.pid":                    uint32(os.Getppid()),
			"process.start":                         time.Unix(0, int64(qp.Proc.TimeBoot)),
			"user.id":                               uint32(0),
			"user.group.id":                         uint32(0),
			"user.effective.id":                     uint32(0),
			"user.saved.id":                         qp.Proc.Suid,
			"user.saved.group.id":                   qp.Proc.Sgid,
			"user.name":                             "root",
			"user.group.name":                       "root",
			"process.interactive":                   interactive,
			"process.thread.capabilities.effective": capEff,
			"process.thread.capabilities.permitted": capPer,
		},
	}

	// Kprobe path depth is limited
	if be != Kprobe {
		cwd, err := os.Getwd()
		require.NoError(t, err)

		self.RootFields["process.working_directory"] = cwd
	}

	if qp.Proc.TtyMajor != 0 {
		self.RootFields["process.tty.char_device.major"] = qp.Proc.TtyMajor
		self.RootFields["process.tty.char_device.minor"] = qp.Proc.TtyMinor
	}

	return self
}

// makeEventOfCmd builds an mb.Event out of cmd and qev
func makeEventOfCmd(t *testing.T, cmd *exec.Cmd, qev quark.Event, be backend) mb.Event {
	// We should get at least FORK|EXEC|EXIT in the aggregation
	require.Equal(t,
		qev.Events&(quark.QUARK_EV_FORK|quark.QUARK_EV_EXEC|quark.QUARK_EV_EXIT),
		quark.QUARK_EV_FORK|quark.QUARK_EV_EXEC|quark.QUARK_EV_EXIT)
	// This is virtually impossible to fail, but we're pedantic
	require.True(t, qev.Process.Proc.Valid)

	qp := qev.Process

	interactive := tty.InteractiveFromTTY(tty.TTYDev{
		Major: qp.Proc.TtyMajor,
		Minor: qp.Proc.TtyMinor,
	})

	capEff, err := capabilities.FromPid(capabilities.Effective, os.Getpid())
	require.NoError(t, err)
	capPer, err := capabilities.FromPid(capabilities.Permitted, os.Getpid())
	require.NoError(t, err)

	cmdEvent := mb.Event{
		RootFields: mapstr.M{
			"event.type":                            []string{"start", "change", "end"},
			"event.action":                          "process_ran",
			"event.category":                        []string{"process"},
			"event.kind":                            "event",
			"process.name":                          "true",
			"process.args":                          []string{"/bin/true"},
			"process.args_count":                    1,
			"process.pid":                           uint32(cmd.Process.Pid),
			"process.executable":                    "/bin/true",
			"process.parent.pid":                    uint32(os.Getpid()),
			"process.start":                         time.Unix(0, int64(qp.Proc.TimeBoot)),
			"user.id":                               uint32(0),
			"user.group.id":                         uint32(0),
			"user.effective.id":                     uint32(0),
			"user.saved.id":                         qp.Proc.Suid,
			"user.saved.group.id":                   qp.Proc.Sgid,
			"user.name":                             "root",
			"user.group.name":                       "root",
			"process.interactive":                   interactive,
			"process.thread.capabilities.effective": capEff,
			"process.thread.capabilities.permitted": capPer,
		},
	}

	// Kprobe path depth is limited
	if be != Kprobe {
		cwd, err := os.Getwd()
		require.NoError(t, err)

		cmdEvent.RootFields["process.working_directory"] = cwd
	}

	return cmdEvent
}

// getConfigForQuark enables quark and allows hashing so we can test
// the cached hasher.
func getConfigForQuark(be backend) map[string]interface{} {
	config := map[string]interface{}{
		"module":   system.ModuleName,
		"datasets": []string{"process"},

		"process.backend": "kernel_tracing",
	}
	quarkForceKprobe = be == Kprobe

	return config
}
