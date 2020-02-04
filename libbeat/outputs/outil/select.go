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

package outil

import (
	"fmt"
	"strings"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/fmtstr"
	"github.com/elastic/beats/libbeat/conditions"
)

type Selector struct {
	sel selectorExpr
}

type Settings struct {
	// single selector key and default option keyword
	Key string

	// multi-selector key in config
	MultiKey string

	// if enabled a selector `key` in config will be generated, if `key` is present
	EnableSingleOnly bool

	// Fail building selector if `key` and `multiKey` are missing
	FailEmpty bool
}

type selectorExpr interface {
	sel(evt *beat.Event) (string, error)
}

type emptySelector struct{}

type listSelector struct {
	selectors []selectorExpr
}

type condSelector struct {
	s    selectorExpr
	cond conditions.Condition
}

type constSelector struct {
	s string
}

type fmtSelector struct {
	f         fmtstr.EventFormatString
	otherwise string
}

type mapSelector struct {
	from      selectorExpr
	otherwise string
	to        map[string]string
}

var nilSelector selectorExpr = &emptySelector{}

func MakeSelector(es ...selectorExpr) Selector {
	switch len(es) {
	case 0:
		return Selector{nilSelector}
	case 1:
		return Selector{es[0]}
	default:
		return Selector{concatSelectorExpr(es...)}
	}
}

// Select runs configured selector against the current event.
// If no matching selector is found, an empty string is returned.
// It's up to the caller to decide if an empty string is an error
// or an expected result.
func (s Selector) Select(evt *beat.Event) (string, error) {
	return s.sel.sel(evt)
}

func (s Selector) IsEmpty() bool {
	return s.sel == nilSelector || s.sel == nil
}

func (s Selector) IsConst() bool {
	if s.sel == nilSelector {
		return true
	}

	_, ok := s.sel.(*constSelector)
	return ok
}

func BuildSelectorFromConfig(
	cfg *common.Config,
	settings Settings,
) (Selector, error) {
	var sel []selectorExpr

	key := settings.Key
	multiKey := settings.MultiKey
	found := false

	if cfg.HasField(multiKey) {
		found = true
		sub, err := cfg.Child(multiKey, -1)
		if err != nil {
			return Selector{}, err
		}

		var table []*common.Config
		if err := sub.Unpack(&table); err != nil {
			return Selector{}, err
		}

		for _, config := range table {
			action, err := buildSingle(config, key)
			if err != nil {
				return Selector{}, err
			}

			if action != nilSelector {
				sel = append(sel, action)
			}
		}
	}

	if settings.EnableSingleOnly && cfg.HasField(key) {
		found = true

		// expect event-format-string
		str, err := cfg.String(key, -1)
		if err != nil {
			return Selector{}, err
		}

		fmtstr, err := fmtstr.CompileEvent(str)
		if err != nil {
			return Selector{}, fmt.Errorf("%v in %v", err, cfg.PathOf(key))
		}

		if fmtstr.IsConst() {
			str, err := fmtstr.Run(nil)
			if err != nil {
				return Selector{}, err
			}

			if str != "" {
				sel = append(sel, constSelectorExpr(str))
			}
		} else {
			sel = append(sel, fmtSelectorExpr(fmtstr, ""))
		}
	}

	if settings.FailEmpty && !found {
		if settings.EnableSingleOnly {
			return Selector{}, fmt.Errorf("missing required '%v' or '%v' in %v",
				key, multiKey, cfg.Path())
		}

		return Selector{}, fmt.Errorf("missing required '%v' in %v",
			multiKey, cfg.Path())
	}

	return MakeSelector(sel...), nil
}

func emptySelectorExpr() selectorExpr {
	return nilSelector
}

func constSelectorExpr(s string) selectorExpr {
	if s == "" {
		return emptySelectorExpr()
	}
	return &constSelector{strings.ToLower(s)}
}

func fmtSelectorExpr(fmt *fmtstr.EventFormatString, fallback string) selectorExpr {
	return &fmtSelector{*fmt, strings.ToLower(fallback)}
}

func concatSelectorExpr(s ...selectorExpr) selectorExpr {
	return &listSelector{s}
}

func conditionalSelectorExpr(
	s selectorExpr,
	cond conditions.Condition,
) selectorExpr {
	return &condSelector{s, cond}
}

func lookupSelectorExpr(
	evtfmt *fmtstr.EventFormatString,
	table map[string]string,
	fallback string,
) (selectorExpr, error) {
	if evtfmt.IsConst() {
		str, err := evtfmt.Run(nil)
		if err != nil {
			return nil, err
		}

		str = table[strings.ToLower(str)]
		if str == "" {
			str = fallback
		}
		return constSelectorExpr(str), nil
	}

	return &mapSelector{
		from:      fmtSelectorExpr(evtfmt, ""),
		to:        table,
		otherwise: fallback,
	}, nil
}

func lowercaseTable(table map[string]string) map[string]string {
	tmp := make(map[string]string, len(table))
	for k, v := range table {
		tmp[strings.ToLower(k)] = strings.ToLower(v)
	}
	return tmp
}

func buildSingle(cfg *common.Config, key string) (selectorExpr, error) {
	// TODO: check for unknown fields

	// 1. extract required key-word handler
	if !cfg.HasField(key) {
		return nil, fmt.Errorf("missing %v", cfg.PathOf(key))
	}

	str, err := cfg.String(key, -1)
	if err != nil {
		return nil, err
	}

	evtfmt, err := fmtstr.CompileEvent(str)
	if err != nil {
		return nil, fmt.Errorf("%v in %v", err, cfg.PathOf(key))
	}

	// 2. extract optional `default` value
	var otherwise string
	if cfg.HasField("default") {
		tmp, err := cfg.String("default", -1)
		if err != nil {
			return nil, err
		}
		otherwise = strings.ToLower(tmp)
	}

	// 3. extract optional `mapping`
	mapping := struct {
		Table map[string]string `config:"mappings"`
	}{nil}
	if cfg.HasField("mappings") {
		if err := cfg.Unpack(&mapping); err != nil {
			return nil, err
		}
	}

	// 4. extract conditional
	var cond conditions.Condition
	if cfg.HasField("when") {
		sub, err := cfg.Child("when", -1)
		if err != nil {
			return nil, err
		}

		condConfig := conditions.Config{}
		if err := sub.Unpack(&condConfig); err != nil {
			return nil, err
		}

		tmp, err := conditions.NewCondition(&condConfig)
		if err != nil {
			return nil, err
		}

		cond = tmp
	}

	// 5. build selector from available fields
	var sel selectorExpr
	if len(mapping.Table) > 0 {
		var err error
		sel, err = lookupSelectorExpr(evtfmt, lowercaseTable(mapping.Table), otherwise)
		if err != nil {
			return nil, err
		}
	} else {
		if evtfmt.IsConst() {
			str, err := evtfmt.Run(nil)
			if err != nil {
				return nil, err
			}
			sel = constSelectorExpr(str)
		} else {
			sel = fmtSelectorExpr(evtfmt, otherwise)
		}
	}

	if cond != nil && sel != nilSelector {
		sel = conditionalSelectorExpr(sel, cond)
	}

	return sel, nil
}

func (s *emptySelector) sel(evt *beat.Event) (string, error) {
	return "", nil
}

func (s *listSelector) sel(evt *beat.Event) (string, error) {
	for _, sub := range s.selectors {
		n, err := sub.sel(evt)
		if err != nil { // TODO: try
			return n, err
		}

		if n != "" {
			return n, nil
		}
	}

	return "", nil
}

func (s *condSelector) sel(evt *beat.Event) (string, error) {
	if !s.cond.Check(evt) {
		return "", nil
	}
	return s.s.sel(evt)
}

func (s *constSelector) sel(_ *beat.Event) (string, error) {
	return s.s, nil
}

func (s *fmtSelector) sel(evt *beat.Event) (string, error) {
	n, err := s.f.Run(evt)
	if err != nil {
		// err will be set if not all keys present in event ->
		// return empty selector result and ignore error
		return s.otherwise, nil
	}

	if n == "" {
		return s.otherwise, nil
	}
	return strings.ToLower(n), nil
}

func (s *mapSelector) sel(evt *beat.Event) (string, error) {
	n, err := s.from.sel(evt)
	if err != nil {
		if s.otherwise == "" {
			return "", err
		}
		return s.otherwise, nil
	}

	if n == "" {
		return s.otherwise, nil
	}

	n = s.to[n]
	if n == "" {
		return s.otherwise, nil
	}
	return n, nil
}
