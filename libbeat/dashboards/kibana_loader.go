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

package dashboards

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"time"

	"github.com/joeshaw/multierror"
	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/kibana"
	"github.com/elastic/beats/v7/libbeat/logp"
)

var importAPI = "/api/saved_objects/_import"

type KibanaLoader struct {
	client       *kibana.Client
	config       *Config
	version      common.Version
	hostname     string
	msgOutputter MessageOutputter
}

// NewKibanaLoader creates a new loader to load Kibana files
func NewKibanaLoader(ctx context.Context, cfg *common.Config, dashboardsConfig *Config, hostname string, msgOutputter MessageOutputter) (*KibanaLoader, error) {

	if cfg == nil || !cfg.Enabled() {
		return nil, fmt.Errorf("Kibana is not configured or enabled")
	}

	client, err := getKibanaClient(ctx, cfg, dashboardsConfig.Retry, 0)
	if err != nil {
		return nil, fmt.Errorf("Error creating Kibana client: %v", err)
	}

	loader := KibanaLoader{
		client:       client,
		config:       dashboardsConfig,
		version:      client.GetVersion(),
		hostname:     hostname,
		msgOutputter: msgOutputter,
	}

	version := client.GetVersion()
	loader.statusMsg("Initialize the Kibana %s loader", version.String())

	return &loader, nil
}

func getKibanaClient(ctx context.Context, cfg *common.Config, retryCfg *Retry, retryAttempt uint) (*kibana.Client, error) {
	client, err := kibana.NewKibanaClient(cfg)
	if err != nil {
		if retryCfg.Enabled && (retryCfg.Maximum == 0 || retryCfg.Maximum > retryAttempt) {
			select {
			case <-ctx.Done():
				return nil, err
			case <-time.After(retryCfg.Interval):
				return getKibanaClient(ctx, cfg, retryCfg, retryAttempt+1)
			}
		}
		return nil, fmt.Errorf("Error creating Kibana client: %v", err)
	}
	return client, nil
}

// ImportIndexFile imports an index pattern from a file
func (loader KibanaLoader) ImportIndexFile(file string) error {
	if loader.version.LessThan(kibana.MinimumRequiredVersionSavedObjects) {
		return fmt.Errorf("Kibana version must be at least " + kibana.MinimumRequiredVersionSavedObjects.String())
	}

	loader.statusMsg("Importing index file from %s", file)

	// read json file
	reader, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("fail to read index-pattern from file %s: %v", file, err)
	}

	var indexContent common.MapStr
	err = json.Unmarshal(reader, &indexContent)
	if err != nil {
		return fmt.Errorf("fail to unmarshal the index content from file %s: %v", file, err)
	}

	return loader.ImportIndex(indexContent)
}

// ImportIndex imports the passed index pattern to Kibana
func (loader KibanaLoader) ImportIndex(pattern common.MapStr) error {
	if loader.version.LessThan(kibana.MinimumRequiredVersionSavedObjects) {
		return fmt.Errorf("Kibana version must be at least " + kibana.MinimumRequiredVersionSavedObjects.String())
	}

	var errs multierror.Errors

	params := url.Values{}
	params.Set("overwrite", "true")

	if err := ReplaceIndexInIndexPattern(loader.config.Index, pattern); err != nil {
		errs = append(errs, errors.Wrapf(err, "error setting index '%s' in index pattern", loader.config.Index))
	}

	if err := loader.client.ImportMultiPartFormFile(importAPI, params, "index-template.ndjson", pattern.String()); err != nil {
		errs = append(errs, errors.Wrap(err, "error loading index pattern"))
	}
	return errs.Err()
}

// ImportDashboard imports the dashboard file
func (loader KibanaLoader) ImportDashboard(file string) error {
	if loader.version.LessThan(kibana.MinimumRequiredVersionSavedObjects) {
		return fmt.Errorf("Kibana version must be at least " + kibana.MinimumRequiredVersionSavedObjects.String())
	}

	loader.statusMsg("Importing dashboard from %s", file)

	params := url.Values{}
	params.Set("overwrite", "true")

	// read json file
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("fail to read dashboard from file %s: %v", file, err)
	}

	content = ReplaceIndexInDashboardObject(loader.config.Index, content)

	content = ReplaceStringInDashboard("CHANGEME_HOSTNAME", loader.hostname, content)

	err = loader.importReferences(file, content)
	if err != nil {
		return fmt.Errorf("error loading references of dashboard: %+v", err)
	}

	var obj common.MapStr
	err = json.Unmarshal(content, &obj)
	if err != nil {
		return err
	}

	if err := loader.client.ImportMultiPartFormFile(importAPI, params, correctExtension(file), obj.String()); err != nil {
		return fmt.Errorf("error dashboard asset: %+v", err)
	}
	return nil
}

type dashboardObj struct {
	References []dashboardReference `json:"references"`
}
type dashboardReference struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (loader KibanaLoader) importReferences(path string, dashboard []byte) error {
	var d dashboardObj
	err := json.Unmarshal(dashboard, &d)
	if err != nil {
		return fmt.Errorf("failed to parse dashboard references: %+v", err)
	}

	base := filepath.Dir(path)
	for _, ref := range d.References {
		if ref.Type == "index-pattern" {
			continue
		}
		referencePath := filepath.Join(base, "..", ref.Type, ref.ID+".json")
		err := loader.ImportDashboard(referencePath)
		if err != nil {
			return fmt.Errorf("error loading reference of %s: %s %s: %+v", path, ref.Type, ref.ID, err)
		}
	}
	return nil
}

func correctExtension(file string) string {
	return filepath.Base(file[:len(file)-len("json")]) + "ndjson"
}

// Close closes the client
func (loader KibanaLoader) Close() error {
	return loader.client.Close()
}

func (loader KibanaLoader) statusMsg(msg string, a ...interface{}) {
	if loader.msgOutputter != nil {
		loader.msgOutputter(msg, a...)
	} else {
		logp.Debug("dashboards", msg, a...)
	}
}
