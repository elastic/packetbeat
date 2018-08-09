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

package licenser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/outputs/elasticsearch"
)

func newServerClientPair(t *testing.T, handler http.HandlerFunc) (*httptest.Server, *elasticsearch.Client) {
	mux := http.NewServeMux()
	mux.Handle("/_xpack/", http.HandlerFunc(handler))

	server := httptest.NewServer(mux)

	client, err := elasticsearch.NewClient(elasticsearch.ClientSettings{URL: server.URL}, nil)
	if err != nil {
		t.Fatalf("could not create the elasticsearch client, error: %s", err)
	}

	return server, client
}

func TestParseJSON(t *testing.T) {
	t.Run("OSS release of Elasticsearch", func(t *testing.T) {
		h := func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Method Not Allowed", 405)
		}
		s, c := newServerClientPair(t, h)
		defer s.Close()
		defer c.Close()

		fetcher := NewElasticFetcher(c)
		oss, err := fetcher.Fetch()
		if assert.NoError(t, err) {
			return
		}

		assert.Equal(t, OSSLicense, oss)
	})

	t.Run("malformed JSON", func(t *testing.T) {
		h := func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello bad JSON"))
		}
		s, c := newServerClientPair(t, h)
		defer s.Close()
		defer c.Close()

		fetcher := NewElasticFetcher(c)
		_, err := fetcher.Fetch()
		assert.Error(t, err)
	})

	t.Run("401 response", func(t *testing.T) {
		h := func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", 401)
		}
		s, c := newServerClientPair(t, h)
		defer s.Close()
		defer c.Close()

		fetcher := NewElasticFetcher(c)
		_, err := fetcher.Fetch()
		assert.Equal(t, err.Error(), "Unauthorized access, could not connect to the xpack endpoint, verify your credentials")
	})

	t.Run("any error from the server", func(t *testing.T) {
		h := func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Not found", 404)
		}
		s, c := newServerClientPair(t, h)
		defer s.Close()
		defer c.Close()

		fetcher := NewElasticFetcher(c)
		_, err := fetcher.Fetch()
		assert.Error(t, err)
	})

	t.Run("200 response", func(t *testing.T) {
		filepath.Walk("data/", func(path string, i os.FileInfo, err error) error {
			if i.IsDir() {
				return nil
			}

			t.Run(path, func(t *testing.T) {
				h := func(w http.ResponseWriter, r *http.Request) {
					json, err := ioutil.ReadFile(path)
					if err != nil {
						t.Fatal("could not read JSON")
					}
					w.Write(json)
				}

				s, c := newServerClientPair(t, h)
				defer s.Close()
				defer c.Close()

				fetcher := NewElasticFetcher(c)
				license, err := fetcher.Fetch()
				if !assert.NoError(t, err) {
					return
				}

				id, _ := uuid.FromString("936183d8-f48c-4a3f-959a-a52aa2563279")
				assert.Equal(t, id, license.UUID)

				assert.NotNil(t, license.Type)
				assert.NotNil(t, license.Mode)
				assert.NotNil(t, license.Status)

				assert.False(t, license.Features.Graph.Available)
				assert.True(t, license.Features.Graph.Enabled)

				assert.False(t, license.Features.Logstash.Available)
				assert.True(t, license.Features.Logstash.Enabled)

				assert.False(t, license.Features.ML.Available)
				assert.True(t, license.Features.ML.Enabled)

				assert.True(t, license.Features.Monitoring.Available)
				assert.True(t, license.Features.Monitoring.Enabled)

				assert.True(t, license.Features.Rollup.Available)
				assert.True(t, license.Features.Rollup.Enabled)

				assert.False(t, license.Features.Security.Available)
				assert.True(t, license.Features.Security.Enabled)

				assert.False(t, license.Features.Watcher.Available)
				assert.True(t, license.Features.Watcher.Enabled)
			})

			return nil
		})
	})

	t.Run("parse milliseconds", func(t *testing.T) {
		t.Run("invalid", func(t *testing.T) {
			b := []byte("{ \"v\": \"\"}")
			ts := struct {
				V expiryTime `json:"v"`
			}{}

			err := json.Unmarshal(b, &ts)
			assert.Error(t, err)
		})

		t.Run("valid", func(t *testing.T) {
			b := []byte("{ \"v\": 1538060781728 }")
			ts := struct {
				V expiryTime `json:"v"`
			}{}

			err := json.Unmarshal(b, &ts)
			if !assert.NoError(t, err) {
				return
			}

			// 2018-09-27 15:06:21.728 +0000 UTC
			d := time.Date(2018, 9, 27, 15, 6, 21, 728000000, time.UTC).Sub((time.Time(ts.V)))
			assert.Equal(t, time.Duration(0), d)
		})
	})
}
