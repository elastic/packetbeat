// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package synthexec

import (
	"fmt"
	"github.com/gofrs/uuid"
	"time"

	"github.com/elastic/beats/v7/heartbeat/eventext"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
)

type enricher func (event *beat.Event, se *SynthEvent) error

type streamEnricher struct {
	je *journeyEnricher
}

func (e *streamEnricher) enrich(event *beat.Event, se *SynthEvent) error {
	if e.je == nil || (se != nil && se.Type == "journey/start") {
		e.je = newJourneyEnricher()
	}

	return e.je.enrich(event, se)
}

// journeyEnricher holds state across received SynthEvents retaining fields
// where relevant to properly enrich *beat.Event instances.
type journeyEnricher struct {
	journeyComplete bool
	journey *Journey
	checkGroup      string
	errorCount      int
	lastError       error
	stepCount       int
	// The first URL we visit is the URL for this journey, which is set on the summary event.
	// We store the URL fields here for use on the summary event.
	urlFields common.MapStr
	start     time.Time
	end       time.Time
}

func newJourneyEnricher() *journeyEnricher {
	return &journeyEnricher{
		checkGroup: makeUuid(),
	}
}

func makeUuid() string {
	u, err := uuid.NewV1()
	if err != nil {
		panic("Cannot generate v1 UUID, this should never happen!")
	}
	return u.String()
}

func (je *journeyEnricher) enrich(event *beat.Event, se *SynthEvent) error {
	if se != nil && !se.Timestamp().IsZero() {
		event.Timestamp = se.Timestamp()
		// Record start and end so we can calculate journey duration accurately later
		switch se.Type {
		case "journey/start":
			je.lastError = nil
			je.checkGroup = makeUuid()
			je.journey = se.Journey
			je.start = event.Timestamp
		case "journey/end":
			je.end = event.Timestamp
		}
	} else {
		event.Timestamp = time.Now()
	}

	// No more synthEvents? In this case this is the summary event
	if se == nil {
		return nil
	}

	eventext.MergeEventFields(event, common.MapStr{
		"monitor": common.MapStr{
			"id": je.journey.Id,
			"name": je.journey.Name,
			"check_group": je.checkGroup,
		},
	})

	return je.enrichSynthEvent(event, se)
}

func (je *journeyEnricher) enrichSynthEvent(event *beat.Event, se *SynthEvent) error {
	switch se.Type {
	case "journey/end":
		je.journeyComplete = true
		return je.createSummary(event)
	case "step/end":
		je.stepCount++
	}

	eventext.MergeEventFields(event, se.ToMap())

	if je.urlFields == nil {
		if urlFields, err := event.GetValue("url"); err == nil {
			if ufMap, ok := urlFields.(common.MapStr); ok {
				je.urlFields = ufMap
			}
		}
	}

	var jobErr error
	if se.Error != nil {
		jobErr = stepError(se.Error)
		je.errorCount++
		je.lastError = jobErr
	}

	return jobErr
}

func (je *journeyEnricher) createSummary(event *beat.Event) error {
	var up, down int
	if je.errorCount > 0 {
		up = 0
		down = 1
	} else {
		up = 1
		down = 0
	}

	if je.journeyComplete {
		eventext.MergeEventFields(event, common.MapStr{
			"url": je.urlFields,
			"synthetics": common.MapStr{
				"type": "heartbeat/summary",
			},
			"monitor": common.MapStr{
				"duration": common.MapStr{
					"us": int64(je.end.Sub(je.start) / time.Microsecond),
				},
			},
			"summary": common.MapStr{
				"up": up,
				"down": down,
			},
		})
		return je.lastError
	}

	return fmt.Errorf("journey did not finish executing, %d steps ran", je.stepCount)
}

func stepError(e *SynthError) error {
	return fmt.Errorf("error executing step: %s", e.String())
}
