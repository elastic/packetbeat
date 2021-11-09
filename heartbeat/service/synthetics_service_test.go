package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/elastic/beats/v7/heartbeat/config"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func init() {
	Client = &MockClient{}

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return nil, errors.New(
			"error from web server",
		)
	}
}

func MockMonitorConfig(t *testing.T, rawConfigStr string) (config.Config, *common.Config )  {
	myJsonString := `{
	  "monitors": [{
		"type":     "test",
		"urls":     ["https://google.com"],
		"schedule": "@every 10m",
		"service_locations": ["us-east"]
	  }],
	  "service": {
		"username": "admin",
		"password": "changeme",
		"manifest_url": "http://localhost:8220"
 		}
	}`

	if rawConfigStr != "" {
		myJsonString = rawConfigStr
	}

	confMap := common.MapStr{}

	err := json.Unmarshal([]byte(myJsonString), &confMap)
	if err != nil {
		t.Error("Unable to parse config map,", err)
	}

	rawConfig, _ := common.NewConfigFrom(confMap)
	parsedConfig := config.DefaultConfig
	err1 := rawConfig.Unpack(&parsedConfig)

	if err1 != nil {
		t.Error(err1)
	}
	return parsedConfig, rawConfig
}

func MockSyntheticService(t *testing.T, rawConfigStr string) *SyntheticService {

	cfg, _ := MockMonitorConfig(t, rawConfigStr)

	return &SyntheticService{
		config: cfg,
	}

}

func TestPushConfiguration(t *testing.T) {

	myJsonString := `{
	  "monitors": [{
		"type":     "test",
		"urls":     ["https://google.com"],
		"schedule": "@every 10m",
		"service_locations": ["us-east"]
	  }],
	  "service": {
		"username": "admin",
		"password": "changeme",
		"manifest_url": "http://localhost:8220"
 		}
	}`

	sv := MockSyntheticService(t, myJsonString)

	payload := SyntheticServicePayload{}

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		bd, _ := ioutil.ReadAll(req.Body)
		err := json.Unmarshal(bd, &payload)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(
			"error from web server",
		)
	}
	username := "elastic"
	password := "changeme"

	sv.servicePushWait.Add(1)
	sv.pushConfigsToSyntheticsService("us-east", config.ServiceLocation{
		Url: "http://localhost:8220/cronjob",
	}, Output{
		Hosts:    []string{"http:localhost:9200"},
		Username: username,
		Password: password,
	})

	if len(payload.Monitors) != 1 {
		t.Error("expected payload monitors length to be 1")
	}

	assert.Equal(t, username, payload.Output.Username)
	assert.Equal(t, password, payload.Output.Password)
}

func TestRunViaSyntheticsService(t *testing.T) {

	myJsonString := `{
	  "monitors": [{
		"type":     "test",
		"urls":     ["https://google.com"],
		"schedule": "@every 10m",
		"service_locations": ["us_central"]
	  }],
	  "service": {
		"username": "admin",
		"password": "changeme",
		"manifest_url": "http://localhost:8220"
 		}
	}`

	sv := MockSyntheticService(t, myJsonString)

	binfo := beat.Info{
		Beat:        "heartbeat",
		IndexPrefix: "heartbeat",
		Version:     "8.0.0",
	}

	username := "elastic"
	password := "changeme"
	hosts := []string{"http://localhost:9200"}

	cfgMap := common.MapStr{
		"hosts":    hosts,
		"username": username,
		"password": password,
	}

	cfg, _ := common.NewConfigFrom(cfgMap)

	output := common.NewConfigNameSpace("output", cfg)

	bConfig := beat.BeatConfig{Output: *output}

	b := beat.Beat{
		Info:   binfo,
		Config: &bConfig,
	}

	payload := SyntheticServicePayload{}

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		if req.Body != nil {
			bd, _ := ioutil.ReadAll(req.Body)
			err := json.Unmarshal(bd, &payload)
			if err != nil {
				return nil, err
			}
			serviceUser, servicePass, _ := req.BasicAuth()

			assert.Equal(t, serviceUser, "admin")
			assert.Equal(t, servicePass, "changeme")

			return nil, errors.New(
				"error from web server",
			)
		}

		jsonStr := `{
					  "locations": {
						"us_central": {
						  "url": "us-central.synthetics.elastic.dev",
						  "geo": {
							"name": "US Central",
							"location": {"lat": 41.25, "lon": -95.86}
						  },
						  "status": "beta"
						}
					  }
					}`
		body := ioutil.NopCloser(bytes.NewReader([]byte(jsonStr)))

		return &http.Response{
			StatusCode: 200,
			Body:       body,
		}, nil
	}

	err2 := sv.Run(&b)
	if err2 != nil {
		t.Error(err2)
	}
	defer sv.servicePushWait.Wait()

	// wait for go routine
	time.Sleep(time.Second * 5)

	if len(payload.Monitors) != 1 {
		t.Error("expected payload monitors length to be 1")
	}

	assert.Equal(t, username, payload.Output.Username)
	assert.Equal(t, password, payload.Output.Password)
	assert.Equal(t, hosts, payload.Output.Hosts)

	sv.servicePushWait.Done()

	sv.Stop()
}

func TestValidateMonitorsSchedule(t *testing.T) {

	invalidMonitorCfg := `{
	  "monitors": [{
		"type":     "test",
		"urls":     ["https://google.com"],
		"schedule": "@every 10s",
		"service_locations": ["us-east"]
	  }]
	}`

	sv := MockSyntheticService(t, invalidMonitorCfg)

	err := sv.validateMonitorsSchedule()

	if err == nil {
		t.Error("it should return error of an invalid monitor")
	}

	validMonitorCfg := `{
	  "monitors": [{
		"type":     "test",
		"urls":     ["https://google.com"],
		"schedule": "@every 10m",
		"service_locations": ["us-east"]
	  }]
	}`

	sv = MockSyntheticService(t, validMonitorCfg)

	err = sv.validateMonitorsSchedule()

	if err != nil {
		t.Error("it should not return an error of a valid monitor")
	}

}
