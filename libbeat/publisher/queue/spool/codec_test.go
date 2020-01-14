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

package spool

import (
	"testing"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	tests := map[string]codecID{
		"jsob":   codecJSON,
		"ubjson": codecUBJSON,
		"cborl":  codecCBORL,
	}

	event := publisher.Event{
		Content: beat.Event{
			Timestamp: time.Now().Round(0),
			Fields: common.MapStr{
				"time":       time.Now().UTC().Round(time.Millisecond),
				"commontime": common.Time(time.Now().UTC().Round(time.Millisecond)),
			},
		},
	}
	expected := publisher.Event{
		Content: beat.Event{
			Timestamp: event.Content.Timestamp,
			Fields: common.MapStr{
				"time":       event.Content.Fields["time"].(time.Time).Format(time.RFC3339Nano),
				"commontime": event.Content.Fields["commontime"].(common.Time).String(),
			},
		},
	}

	for name, codec := range tests {
		t.Run(name, func(t *testing.T) {
			encoder, err := newEncoder(codec)
			assert.NoError(t, err)

			encoded, err := encoder.encode(&event)
			assert.NoError(t, err)

			decoder := newDecoder()
			decoder.buf = encoded

			observed, err := decoder.Decode()
			assert.NoError(t, err)

			assert.Equal(t, expected, observed)
		})
	}
}
