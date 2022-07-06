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

package wrappers

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/mitchellh/hashstructure"

	"github.com/elastic/beats/v7/heartbeat/ecserr"
	"github.com/elastic/beats/v7/heartbeat/eventext"
	"github.com/elastic/beats/v7/heartbeat/look"
	"github.com/elastic/beats/v7/heartbeat/monitors/jobs"
	"github.com/elastic/beats/v7/heartbeat/monitors/logger"
	"github.com/elastic/beats/v7/heartbeat/monitors/stdfields"
	"github.com/elastic/beats/v7/heartbeat/monitors/wrappers/monitorstate"
	"github.com/elastic/beats/v7/heartbeat/scheduler/schedule"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// WrapCommon applies the common wrappers that all monitor jobs get.
func WrapCommon(js []jobs.Job, stdMonFields stdfields.StdMonitorFields) []jobs.Job {
	mst := monitorstate.NewMonitorStateTracker()
	if stdMonFields.Type == "browser" {
		return WrapBrowser(js, stdMonFields)
	} else {
		return WrapLightweight(js, stdMonFields, mst)
	}
}

// WrapLightweight applies to http/tcp/icmp, everything but journeys involving node
func WrapLightweight(js []jobs.Job, stdMonFields stdfields.StdMonitorFields, mst *monitorstate.MonitorStateTracker) []jobs.Job {
	return jobs.WrapAllSeparately(
		jobs.WrapAll(
			js,
			addMonitorTimespan(stdMonFields),
			addServiceName(stdMonFields),
			addMonitorMeta(stdMonFields, len(js) > 1),
			addMonitorStatus(false),
			addMonitorDuration,
			addMonitorState(stdMonFields, mst),
		),
		func() jobs.JobWrapper {
			return makeAddSummary()
		})
}

// WrapBrowser is pretty minimal in terms of fields added. The browser monitor
// type handles most of the fields directly, since it runs multiple jobs in a single
// run it needs to take this task on in a unique way.
func WrapBrowser(js []jobs.Job, stdMonFields stdfields.StdMonitorFields) []jobs.Job {
	return jobs.WrapAll(
		js,
		addMonitorTimespan(stdMonFields),
		addServiceName(stdMonFields),
		addMonitorMeta(stdMonFields, false),
		addMonitorStatus(true),
		logJourneySummaries,
	)
}

// addMonitorMeta adds the id, name, and type fields to the monitor.
func addMonitorState(sf stdfields.StdMonitorFields, mst *monitorstate.MonitorStateTracker) jobs.JobWrapper {
	return func(job jobs.Job) jobs.Job {
		return func(event *beat.Event) ([]jobs.Job, error) {
			cont, err := job(event)

			// Only process terminal docs
			if len(cont) > 0 {
				return cont, err
			}

			trackerId := sf.ID
			if ip, err := event.GetValue("monitor.ip"); err != nil {
				trackerId = fmt.Sprintf("%s-%s", sf.ID, ip)
			}
			status, err := event.GetValue("monitor.status")

			ms, newMs, oldMs := mst.Compute(trackerId, status == "up")

			stateFields := common.MapStr{
				"id":     ms.Id(),
				"status": ms.Status(),
			}

			if newMs != nil {
				stateFields["starting"] = common.MapStr{
					"id":         newMs.Id(),
					"started_at": newMs.StartedAt.UnixMilli(),
					"status":     newMs.Status(),
				}
			}

			if oldMs != nil {
				logp.Warn("!!!ADD  END")
				endedAt := newMs.StartedAt.Add(-time.Millisecond)

				stateFields["ending"] = common.MapStr{
					"id":         oldMs.Id(),
					"started_at": oldMs.StartedAt.UnixMilli(),
					// Set the end time to 1ms before the new state
					"ended_at":    endedAt.UnixMilli(),
					"duration_ms": endedAt.Sub(oldMs.StartedAt).Milliseconds(),
					"checks":      oldMs.Checks,
					"up":          oldMs.Up,
					"down":        oldMs.Down,
					"status":      oldMs.Status(),
				}
			}
			logp.Warn("CHECKS: %s - u:%d d:%d", ms.Status(), ms.Up, ms.Down)

			eventext.MergeEventFields(event, common.MapStr{"state": stateFields})

			return cont, err
		}
	}
}

// addMonitorMeta adds the id, name, and type fields to the monitor.
func addMonitorMeta(sFields stdfields.StdMonitorFields, hashURLIntoID bool) jobs.JobWrapper {
	return func(job jobs.Job) jobs.Job {
		return func(event *beat.Event) ([]jobs.Job, error) {
			cont, err := job(event)

			id := sFields.ID
			name := sFields.Name
			// If multiple jobs are listed for this monitor, we can't have a single ID, so we hash the
			// unique URLs to create unique suffixes for the monitor.
			if hashURLIntoID {
				url, err := event.GetValue("url.full")
				if err != nil {
					logp.Error(fmt.Errorf("mandatory url.full key missing: %w", err))
					url = "n/a"
				}
				urlHash, _ := hashstructure.Hash(url, nil)
				id = fmt.Sprintf("%s-%x", sFields.ID, urlHash)
			}

			fields := mapstr.M{
				"type": sFields.Type,
			}

			if !sFields.IsLegacyBrowserSource || sFields.Type != "browser" {
				fields["id"] = id
				fields["name"] = name
			} else {
				// Id and name differs for local
				// - We use the monitor id and name for inline journeys ignoring the
				//   autogenerated `inline`journey id and name.
				// - Monitor id/name is concatenated with the journey id/name for
				// 	 project monitors
				journeyId, err := event.GetValue("monitor.id")
				if err == nil {
					fields["id"] = fmt.Sprintf("%s-%s", sFields.ID, journeyId)
				}
				journeyName, err := event.GetValue("monitor.name")
				if err == nil {
					fields["name"] = fmt.Sprintf("%s - %s", sFields.Name, journeyName)
				}
				fields["project"] = mapstr.M{
					"id":   sFields.ID,
					"name": sFields.Name,
				}
			}

			if sFields.Origin != "" {
				fields["origin"] = sFields.Origin
			}

			eventext.MergeEventFields(event, mapstr.M{"monitor": fields})
			return cont, err
		}
	}
}

func addMonitorTimespan(sf stdfields.StdMonitorFields) jobs.JobWrapper {
	return func(origJob jobs.Job) jobs.Job {
		return func(event *beat.Event) ([]jobs.Job, error) {
			cont, err := origJob(event)

			eventext.MergeEventFields(event, mapstr.M{
				"monitor": mapstr.M{
					"timespan": timespan(time.Now(), sf.Schedule, sf.Timeout),
				},
			})
			return cont, err
		}
	}
}

// Add service.name to monitors for APM interop
func addServiceName(sf stdfields.StdMonitorFields) jobs.JobWrapper {
	return func(origJob jobs.Job) jobs.Job {
		return func(event *beat.Event) ([]jobs.Job, error) {
			cont, err := origJob(event)

			if sf.Service.Name != "" {
				eventext.MergeEventFields(event, mapstr.M{
					"service": mapstr.M{
						"name": sf.Service.Name,
					},
				})
			}
			return cont, err
		}
	}
}

func timespan(started time.Time, sched *schedule.Schedule, timeout time.Duration) mapstr.M {
	maxEnd := sched.Next(started)

	if maxEnd.Sub(started) < timeout {
		maxEnd = started.Add(timeout)
	}

	return mapstr.M{
		"gte": started,
		"lt":  maxEnd,
	}
}

// addMonitorStatus wraps the given Job's execution such that any error returned
// by the original Job will be set as a field. The original error will not be
// passed through as a return value. Errors may still be present but only if there
// is an actual error wrapping the error.
func addMonitorStatus(summaryOnly bool) jobs.JobWrapper {
	return func(origJob jobs.Job) jobs.Job {
		return func(event *beat.Event) ([]jobs.Job, error) {
			cont, err := origJob(event)

			if summaryOnly {
				hasSummary, _ := event.Fields.HasKey("summary.up")
				if !hasSummary {
					return cont, nil
				}
			}

			fields := mapstr.M{
				"monitor": mapstr.M{
					"status": look.Status(err),
				},
			}
			if err != nil {
				var asECS *ecserr.ECSErr
				if errors.As(err, &asECS) {
					// Override the message of the error in the event it was wrapped
					asECS.Message = err.Error()
					fields["error"] = asECS
				} else {
					fields["error"] = look.Reason(err)
				}
			}
			eventext.MergeEventFields(event, fields)
			return cont, nil
		}
	}
}

// addMonitorDuration adds duration correctly for all non-browser jobs
func addMonitorDuration(job jobs.Job) jobs.Job {
	return func(event *beat.Event) ([]jobs.Job, error) {
		start := time.Now()
		cont, err := job(event)
		duration := time.Since(start)

		if event != nil {
			eventext.MergeEventFields(event, mapstr.M{
				"monitor": mapstr.M{
					"duration": look.RTT(duration),
				},
			})
			event.Timestamp = start

			logger.LogRun(event, nil)
		}

		return cont, err
	}
}

const META_STEP_COUNT = "__HEARTBEAT_STEP_COUNT__"

// logJourneySummaries emits a metric for the service when summary events are complete.
// Only applies to browser journeys.
func logJourneySummaries(job jobs.Job) jobs.Job {
	return func(event *beat.Event) ([]jobs.Job, error) {
		conts, err := job(event)

		summary, _ := event.GetValue("summary")
		if summary != nil {
			sc, _ := event.Meta.GetValue(META_STEP_COUNT)
			var scInt int
			// If we don't have it we have zero steps
			scInt, _ = sc.(int)
			logger.LogRun(event, &scInt)
		}

		return conts, err
	}
}

// makeAddSummary summarizes the job, adding the `summary` field to the last event emitted.
func makeAddSummary() jobs.JobWrapper {
	// This is a tricky method. The way this works is that we track the state across jobs in the
	// state struct here.
	state := struct {
		mtx        sync.Mutex
		monitorId  string
		remaining  uint16
		up         uint16
		down       uint16
		checkGroup string
		generation uint64
	}{
		mtx: sync.Mutex{},
	}
	// Note this is not threadsafe, must be called from a mutex
	resetState := func() {
		state.remaining = 1
		state.up = 0
		state.down = 0
		state.generation++
		u, err := uuid.NewV1()
		if err != nil {
			panic(fmt.Sprintf("cannot generate UUIDs on this system: %s", err))
		}
		state.checkGroup = u.String()
	}
	resetState()

	return func(job jobs.Job) jobs.Job {
		return func(event *beat.Event) ([]jobs.Job, error) {
			cont, jobErr := job(event)
			state.mtx.Lock()
			defer state.mtx.Unlock()

			// If the event is cancelled we don't record it as being either up or down since
			// we discard the event anyway.
			var eventStatus interface{}
			if !eventext.IsEventCancelled(event) {
				// After each job
				eventStatus, _ = event.GetValue("monitor.status")
				if eventStatus == "up" {
					state.up++
				} else {
					state.down++
				}
			}

			_, _ = event.PutValue("monitor.check_group", state.checkGroup)

			// Adjust the total remaining to account for new continuations
			state.remaining += uint16(len(cont))
			// Reduce total remaining to account for the just executed job
			state.remaining--

			// After last job
			if state.remaining == 0 {
				up := state.up
				down := state.down

				eventext.MergeEventFields(event, mapstr.M{
					"summary": mapstr.M{
						"up":   up,
						"down": down,
					},
				})
				resetState()
			}

			logp.Warn("RET CONT %s %d / %s", state.checkGroup, len(cont), jobErr)
			return cont, jobErr
		}
	}
}
