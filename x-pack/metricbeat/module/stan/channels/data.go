// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package channels

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
	"github.com/elastic/beats/metricbeat/mb"
)

var channelSchema = s.Schema{
	"cluster_id": c.Str("cluster_id"),
	"server_id":  c.Str("server_id"),
	"name":       c.Str("name"),
	"msgs":       c.Int("msgs"),
	"bytes":      c.Int("bytes"),
	"first_seq":  c.Int("first_seq"),
	"last_seq":   c.Int("last_seq"),
	"depth":      c.Int("depth", s.Optional), // aggregated by the module
}

// Subscription stores subscription related information
type Subscription struct {
	QueueName    string `json:"queue_name"`
	IsDurable    bool   `json:"is_durable"`
	IsOffline    bool   `json:"is_offline"`
	IsStalled    bool   `json:"is_stalled"`
	PendingCount int64  `json:"pending_count"`
	LastSent     int64  `json:"last_sent"`
}

// Channel stores channel related information
type Channel struct {
	Name          string         `json:"name"`
	Msgs          int64          `json:"msgs"`
	Bytes         int64          `json:"bytes"`
	FirstSeq      int64          `json:"first_seq"`
	LastSeq       int64          `json:"last_seq"`
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

// Channels stores channels related information
type Channels struct {
	ClusterID string    `json:"cluster_id"`
	ServerID  string    `json:"server_id"`
	Limit     uint64    `json:"limit"`
	Total     uint64    `json:"total"`
	Channels  []Channel `json:"channels,omitempty"`
}

// eventMapping map a channel to a Metricbeat event
func eventMapping(content map[string]interface{}) (mb.Event, error) {
	fields, err := channelSchema.Apply(content)
	if err != nil {
		return mb.Event{}, errors.Wrap(err, "failure applying channels schema")
	}

	event := mb.Event{
		MetricSetFields: fields,
		ModuleFields:    common.MapStr{},
	}
	return event, nil
}

// eventsMapping maps the top-level channel metrics AND also per-channel metrics AND subscriptions
func eventsMapping(content []byte, r mb.ReporterV2) error {
	channelsIn := Channels{}
	if err := json.Unmarshal(content, &channelsIn); err != nil {
		return errors.Wrap(err, "failure unmarshaling Nats streaming channels response to JSON")
	}

	for _, ch := range channelsIn.Channels {
		var evt mb.Event
		var err error
		var maxSubSeq int64
		for _, sub := range ch.Subscriptions {
			if sub.LastSent > maxSubSeq {
				maxSubSeq = sub.LastSent
			}
		}
		chWrapper := map[string]interface{}{
			"cluster_id": channelsIn.ClusterID,
			"server_id":  channelsIn.ServerID,
			"name":       ch.Name,
			"msgs":       ch.Msgs,
			"bytes":      ch.Bytes,
			"first_seq":  ch.FirstSeq,
			"last_seq":   ch.LastSeq,
			"depth":      ch.LastSeq - maxSubSeq, // queue depth is known channel seq number - maximum consumed by subscribers
		}

		if evt, err = eventMapping(chWrapper); err != nil {
			r.Error(errors.Wrap(err, "failure to map channel to its schema"))
		}
		if !r.Event(evt) {
			return nil
		}
	}

	return nil
}
