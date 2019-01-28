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

package status

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
	"github.com/elastic/beats/metricbeat/helper/elastic"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/module/kibana"
)

var (
	schema = s.Schema{
		"uuid": c.Str("uuid"),
		"name": c.Str("name"),
		"version": c.Dict("version", s.Schema{
			"number": c.Str("number"),
		}),
		"status": c.Dict("status", s.Schema{
			"overall": c.Dict("overall", s.Schema{
				"state": c.Str("state"),
			}),
		}),
		"metrics": c.Dict("metrics", s.Schema{
			"requests": c.Dict("requests", s.Schema{
				"total":       c.Int("total"),
				"disconnects": c.Int("disconnects"),
			}),
			"concurrent_connections": c.Int("concurrent_connections"),
		}),
	}
)

func eventMapping(r mb.ReporterV2, content []byte) error {
	var event mb.Event
	event.RootFields = common.MapStr{}
	event.RootFields.Put("service.name", kibana.ModuleName)

	var data map[string]interface{}
	err := json.Unmarshal(content, &data)
	if err != nil {
		event.Error = errors.Wrap(err, "failure parsing Kibana Status API response")
		r.Event(event)
		return event.Error
	}

	dataFields, err := schema.Apply(data)
	if err != nil {
		event.Error = errors.Wrap(err, "failure to apply status schema")
		r.Event(event)
		return event.Error
	}

	// Set service ID
	uuid, err := dataFields.GetValue("uuid")
	if err != nil {
		event.Error = elastic.MakeErrorForMissingField("kibana.uuid", elastic.Kibana)
		r.Event(event)
		return event.Error
	}
	event.RootFields.Put("service.id", uuid)
	dataFields.Delete("uuid")

	event.MetricSetFields = dataFields

	r.Event(event)
	return nil
}
