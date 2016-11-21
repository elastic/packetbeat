// Package pluscache reads server cache status from Nginx, ngx_http_status_module is required.
package pluscache

import (
	"fmt"
	"net/http"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/module/nginx"
)

const (
	// defaultScheme is the default scheme to use when it is not specified in
	// the host config.
	defaultScheme = "http"

	// defaultPath is the default path to the ngx_http_status_module server cache endpoint on Nginx.
	defaultPath = "/status/caches"
)

var (
	debugf = logp.MakeDebug("nginx-status")
)

func init() {
	if err := mb.Registry.AddMetricSet("nginx", "pluscache", New); err != nil {
		panic(err)
	}
}

// MetricSet for fetching Nginx plus status.
type MetricSet struct {
	mb.BaseMetricSet

	client *http.Client // HTTP client that is reused across requests.
	url    string       // Nginx pluscache endpoint URL.

	requests int
}

// New creates new instance of MetricSet
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	// Additional configuration options
	config := struct {
		ServerStatusPath string `config:"server_status_path"`
	}{
		ServerStatusPath: defaultPath,
	}

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	u, err := nginx.GetURL(config.ServerStatusPath, base.Host())
	if err != nil {
		return nil, err
	}

	debugf("nginx-pluscache URL=%s", u)
	return &MetricSet{
		BaseMetricSet: base,
		url:           u.String(),
		client:        &http.Client{Timeout: base.Module().Config().Timeout},
		requests:      0,
	}, nil
}

// Fetch makes an HTTP request to fetch status metrics from the pluscache endpoint.
func (m *MetricSet) Fetch() ([]common.MapStr, error) {
	req, err := http.NewRequest("GET", m.url, nil)
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, resp.Status)
	}

	return eventMapping(m, resp.Body, m.Host(), m.Name())
}
