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

package idxmgmt

import (
	"errors"
	"fmt"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/atomic"
	"github.com/elastic/beats/libbeat/idxmgmt/ilm"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/outputs/outil"
	"github.com/elastic/beats/libbeat/template"
)

type indexSupport struct {
	log         *logp.Logger
	ilm         ilm.Supporter
	info        beat.Info
	templateCfg template.TemplateConfig

	st indexState
}

type indexState struct {
	withILM atomic.Bool
}

type indexManager struct {
	support *indexSupport
	ilm     ilm.Manager

	client    ESClient
	fields    []byte
	migration bool
}

type indexSelector outil.Selector

type ilmIndexSelector struct {
	index outil.Selector
	alias outil.Selector
	st    *indexState
}

func newIndexSupport(
	log *logp.Logger,
	info beat.Info,
	ilmFactory ilm.SupportFactory,
	tmplConfig *common.Config,
	ilmConfig *common.Config,
) (*indexSupport, error) {
	if ilmFactory == nil {
		ilmFactory = ilm.DefaultSupport
	}

	ilm, err := ilmFactory(log, info, ilmConfig)
	if err != nil {
		return nil, err
	}

	tmplCfg, err := unpackTemplateConfig(tmplConfig)
	if err != nil {
		return nil, err
	}

	return &indexSupport{
		log:         log,
		ilm:         ilm,
		info:        info,
		templateCfg: tmplCfg,
	}, nil
}

func (s *indexSupport) Enabled() bool {
	return s.templateCfg.Enabled || (s.ilm.Mode() != ilm.ModeDisabled)
}

func (s *indexSupport) ILM() ilm.Supporter {
	return s.ilm
}

func (s *indexSupport) TemplateConfig(withILM bool) (template.TemplateConfig, error) {
	log := s.log

	cfg := s.templateCfg
	if mode := s.ilm.Mode(); mode == ilm.ModeDisabled {
		withILM = false
	} else if mode == ilm.ModeEnabled {
		withILM = true
	}

	var err error
	if withILM {
		ilmSettings := s.ilm.Template()
		cfg, err = applyILMSettings(log, cfg, ilmSettings)
	}
	return cfg, err
}

func (s *indexSupport) Manager(
	client ESClient,
	fields []byte,
	migration bool,
) Manager {
	ilm := s.ilm.Manager(ilm.ESClientHandler(client))
	return &indexManager{
		support:   s,
		ilm:       ilm,
		client:    client,
		fields:    fields,
		migration: migration,
	}
}

func (s *indexSupport) BuildSelector(cfg *common.Config) (outputs.IndexSelector, error) {
	log := s.log

	selCfg := common.NewConfig()
	if cfg.HasField("index") {
		s, err := cfg.String("index", -1)
		if err != nil {
			return nil, err
		}
		selCfg.SetString("index", -1, s)
	}

	if cfg.HasField("indices") {
		sub, err := cfg.Child("indices", -1)
		if err != nil {
			return nil, err
		}
		selCfg.SetChild("indices", -1, sub)
	}

	mode := s.ilm.Mode()
	alias := s.ilm.Template().Alias
	if mode == ilm.ModeEnabled {
		log.Infof("Set %v to '%s' as ILM is enabled.", cfg.PathOf("index"), alias)
		selCfg.SetString("index", -1, alias)
	}

	buildSettings := outil.Settings{
		Key:              "index",
		MultiKey:         "indices",
		EnableSingleOnly: true,
		FailEmpty:        mode != ilm.ModeEnabled,
	}

	indexSel, err := outil.BuildSelectorFromConfig(cfg, buildSettings)
	if err != nil {
		return nil, err
	}

	if mode != ilm.ModeAuto {
		return indexSelector(indexSel), nil
	}

	selCfg.SetString("index", -1, alias)
	aliasSel, err := outil.BuildSelectorFromConfig(selCfg, buildSettings)
	return &ilmIndexSelector{
		index: indexSel,
		alias: aliasSel,
		st:    &s.st,
	}, nil
}

func (m *indexManager) Setup(template, policy bool) error {
	return m.load(template, policy)
}

func (m *indexManager) Load() error {
	return m.load(false, false)
}

func (m *indexManager) load(forceTemplate, forcePolicy bool) error {
	var err error
	log := m.support.log

	withILM := m.support.st.withILM.Load()
	if !withILM {
		withILM, err = m.ilm.Enabled()
		if err != nil {
			return err
		}

		if withILM {
			log.Info("Auto ILM enable success.")
		}
	}

	// mark ILM as enabled in indexState if withILM is true
	if withILM {
		m.support.st.withILM.CAS(false, withILM)
	}

	// install ilm policy
	if withILM {
		if err := m.ilm.EnsurePolicy(forcePolicy); err != nil {
			return err
		}
		log.Info("ILM policy successfully loaded.")
	}

	// create and install template
	if m.support.templateCfg.Enabled {
		tmplCfg := m.support.templateCfg
		if withILM {
			ilmSettings := m.support.ilm.Template()
			tmplCfg, err = applyILMSettings(log, tmplCfg, ilmSettings)
			if err != nil {
				return err
			}
		}

		if forceTemplate {
			tmplCfg.Overwrite = true
		}

		loader, err := template.NewLoader(tmplCfg, m.client, m.support.info, m.fields, m.migration)
		if err != nil {
			return fmt.Errorf("Error creating Elasticsearch template loader: %v", err)
		}

		err = loader.Load()
		if err != nil {
			return fmt.Errorf("Error loading Elasticsearch template: %v", err)
		}

		log.Info("Loaded index template.")
	}

	// create alias
	if withILM {
		if err := m.ilm.EnsureAlias(); err != nil {
			return err
		}
		log.Info("Write alias successfully generated.")
	}

	return nil
}

func (s *ilmIndexSelector) Select(evt *beat.Event) (string, error) {
	if idx := getEventCustomIndex(evt); idx != "" {
		return idx, nil
	}

	if s.st.withILM.Load() {
		idx, err := s.alias.Select(evt)
		return idx, err
	}

	idx, err := s.index.Select(evt)
	return idx, err
}

func (s indexSelector) Select(evt *beat.Event) (string, error) {
	if idx := getEventCustomIndex(evt); idx != "" {
		return idx, nil
	}
	return outil.Selector(s).Select(evt)
}

func getEventCustomIndex(evt *beat.Event) string {
	if len(evt.Meta) == 0 {
		return ""
	}

	if tmp := evt.Meta["alias"]; tmp != nil {
		if alias, ok := tmp.(string); ok {
			return alias
		}
	}

	if tmp := evt.Meta["index"]; tmp != nil {
		if idx, ok := tmp.(string); ok {
			ts := evt.Timestamp.UTC()
			return fmt.Sprintf("%s-%d.%02d.%02d",
				idx, ts.Year(), ts.Month(), ts.Day())
		}
	}

	return ""
}

func unpackTemplateConfig(cfg *common.Config) (template.TemplateConfig, error) {
	if cfg == nil {
		cfg = common.NewConfig()
	}

	config := template.DefaultConfig
	err := cfg.Unpack(&config)
	return config, err
}

func applyILMSettings(
	log *logp.Logger,
	tmpl template.TemplateConfig,
	settings ilm.TemplateSettings,
) (template.TemplateConfig, error) {
	if !tmpl.Enabled {
		return tmpl, nil
	}

	alias := settings.Alias
	if alias == "" {
		return tmpl, errors.New("no ilm rollover alias configured")
	}

	policy := settings.PolicyName
	if policy == "" {
		return tmpl, errors.New("no ilm policy name configured")
	}

	tmpl.Name = alias
	if log != nil {
		log.Infof("Set setup.template.name to '%s' as ILM is enabled.", alias)
	}

	tmpl.Pattern = fmt.Sprintf("%s-*", alias)
	if log != nil {
		log.Infof("Set setup.template.pattern to '%s' as ILM is enabled.", tmpl.Pattern)
	}

	// rollover_alias and lifecycle.name can't be configured and will be overwritten

	// init/copy index settings
	idxSettings := tmpl.Settings.Index
	if idxSettings == nil {
		idxSettings = map[string]interface{}{}
	} else {
		tmp := make(map[string]interface{}, len(idxSettings))
		for k, v := range idxSettings {
			tmp[k] = v
		}
		idxSettings = tmp
	}
	tmpl.Settings.Index = idxSettings

	// init/copy index.lifecycle settings
	var lifecycle map[string]interface{}
	if ifcLifecycle := idxSettings["lifecycle"]; ifcLifecycle == nil {
		lifecycle = map[string]interface{}{}
	} else if tmp, ok := ifcLifecycle.(map[string]interface{}); ok {
		lifecycle = make(map[string]interface{}, len(tmp))
		for k, v := range tmp {
			lifecycle[k] = v
		}
	} else {
		return tmpl, errors.New("settings.index.lifecycle must be an object")
	}
	idxSettings["lifecycle"] = lifecycle

	// add rollover_alias and name to index.lifecycle settings
	if _, exists := lifecycle["rollover_alias"]; !exists {
		log.Infof("Set settings.index.lifecycle.rollover_alias in template to %s as ILM is enabled.", alias)
		lifecycle["rollover_alias"] = alias
	}
	if _, exists := lifecycle["name"]; !exists {
		log.Infof("Set settings.index.lifecycle.name in template to %s as ILM is enabled.", policy)
		lifecycle["name"] = policy
	}

	return tmpl, nil
}
