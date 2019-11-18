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

package scheduler

import (
	"errors"
	"sync"
	"time"

	"github.com/elastic/beats/heartbeat/hbregistry"

	"github.com/elastic/beats/heartbeat/scheduler/throttler"
	"github.com/elastic/beats/libbeat/common/atomic"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/monitoring"
)

const (
	statePreRunning int = iota + 1
	stateRunning
	stateDone
)

// countersRegistry is for tracking scheduler counters, confusingly known as 'stats' elsewhere in the stack.
var countersRegistry = hbregistry.StatsRegistry.NewRegistry("scheduler")

var jobsMissedDeadlineCounter = monitoring.NewUint(countersRegistry, "jobs.missed_deadline")

// gaugesRegistry is for tracking scheduler gauges, confusingly known as 'state' elsewhere in the stack.
var gaugesRegistry = hbregistry.StateRegistry.NewRegistry("scheduler")

var activeJobsGauge = monitoring.NewUint(gaugesRegistry, "jobs.active")
var activeTasksGauge = monitoring.NewUint(gaugesRegistry, "tasks.active")
var waitingTasksGauge = monitoring.NewUint(gaugesRegistry, "tasks.waiting")

type Scheduler struct {
	limit uint
	state atomic.Int
	done  chan int

	location *time.Location

	jobs               []*job
	activeJobs         *monitoring.Uint // gauge showing number of active jobs
	activeTasks        *monitoring.Uint // gauge showing number of active tasks
	waitingTasks       *monitoring.Uint // number of tasks waiting to run, but constrained by scheduler limit
	jobsPerSecond      *monitoring.Uint // rate of job processing computed over the past hour
	jobsMissedDeadline *monitoring.Uint // counter for number of jobs that missed start deadline

	throttler *throttler.Throttler
}

type Canceller func() error

// A job is a re-schedulable entry point in a set of tasks. Each task can return
// a new set of tasks being executed (subject to activeJobs task limits). Only after
// all tasks of a job have been finished, the job is marked as done and subject
// to be re-scheduled.
type job struct {
	id       string
	next     time.Time
	schedule Schedule
	fn       TaskFunc

	registered bool
	running    uint32 // count number of activeJobs task for job
}

// A single task in an activeJobs job.
type task struct {
	job *job
	fn  TaskFunc
}

// Single task in an activeJobs job. Optionally returns continuation of tasks to
// be executed within current job.
type TaskFunc func() []TaskFunc

type taskOverSignal struct {
	entry *job
	cont  []task // continuation tasks to be executed by concurrently for job at hand
}

// Schedule defines an interface for getting the next scheduled runtime for a job
type Schedule interface {
	// Next returns the next time a scheduled event occurs after the given time
	Next(now time.Time) (next time.Time)
	// Returns true if this schedule type should run once immediately before checking Next.
	// Cron tasks run at exact times so should set this to false.
	RunOnInit() bool
}

var debugf = logp.MakeDebug("scheduler")

// New creates a new Scheduler
func New(limit uint) *Scheduler {
	return NewWithLocation(limit, time.Local)
}

// NewWithLocation creates a new Scheduler using the given time zone.
func NewWithLocation(limit uint, location *time.Location) *Scheduler {
	stateInitial := statePreRunning
	return &Scheduler{
		limit:    limit,
		location: location,

		state: atomic.MakeInt(stateInitial),
		done:  make(chan int),

		activeJobs:         activeJobsGauge,
		activeTasks:        activeTasksGauge,
		waitingTasks:       waitingTasksGauge,
		jobsMissedDeadline: jobsMissedDeadlineCounter,

		throttler: throttler.NewThrottler(limit),
	}
}

// Start the scheduler.
func (s *Scheduler) Start() error {
	s.state.Store(stateRunning)
	s.throttler.Start()

	interval := time.Second * 10

	// Missed deadline reporter
	go func() {
		t := time.NewTicker(interval)

		// Counter used to check if we're missing more checks now than before
		missedAtLastCheck := uint64(0)
		for {
			select {
			case <-s.done:
				t.Stop()
				return
			case <-t.C:
				// Take a snapshot of items that have missed the deadline.
				// We only need to mutex long enough to do this. Since maps are reference types we can
				// do this through assignment

				missingNow := s.jobsMissedDeadline.Get()
				missedDelta := missingNow - missedAtLastCheck
				if missedDelta > 0 {
					logp.Warn("%d tasks have missed their schedule deadlines in the last %s.", missedDelta, interval)
				}
				missedAtLastCheck = missingNow
			}
		}
	}()

	return nil
}

// Stop all executing tasks in the scheduler. Cannot be restarted after Stop.
func (s *Scheduler) Stop() error {
	s.state.Store(stateDone)
	close(s.done)
	s.throttler.Stop()

	return nil
}

// ErrAlreadyStopped is returned when an Add operation is attempted after the scheduler
// has already stopped.
var ErrAlreadyStopped = errors.New("attempted to add job to already stopped scheduler")

// Add adds the given TaskFunc to the current scheduler. Will return an error if the scheduler
// is done.
func (s *Scheduler) Add(sched Schedule, id string, entrypoint TaskFunc) (removeFn func() error, err error) {
	if s.state.Load() == stateDone {
		return nil, ErrAlreadyStopped
	}

	removed := atomic.MakeBool(false)
	removeCh := make(chan bool)
	removeFn = func() error {
		// Safely close the channel exactly once
		// Clients may invoke this function multiple times
		if removed.CAS(false, true) {
			close(removeCh)
		}
		return nil
	}

	var timer *time.Timer
	go func() {
		// lastRanAt stores the last time the task was invoked
		// The initial value is time.Now() because we use it to get the next time a job is scheduled to run
		lastRanAt := time.Now()
		// If this job should be run immediately set the timestamp to the epoch.
		// That will cause any (plausible) schedule to run immediately
		if sched.RunOnInit() {
			lastRanAt = time.Unix(0, 0)
		}

		// true for the  first iteration
		for i := uint64(0); true; i++ {
			now := time.Now()
			// We use the time the last task was invoked to figure out when to next run it.
			next := sched.Next(lastRanAt)

			// If we are running behind schedule there's no need to invoke the timer, we can run right away.
			// This can happen if the task is slow, and also on first run when we intentionally set the lastRan
			// to the epoch.
			if next.Before(now) {
				lastRanAt = now
				s.runOnce(id, entrypoint)

				// Record the missed deadline except in the case of the first run, where this would otherwise
				// always trigger
				if i > 0 {
					s.jobsMissedDeadline.Inc()
				}
			}

			if timer == nil {
				timer = time.NewTimer(next.Sub(now))
			} else {
				timer.Reset(next.Sub(now))
			}

			select {
			case <-timer.C:
				// time may have elapsed between when we scheduled the task and when it was actually invoked
				// it may seem to make more sense to just use `now` rather than `time.Now`, however, there's an
				// advantage to being more precise here. In the case where more jobs are scheduled than can
				// be executed simultaneously, and schedules like `@every 5s` are in use this will cause any delayed
				// job to be permanently offset. This will naturally lead to a more even distribution of jobs over
				// the timeline in short order, rather than repeatedly bursting schedules. This should lead to more
				// reliability in those high concurrency scenarios.
				//
				// For cron scheduling this does nothing since cron schedules run at exact times.
				lastRanAt = time.Now()
				s.runOnce(id, entrypoint)
			case <-removeCh:
				return
			}
		}
	}()

	return removeFn, nil
}

// runOnce runs a TaskFunc and its continuations once.
func (s *Scheduler) runOnce(id string, entrypoint TaskFunc) {
	// Since we run all continuations asynchronously we use a wait group to block until we're done.
	wg := sync.WaitGroup{}

	// Since task funcs can emit continuations recursively we need a function to execute
	// recursively. We declare the function variable before definition to allow for recursion.
	var runRecursive func(task TaskFunc)
	runRecursive = func(task TaskFunc) {
		wg.Add(1)
		defer wg.Done()

		// The accounting for waitingTasks/activeJobs is done using atomics. Absolute accuracy is not critical here so the gap
		// between modifying waitingTasks and activeJobs is acceptable.
		s.waitingTasks.Inc()

		// Acquire an execution slot in keeping with heartbeat.scheduler.limit
		acquired, releaseSlot := s.throttler.AcquireSlot()
		defer releaseSlot()

		s.activeTasks.Inc()
		s.waitingTasks.Dec()

		// The only situation in which we can't acquire a slot is during shutdown. In that case we can just return
		// without worrying about any of the counters since we're going away soon.
		if !acquired {
			return
		}
		continuations := task()
		s.activeTasks.Dec()

		for _, cont := range continuations {
			go runRecursive(cont) // Run continuations in parallel, note that these each will acquire their own slots
		}
	}

	s.activeJobs.Inc()
	runRecursive(entrypoint)

	wg.Wait()
	s.activeJobs.Dec()
}
