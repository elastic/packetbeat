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

package varz

import (
	"encoding/json"
	"github.com/elastic/beats/libbeat/common"
	"time"
)

type Varz struct {
	ServerID       string `json:"server_id"`
	Version        string `json:"version"`
	Proto          int    `json:"proto"`
	GitCommit      string `json:"git_commit"`
	Go             string `json:"go"`
	Host           string `json:"host"`
	Addr           string `json:"addr"`
	MaxConnections int    `json:"max_connections"`
	PingInterval   int64  `json:"ping_interval"`
	PingMax        int    `json:"ping_max"`
	HTTPHost       string `json:"http_host"`
	HTTPPort       int    `json:"http_port"`
	HTTPSPort      int    `json:"https_port"`
	AuthTimeout    int    `json:"auth_timeout"`
	MaxControlLine int    `json:"max_control_line"`
	Cluster        struct {
		Addr        string `json:"addr"`
		ClusterPort int    `json:"cluster_port"`
		AuthTimeout int    `json:"auth_timeout"`
	} `json:"cluster"`
	TLSTimeout       float64   `json:"tls_timeout"`
	Port             int       `json:"port"`
	MaxPayload       int       `json:"max_payload"`
	Start            time.Time `json:"start"`
	Now              time.Time `json:"now"`
	Uptime           string    `json:"uptime"`
	Mem              int       `json:"mem"`
	Cores            int       `json:"cores"`
	CPU              int       `json:"cpu"`
	Connections      int       `json:"connections"`
	TotalConnections int       `json:"total_connections"`
	Routes           int       `json:"routes"`
	Remotes          int       `json:"remotes"`
	InMsgs           int       `json:"in_msgs"`
	OutMsgs          int       `json:"out_msgs"`
	InBytes          int       `json:"in_bytes"`
	OutBytes         int       `json:"out_bytes"`
	SlowConsumers    int       `json:"slow_consumers"`
	MaxPending       int       `json:"max_pending"`
	WriteDeadline    int       `json:"write_deadline"`
	Subscriptions    int       `json:"subscriptions"`
	HTTPReqStats     interface{} `json:"http_req_stats"`
	ConfigLoadTime time.Time `json:"config_load_time"`
}

func eventMapping(content []byte) common.MapStr {
	var data Varz
	json.Unmarshal(content, &data)
	// TODO: add error handling
	event := common.MapStr{
		"metrics": data,
	}
	return event
}
