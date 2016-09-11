package pluszone

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/module/nginx"
)

// Map body to []MapStr
func eventMapping(m *MetricSet, body io.ReadCloser, hostname string, metricset string) ([]common.MapStr, error) {
	// Nginx plus server zones:
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var zones map[string]interface{}
	if err := json.Unmarshal([]byte(b), &zones); err != nil {
		return nil, err
	}
	zones = nginx.Ftoi(zones)

	events := []common.MapStr{}

	for name, zone := range zones {
		event := common.MapStr{
			"hostname": hostname,
			"name":     name,
		}

		for k, v := range zone.(map[string]interface{}) {
			event[k] = v
		}

		events = append(events, event)
	}

	return events, nil
}
