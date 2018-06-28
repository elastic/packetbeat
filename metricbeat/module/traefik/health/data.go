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

package health

import (
	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
)

var (
	schema = s.Schema{
		"uptime": s.Object{
			"sec": c.Int("uptime_sec"),
		},
		"response": s.Object{
			"count": c.Int("total_count"),
			"avg_time": s.Object{
				"us": c.Int("average_response_time_us"),
			},
			"status_codes": s.Object{},
		},
	}
)

func eventMapping(health map[string]interface{}) (common.MapStr, *s.Errors) {
	if averageResponseTimeSec, ok := health["average_response_time_sec"]; ok {
		health["average_response_time_us"] = averageResponseTimeSec.(float64) * 1000 * 1000
	}

	event, _ := schema.Apply(health)

	statusCodeCountMap := health["total_status_code_count"].(map[string]interface{})
	for code, count := range statusCodeCountMap {
		event.Put("response.status_codes."+code, count)
	}

	return event, nil
}
