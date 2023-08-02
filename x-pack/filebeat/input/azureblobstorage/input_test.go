// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package azureblobstorage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"

	v2 "github.com/elastic/beats/v7/filebeat/input/v2"
	cursor "github.com/elastic/beats/v7/filebeat/input/v2/input-cursor"
	"github.com/elastic/beats/v7/libbeat/beat"
	beattest "github.com/elastic/beats/v7/libbeat/publisher/testing"
	"github.com/elastic/beats/v7/x-pack/filebeat/input/azureblobstorage/mock"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/logp"
)

const (
	beatsContainer              = "beatscontainer"
	beatsContainer2             = "beatscontainer2"
	beatsMultilineJSONContainer = "beatsmultilinejsoncontainer"
	beatsJSONContainer          = "beatsjsoncontainer"
	beatsNdJSONContainer        = "beatsndjsoncontainer"
	beatsGzJSONContainer        = "beatsgzjsoncontainer"
)

func Test_StorageClient(t *testing.T) {
	tests := []struct {
		name          string
		baseConfig    map[string]interface{}
		mockHandler   func() http.Handler
		expected      map[string]bool
		expectedError error
	}{
		{
			name: "SingleContainerWithPoll_NoErr",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": beatsContainer,
					},
				},
			},
			mockHandler: mock.AzureStorageServer,
			expected: map[string]bool{
				mock.Beatscontainer_blob_ata_json:      true,
				mock.Beatscontainer_blob_data3_json:    true,
				mock.Beatscontainer_blob_docs_ata_json: true,
			},
		},
		{
			name: "SingleContainerWithoutPoll_NoErr",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                false,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": beatsContainer,
					},
				},
			},
			mockHandler: mock.AzureStorageServer,
			expected: map[string]bool{
				mock.Beatscontainer_blob_ata_json:      true,
				mock.Beatscontainer_blob_data3_json:    true,
				mock.Beatscontainer_blob_docs_ata_json: true,
			},
		},
		{
			name: "TwoContainersWithPoll_NoErr",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": beatsContainer,
					},
					{
						"name": beatsContainer2,
					},
				},
			},
			mockHandler: mock.AzureStorageServer,
			expected: map[string]bool{
				mock.Beatscontainer_blob_ata_json:      true,
				mock.Beatscontainer_blob_data3_json:    true,
				mock.Beatscontainer_blob_docs_ata_json: true,
				mock.Beatscontainer_2_blob_ata_json:    true,
				mock.Beatscontainer_2_blob_data3_json:  true,
			},
		},
		{
			name: "TwoContainersWithoutPoll_NoErr",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                false,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": beatsContainer,
					},
					{
						"name": beatsContainer2,
					},
				},
			},
			mockHandler: mock.AzureStorageServer,
			expected: map[string]bool{
				mock.Beatscontainer_blob_ata_json:      true,
				mock.Beatscontainer_blob_data3_json:    true,
				mock.Beatscontainer_blob_docs_ata_json: true,
				mock.Beatscontainer_2_blob_ata_json:    true,
				mock.Beatscontainer_2_blob_data3_json:  true,
			},
		},
		{
			name: "SingleContainerPoll_InvalidContainerErr",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": "azuretest",
					},
				},
			},
			mockHandler:   mock.AzureStorageServer,
			expected:      map[string]bool{},
			expectedError: mock.NotFoundErr,
		},
		{
			name: "SingleContainerWithoutPoll_InvalidBucketErr",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                false,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": "azuretest",
					},
				},
			},
			mockHandler:   mock.AzureStorageServer,
			expected:      map[string]bool{},
			expectedError: mock.NotFoundErr,
		},
		{
			name: "TwoContainersWithPoll_InvalidBucketErr",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": "azurenew",
					},
					{
						"name": "azurelatest",
					},
				},
			},
			mockHandler:   mock.AzureStorageServer,
			expected:      map[string]bool{},
			expectedError: mock.NotFoundErr,
		},
		{
			name: "SingleBucketWithPoll_InvalidConfigValue",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         5100,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": beatsContainer,
					},
				},
			},
			mockHandler:   mock.AzureStorageServer,
			expected:      map[string]bool{},
			expectedError: errors.New("requires value <= 5000 accessing 'max_workers'"),
		},
		{
			name: "TwoBucketWithPoll_InvalidConfigValue",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         5100,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": beatsContainer,
					},
					{
						"name": beatsContainer2,
					},
				},
			},
			mockHandler:   mock.AzureStorageServer,
			expected:      map[string]bool{},
			expectedError: errors.New("requires value <= 5000 accessing 'max_workers'"),
		},
		{
			name: "ReadJSON",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": beatsJSONContainer,
					},
				},
			},
			mockHandler: mock.AzureStorageFileServer,
			expected: map[string]bool{
				mock.BeatsFilesContainer_log_json[0]: true,
				mock.BeatsFilesContainer_log_json[1]: true,
				mock.BeatsFilesContainer_log_json[2]: true,
			},
		},
		{
			name: "ReadOctetStreamJSON",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": beatsMultilineJSONContainer,
					},
				},
			},
			mockHandler: mock.AzureStorageFileServer,
			expected: map[string]bool{
				mock.BeatsFilesContainer_multiline_json[0]: true,
				mock.BeatsFilesContainer_multiline_json[1]: true,
			},
		},
		{
			name: "ReadNdJSON",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": beatsNdJSONContainer,
					},
				},
			},
			mockHandler: mock.AzureStorageFileServer,
			expected: map[string]bool{
				mock.BeatsFilesContainer_log_ndjson[0]: true,
				mock.BeatsFilesContainer_log_ndjson[1]: true,
			},
		},
		{
			name: "ReadMultilineGzJSON",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": beatsGzJSONContainer,
					},
				},
			},
			mockHandler: mock.AzureStorageFileServer,
			expected: map[string]bool{
				mock.BeatsFilesContainer_multiline_json_gz[0]: true,
				mock.BeatsFilesContainer_multiline_json_gz[1]: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serv := httptest.NewServer(tt.mockHandler())
			t.Cleanup(serv.Close)

			cfg := conf.MustNewConfigFrom(tt.baseConfig)
			conf := config{}
			err := cfg.Unpack(&conf)
			if err != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
				return
			}
			input := newStatelessInput(conf, serv.URL+"/")

			assert.Equal(t, "azure-blob-storage-stateless", input.Name())
			assert.NoError(t, input.Test(v2.TestContext{}))

			chanClient := beattest.NewChanClient(len(tt.expected))
			t.Cleanup(func() { _ = chanClient.Close() })

			ctx, cancel := newV2Context()
			t.Cleanup(cancel)
			ctx.ID += tt.name

			var g errgroup.Group
			g.Go(func() error {
				return input.Run(ctx, chanClient)
			})

			var timeout *time.Timer
			if conf.PollInterval != nil {
				timeout = time.NewTimer(1*time.Second + *conf.PollInterval)
			} else {
				timeout = time.NewTimer(5 * time.Second)
			}
			t.Cleanup(func() { timeout.Stop() })

			if len(tt.expected) == 0 {
				if tt.expectedError != nil && g.Wait() != nil {
					//nolint:errorlint // This will never be a wrapped error
					if tt.expectedError == mock.NotFoundErr {
						arr := strings.Split(g.Wait().Error(), "\n")
						errStr := strings.Join(arr[1:], "\n")
						assert.Equal(t, tt.expectedError.Error(), errStr)
					} else {
						assert.EqualError(t, g.Wait(), tt.expectedError.Error())
					}
					cancel()
				} else {
					assert.NoError(t, g.Wait())
					cancel()
				}
				return
			}

			var receivedCount int
		wait:
			for {
				select {
				case <-timeout.C:
					t.Errorf("timed out waiting for %d events", len(tt.expected))
					cancel()
					return
				case got := <-chanClient.Channel:
					var val interface{}
					var err error
					val, err = got.Fields.GetValue("message")
					assert.NoError(t, err)
					assert.True(t, tt.expected[val.(string)])
					assert.Equal(t, tt.expectedError, err)
					receivedCount += 1
					if receivedCount == len(tt.expected) {
						cancel()
						break wait
					}
				}
			}
		})
	}
}

func Test_Concurrency(t *testing.T) {
	tests := []struct {
		name        string
		baseConfig  map[string]interface{}
		mockHandler func() http.Handler
		expectedLen int
	}{
		{
			name: "TestConcurrency_100_Workers",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         100,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": mock.ConcurrencyContainer,
					},
				},
			},
			mockHandler: mock.AzureConcurrencyServer,
			expectedLen: mock.TotalRandomDataSets,
		},
		{
			name: "TestConcurrency_1000_Workers",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         1000,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": mock.ConcurrencyContainer,
					},
				},
			},
			mockHandler: mock.AzureConcurrencyServer,
			expectedLen: mock.TotalRandomDataSets,
		},
		{
			name: "TestConcurrency_2000_Workers",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         2000,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": mock.ConcurrencyContainer,
					},
				},
			},
			mockHandler: mock.AzureConcurrencyServer,
			expectedLen: mock.TotalRandomDataSets,
		},
		{
			name: "TestConcurrency_3000_Workers",
			baseConfig: map[string]interface{}{
				"account_name":                        "beatsblobnew",
				"auth.shared_credentials.account_key": "7pfLm1betGiRyyABEM/RFrLYlafLZHbLtGhB52LkWVeBxE7la9mIvk6YYAbQKYE/f0GdhiaOZeV8+AStsAdr/Q==",
				"max_workers":                         3000,
				"poll":                                true,
				"poll_interval":                       "10s",
				"containers": []map[string]interface{}{
					{
						"name": mock.ConcurrencyContainer,
					},
				},
			},
			mockHandler: mock.AzureConcurrencyServer,
			expectedLen: mock.TotalRandomDataSets,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serv := httptest.NewServer(tt.mockHandler())
			t.Cleanup(serv.Close)

			cfg := conf.MustNewConfigFrom(tt.baseConfig)
			conf := config{}
			err := cfg.Unpack(&conf)
			assert.NoError(t, err)
			input := azurebsInput{
				config:     conf,
				serviceURL: serv.URL + "/",
			}
			name := input.Name()
			if name != "azure-blob-storage" {
				t.Errorf(`unexpected input name: got:%q want:"azure-blob-storage"`, name)
			}

			var src cursor.Source
			// This test will always have only one container
			for _, c := range input.config.Containers {
				container := tryOverrideOrDefault(input.config, c)
				src = &Source{
					AccountName:   input.config.AccountName,
					ContainerName: c.Name,
					MaxWorkers:    *container.MaxWorkers,
					Poll:          *container.Poll,
					PollInterval:  *container.PollInterval,
				}
			}
			v2Ctx, cancel := newV2Context()
			t.Cleanup(cancel)
			v2Ctx.ID += tt.name
			var client publisher
			st := newState()

			var g errgroup.Group
			g.Go(func() error {
				return input.run(v2Ctx, src, st, &client)
			})
			timeout := time.NewTimer(100 * time.Second)
			t.Cleanup(func() { timeout.Stop() })
			client.Channel = make(chan beat.Event)
			var receivedCount int
		wait:
			for {
				select {
				case <-timeout.C:
					t.Errorf("timed out waiting for %d events", tt.expectedLen)
					cancel()
					return
				case got := <-client.Channel:
					_, err = got.Fields.GetValue("message")
					assert.NoError(t, err)
					receivedCount += 1
					if receivedCount >= tt.expectedLen {
						cancel()
						break wait
					}
				}
			}
		})
	}
}

type publisher struct {
	Channel chan beat.Event
	mu      sync.Mutex
	cursors []map[string]interface{}
}

func (p *publisher) Publish(e beat.Event, cursor interface{}) error {
	p.mu.Lock()
	p.Channel <- e
	if cursor != nil {
		var c map[string]interface{}
		chkpt, ok := cursor.(*Checkpoint)
		if !ok {
			return fmt.Errorf("invalid cursor type for testing: %T", cursor)
		}
		cursorBytes, err := json.Marshal(chkpt)
		if err != nil {
			return fmt.Errorf("error marshaling cursor data: %w", err)
		}
		err = json.Unmarshal(cursorBytes, &c)
		if err != nil {
			return fmt.Errorf("error converting checkpoint struct to cursor map: %w", err)
		}

		p.cursors = append(p.cursors, c)
	}
	p.mu.Unlock()
	return nil
}

func newV2Context() (v2.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	return v2.Context{
		Logger:      logp.NewLogger("azure-blob-storage_test"),
		ID:          "test_id:",
		Cancelation: ctx,
	}, cancel
}
