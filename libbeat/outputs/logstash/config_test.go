package logstash

import (
	"testing"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	info := beat.Info{Beat: "testbeat", Name: "foo", IndexPrefix: "bar"}
	for name, test := range map[string]struct {
		config         *common.Config
		expectedConfig *Config
		err            bool
	}{
		"default config": {
			config: common.MustNewConfigFrom([]byte(`{ }`)),
			expectedConfig: &Config{
				LoadBalance:      false,
				Pipelining:       2,
				BulkMaxSize:      2048,
				SlowStart:        false,
				CompressionLevel: 3,
				Timeout:          30 * time.Second,
				MaxRetries:       3,
				TTL:              0 * time.Second,
				Backoff: Backoff{
					Init: 1 * time.Second,
					Max:  60 * time.Second,
				},
				EscapeHTML: false,
				Index:      "bar",
			},
		},
		"config given": {
			config: common.MustNewConfigFrom(common.MapStr{
				"index":         "beat-index",
				"loadbalance":   true,
				"bulk_max_size": 1024,
				"slow_start":    false,
			}),
			expectedConfig: &Config{
				LoadBalance:      true,
				BulkMaxSize:      1024,
				Pipelining:       2,
				SlowStart:        false,
				CompressionLevel: 3,
				Timeout:          30 * time.Second,
				MaxRetries:       3,
				TTL:              0 * time.Second,
				Backoff: Backoff{
					Init: 1 * time.Second,
					Max:  60 * time.Second,
				},
				EscapeHTML: false,
				Index:      "beat-index",
			},
		},
		"removed config setting": {
			config: common.MustNewConfigFrom(common.MapStr{
				"port": "8080",
			}),
			expectedConfig: nil,
			err:            true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			cfg, err := readConfig(test.config, info)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, cfg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedConfig, cfg)
			}
		})
	}
}
