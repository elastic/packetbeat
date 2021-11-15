package process

import (
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/opt"
)

type RunState byte

const (
	// RunStateSleep corresponds to a sleep state
	RunStateSleep = 'S'
	// RunStateRun corresponds to a running state
	RunStateRun = 'R'
	// RunStateStop corresponds to a stopped state
	RunStateStop = 'T'
	// RunStateZombie marks a zombie process
	RunStateZombie = 'Z'
	// RunStateIdle corresponds to an idle state
	RunStateIdle = 'D'
	// RunStateUnknown corresponds to a process in an unknown state
	RunStateUnknown = '?'
)

type ProcState struct {
	// Basic Process data
	Name     string  `struct:"name,omitempty"`
	State    string  `struct:"state,omitempty"`
	Username string  `struct:"username,omitempty"`
	pid      opt.Int `struct:"pid,omitempty"`
	Ppid     opt.Int `struct:"ppid,omitempty"`
	Pgid     opt.Int `struct:"pgid,omitempty"`

	// Extended Process Data
	Args    []string      `struct:"args,omitempty"`
	Cmdline opt.String    `struct:"cmdline,omitempty"`
	Cwd     opt.String    `struct:"cwd,omitempty"`
	Exe     opt.String    `struct:"exe,omitempty"`
	Env     common.MapStr `struct:"env,omitempty"`

	// Resource Metrics
	Memory ProcMeminfo `struct:"memory,omitempty"`
	CPU    ProcCPUInfo `struct:"cpu,omitempty"`
}

type ProcCPUInfo struct {
	StartTime common.Time `struct:"start_time,omitempty"`
}

type CPUTotal struct {
	Value opt.Float  `struct:"value,omitempty"`
	Pct   opt.Float  `struct:"pct,omitempty"`
	Norm  opt.PctOpt `struct:"norm,omitempty"`
}

type ProcMeminfo struct {
	Size  opt.Uint   `struct:"size,omitempty"`
	Share opt.Uint   `struct:"share,omitempty"`
	Rss   MemBytePct `struct:"rss,omitempty"`
}

type MemBytePct struct {
	Bytes opt.Uint  `struct:"bytes,omitempty"`
	Pct   opt.Float `struct:"pct,omitempty"`
}
