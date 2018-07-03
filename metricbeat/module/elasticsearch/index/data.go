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

package index

import (
	"encoding/json"

	"github.com/joeshaw/multierror"

	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/module/elasticsearch"
)

var (
	schema = s.Schema{
		"total": c.Dict("total", s.Schema{
			"docs": c.Dict("docs", s.Schema{
				"count":   c.Int("count"),
				"deleted": c.Int("deleted"),
			}),
			"store": c.Dict("store", s.Schema{
				"size": s.Object{
					"bytes": c.Int("size_in_bytes"),
				},
			}),
			"segments": c.Dict("segments", s.Schema{
				"count": c.Int("count"),
				"memory": s.Object{
					"bytes": c.Int("memory_in_bytes"),
				},
			}),
		}),
	}
)

func eventsMapping(r mb.ReporterV2, info elasticsearch.Info, content []byte) error {
	var indicesStruct struct {
		Indices map[string]map[string]interface{} `json:"indices"`
	}

	err := json.Unmarshal(content, &indicesStruct)
	if err != nil {
		r.Error(err)
		return err
	}

	var errors multierror.Errors
	for name, index := range indicesStruct.Indices {
		event := mb.Event{}
		event.MetricSetFields, err = schema.Apply(index)
		if err != nil {
			r.Error(err)
			errors = append(errors, err)
		}
		// Write name here as full name only available as key
		event.MetricSetFields["name"] = name
		event.RootFields = common.MapStr{}
		event.RootFields.Put("service.name", "elasticsearch")
		event.ModuleFields = common.MapStr{}
		event.ModuleFields.Put("cluster.name", info.ClusterName)
		event.ModuleFields.Put("cluster.id", info.ClusterID)
		r.Event(event)
	}

	return errors.Err()
}
