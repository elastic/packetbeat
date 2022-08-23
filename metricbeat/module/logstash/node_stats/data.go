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

package node_stats

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/elastic/beats/v7/metricbeat/helper/elastic"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/elastic-agent-libs/version"

	"github.com/elastic/beats/v7/metricbeat/module/logstash"

	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

type jvm struct {
	GC  map[string]interface{} `json:"gc"`
	Mem struct {
		HeapMaxInBytes  int `json:"heap_max_in_bytes"`
		HeapUsedInBytes int `json:"heap_used_in_bytes"`
		HeapUsedPercent int `json:"heap_used_percent"`
	} `json:"mem"`
	UptimeInMillis int `json:"uptime_in_millis"`
}

type events struct {
	DurationInMillis int `json:"duration_in_millis"`
	In               int `json:"in"`
	Filtered         int `json:"filtered"`
	Out              int `json:"out"`
}

type commonStats struct {
	Events  events                 `json:"events"`
	JVM     jvm                    `json:"jvm"`
	Reloads map[string]interface{} `json:"reloads"`
	Queue   struct {
		EventsCount int `json:"events_count"`
	} `json:"queue"`
}

type cpu struct {
	Percent     int                    `json:"percent"`
	LoadAverage map[string]interface{} `json:"load_average,omitempty"`
}

type process struct {
	OpenFileDescriptors int `json:"open_file_descriptors"`
	MaxFileDescriptors  int `json:"max_file_descriptors"`
	CPU                 cpu `json:"cpu"`
}

type cgroup struct {
	CPUAcct map[string]interface{} `json:"cpuacct"`
	CPU     struct {
		Stat         map[string]interface{} `json:"stat"`
		ControlGroup string                 `json:"control_group"`
	} `json:"cpu"`
}

type os struct {
	CPU    cpu    `json:"cpu"`
	CGroup cgroup `json:"cgroup,omitempty"`
}

type pipeline struct {
	BatchSize int `json:"batch_size"`
	Workers   int `json:"workers"`
}

type nodeInfo struct {
	ID          string   `json:"id,omitempty"`
	UUID        string   `json:"uuid"`
	EphemeralID string   `json:"ephemeral_id"`
	Name        string   `json:"name"`
	Host        string   `json:"host"`
	Version     string   `json:"version"`
	Snapshot    bool     `json:"snapshot"`
	Status      string   `json:"status"`
	HTTPAddress string   `json:"http_address"`
	Pipeline    pipeline `json:"pipeline"`
}

// inNodeInfo represents the Logstash node info to be parsed from the Logstash API
// response. It contains nodeInfo (which is also used as-is elsewhere) + monitoring
// information.
type inNodeInfo struct {
	nodeInfo
	Monitoring struct {
		ClusterID string `json:"cluster_uuid"`
	} `json:"monitoring"`
}

type reloads struct {
	Successes int `json:"successes"`
	Failures  int `json:"failures"`
}

// NodeStats represents the stats of a Logstash node
type NodeStats struct {
	inNodeInfo
	commonStats
	Process   process                  `json:"process"`
	OS        os                       `json:"os"`
	Pipelines map[string]PipelineStats `json:"pipelines"`
}

// LogstashStats represents the logstash_stats sub-document indexed into .monitoring-logstash-*
type LogstashStats struct {
	commonStats
	Process   process         `json:"process"`
	OS        os              `json:"os"`
	Pipelines []PipelineStats `json:"pipelines"`
	Logstash  nodeInfo        `json:"logstash"`
	Timestamp common.Time     `json:"timestamp"`
}

// PipelineStats represents the stats of a Logstash pipeline
type PipelineStats struct {
	ID          string        `json:"id"`
	Hash        string        `json:"hash"`
	EphemeralID string        `json:"ephemeral_id"`
	Events      LongEnforce   `json:"events"`
	Reloads     reloads       `json:"reloads"`
	Queue       LongEnforce   `json:"queue"`
	Vertices    []LongEnforce `json:"vertices"`
}

// the default JSON encoder will attempt to unmarshall interface{}-typed
// digits as a float. None of these values are actually floats,
// and we get mapping errors when the JSON encoder on the other end
// will sometimes format floats over a certain size with scientific notation.
// Hard-defining these fields as a struct seems to be a bit touchy, as they're
// not documented in any one place on the logstash side, but they're still in the critical path for a number of
// O11y apps, dashboards, etc. So, force the JSON marshaller to treat digit values as floats; this is the most cautious solution.
type LongEnforce map[string]interface{}

func (le *LongEnforce) UnmarshalJSON(b []byte) error {
	base := make(map[string]interface{})
	if err := json.Unmarshal(b, &base); err != nil {
		return err
	}
	*le = make(LongEnforce)

	for key, item := range base {
		switch underlying := item.(type) {
		case float64:
			ival, fval := math.Modf(underlying)
			// based on the docs/code I've seen,
			// none of these are actual floats,
			// but check just to be sure
			if fval == 0 {
				(*le)[key] = int64(ival)
			} else {
				(*le)[key] = underlying
			}
		default:
			(*le)[key] = item

		}
	}
	return nil
}

func eventMapping(r mb.ReporterV2, content []byte, isXpack bool, logger *logp.Logger) error {
	var nodeStats NodeStats
	err := json.Unmarshal(content, &nodeStats)
	if err != nil {
		return errors.Wrap(err, "could not parse node stats response")
	}

	timestamp := common.Time(time.Now())

	// Massage Logstash node basic info
	nodeStats.nodeInfo.UUID = nodeStats.nodeInfo.ID
	nodeStats.nodeInfo.ID = ""

	proc := process{
		nodeStats.Process.OpenFileDescriptors,
		nodeStats.Process.MaxFileDescriptors,
		cpu{
			Percent: nodeStats.Process.CPU.Percent,
		},
	}

	o := os{
		cpu{
			LoadAverage: nodeStats.Process.CPU.LoadAverage,
		},
		nodeStats.OS.CGroup,
	}

	var pipelines []PipelineStats

	// The version from which we expect the node stats API to return a hash with pipeline documents. This is really just a formality so that
	// unit tests with much older fixture versions still work.
	var PipelineDocumentsContainHashVersion = version.MustNew("7.3.0")
	var StatsVersion = version.MustNew(nodeStats.Version)
	pipelineDocumentsShouldContainHash := !StatsVersion.LessThan(PipelineDocumentsContainHashVersion)

	for pipelineID, pipeline := range nodeStats.Pipelines {
		if pipelineDocumentsShouldContainHash && (pipeline.Hash == "" || pipeline.Vertices == nil) {
			logger.Warn(fmt.Sprintf("Pipeline document was discarded due to missing properties. This can happen when the Logstash node stats API is polled before the pipeline setup has completed. Pipeline ID: %s", pipelineID))
		} else {
			pipeline.ID = pipelineID
			pipelines = append(pipelines, pipeline)
		}
	}

	pipelines = getUserDefinedPipelines(pipelines)
	clusterToPipelinesMap := makeClusterToPipelinesMap(pipelines, nodeStats.Monitoring.ClusterID)

	for clusterUUID, clusterPipelines := range clusterToPipelinesMap {
		logstashStats := LogstashStats{
			nodeStats.commonStats,
			proc,
			o,
			clusterPipelines,
			nodeStats.nodeInfo,
			timestamp,
		}

		event := mb.Event{
			RootFields: mapstr.M{
				"service": mapstr.M{"name": logstash.ModuleName},
			},
			ModuleFields: mapstr.M{},
		}

		event.ModuleFields.Put("node.stats", logstashStats)
		event.RootFields.Put("service.id", nodeStats.ID)
		event.RootFields.Put("service.hostname", nodeStats.Host)
		event.RootFields.Put("service.version", nodeStats.Version)

		if clusterUUID != "" {
			event.ModuleFields.Put("cluster.id", clusterUUID)
			event.ModuleFields.Put("elasticsearch.cluster.id", clusterUUID)
		}

		// xpack.enabled in config using standalone metricbeat writes to `.monitoring` instead of `metricbeat-*`
		// When using Agent, the index name is overwritten anyways.
		if isXpack {
			index := elastic.MakeXPackMonitoringIndexName(elastic.Logstash)
			event.Index = index
		}

		r.Event(event)
	}

	return nil
}

func makeClusterToPipelinesMap(pipelines []PipelineStats, overrideClusterUUID string) map[string][]PipelineStats {
	var clusterToPipelinesMap map[string][]PipelineStats
	clusterToPipelinesMap = make(map[string][]PipelineStats)

	if overrideClusterUUID != "" {
		clusterToPipelinesMap[overrideClusterUUID] = pipelines
		return clusterToPipelinesMap
	}

	for _, pipeline := range pipelines {
		clusterUUIDs := common.StringSet{}
		for _, vertex := range pipeline.Vertices {
			clusterUUID := logstash.GetVertexClusterUUID(vertex, overrideClusterUUID)
			if clusterUUID != "" {
				clusterUUIDs.Add(clusterUUID)
			}
		}

		// If no cluster UUID was found in this pipeline, assign it a blank one
		if len(clusterUUIDs) == 0 {
			clusterUUIDs.Add("")
		}

		for clusterUUID := range clusterUUIDs {
			clusterPipelines := clusterToPipelinesMap[clusterUUID]
			if clusterPipelines == nil {
				clusterToPipelinesMap[clusterUUID] = []PipelineStats{}
			}

			clusterToPipelinesMap[clusterUUID] = append(clusterPipelines, pipeline)
		}
	}

	return clusterToPipelinesMap
}

func getUserDefinedPipelines(pipelines []PipelineStats) []PipelineStats {
	userDefinedPipelines := []PipelineStats{}
	for _, pipeline := range pipelines {
		if pipeline.ID[0] != '.' {
			userDefinedPipelines = append(userDefinedPipelines, pipeline)
		}
	}
	return userDefinedPipelines
}
