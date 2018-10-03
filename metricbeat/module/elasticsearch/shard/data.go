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

package shard

import (
	"encoding/json"

	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/module/elasticsearch"
)

var (
	schema = s.Schema{
		"state":   c.Str("state"),
		"primary": c.Bool("primary"),
		"index":   c.Str("index"),
		"shard":   c.Int("shard"),
	}
)

type stateStruct struct {
	ClusterName string `json:"cluster_name"`
	StateID     string `json:"state_uuid"`
	MasterNode  string `json:"master_node"`
	Nodes       map[string]struct {
		Name string `json:"name"`
	} `json:"nodes"`
	RoutingTable struct {
		Indices map[string]struct {
			Shards map[string][]map[string]interface{} `json:"shards"`
		} `json:"indices"`
	} `json:"routing_table"`
}

func eventsMapping(r mb.ReporterV2, content []byte) {
	stateData := &stateStruct{}
	err := json.Unmarshal(content, stateData)
	if err != nil {
		r.Error(err)
		return
	}

	for _, index := range stateData.RoutingTable.Indices {
		for _, shards := range index.Shards {
			for _, shard := range shards {
				event := mb.Event{}

				fields, err := schema.Apply(shard)
				if err != nil {
					r.Error(err)
					continue
				}

				// Handle node field: could be string or null
				err = elasticsearch.PassThruField("node", shard, fields)
				if err != nil {
					continue
				}

				// Handle relocating_node field: could be string or null
				err = elasticsearch.PassThruField("relocating_node", shard, fields)
				if err != nil {
					continue
				}

				event.ModuleFields = common.MapStr{}
				event.ModuleFields.Put("node.name", fields["node"])
				delete(fields, "node")

				event.ModuleFields.Put("index.name", fields["index"])
				delete(fields, "index")

				event.MetricSetFields = fields
				event.MetricSetFields.Put("number", fields["shard"])
				delete(event.MetricSetFields, "shard")

				delete(event.MetricSetFields, "relocating_node")
				event.MetricSetFields.Put("relocating_node.name", fields["relocating_node"])

				event.ModuleFields.Put("cluster.state.id", stateData.StateID)
				event.ModuleFields.Put("cluster.name", stateData.ClusterName)

				event.RootFields = common.MapStr{}
				event.RootFields.Put("service.name", "elasticsearch")

				r.Event(event)
			}
		}
	}
}
