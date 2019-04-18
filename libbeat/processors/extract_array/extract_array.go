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

package extract_array

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/processors"
	"github.com/elastic/beats/libbeat/processors/checks"
)

type extractArrayConfig struct {
	Field         string        `config:"field"`
	Mappings      common.MapStr `config:"mappings"`
	IgnoreMissing bool          `config:"ignore_missing"`
	OverwriteKeys bool          `config:"overwrite_keys"`
	FailOnError   bool          `config:"fail_on_error"`
}

type fieldMapping struct {
	from int
	to   string
}

type extractArrayProcessor struct {
	config   extractArrayConfig
	mappings []fieldMapping
}

var (
	defaultExtractArrayConfig = extractArrayConfig{
		FailOnError: true,
	}
	errNoMappings = errors.New("no mappings defined in extract_array processor")
)

func init() {
	processors.RegisterPlugin("extract_array",
		checks.ConfigChecked(NewExtractArray,
			checks.RequireFields("field", "mappings"),
			checks.AllowedFields("field", "mappings", "ignore_missing", "overwrite_keys", "fail_on_error", "when")))
}

// NewExtractArray builds a new extract_array processor.
func NewExtractArray(c *common.Config) (processors.Processor, error) {
	config := defaultExtractArrayConfig

	err := c.Unpack(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack the extract_array configuration: %s", err)
	}

	if len(config.Mappings) == 0 {
		return nil, errNoMappings
	}

	p := &extractArrayProcessor{
		config:   config,
		mappings: make([]fieldMapping, 0, len(config.Mappings)),
	}
	for field, column := range config.Mappings.Flatten() {
		colIdx, ok := common.TryToInt(column)
		if !ok || colIdx < 0 {
			return nil, fmt.Errorf("bad extract_array mapping for field %s: %+v is not a positive integer", field, column)
		}
		p.mappings = append(p.mappings, fieldMapping{from: colIdx, to: field})
	}
	sort.Slice(p.mappings, func(i, j int) bool {
		return p.mappings[i].from < p.mappings[j].from
	})
	return p, nil
}

func (f *extractArrayProcessor) Run(event *beat.Event) (*beat.Event, error) {
	iValue, err := event.GetValue(f.config.Field)
	if err != nil {
		if f.config.IgnoreMissing && errors.Cause(err) == common.ErrKeyNotFound {
			return event, nil
		}
		return event, errors.Wrapf(err, "could not fetch value for field %s", f.config.Field)
	}

	v := reflect.ValueOf(iValue)
	if t := v.Type(); t.Kind() != reflect.Slice {
		if !f.config.FailOnError {
			return event, nil
		}
		return event, errors.Wrapf(err, "unsupported type for field %s: got: %s needed: array", f.config.Field, t.String())
	}

	saved := *event
	if f.config.FailOnError {
		saved.Fields = event.Fields.Clone()
		saved.Meta = event.Meta.Clone()
	}

	n := v.Len()
	for _, mapping := range f.mappings {
		if mapping.from >= n {
			if !f.config.FailOnError {
				continue
			}
			return &saved, errors.Errorf("index %d exceeds length of %d when processing mapping for field %s", mapping.from, n, mapping.to)
		}
		if !f.config.OverwriteKeys {
			if _, err = event.GetValue(mapping.to); err == nil {
				if !f.config.FailOnError {
					continue
				}
				return &saved, errors.Errorf("target field %s already has a value. Set the overwrite_keys flag or drop/rename the field first", mapping.to)
			}
		}
		if _, err = event.PutValue(mapping.to, clone(v.Index(mapping.from).Interface())); err != nil {
			if !f.config.FailOnError {
				continue
			}
			return &saved, errors.Wrapf(err, "failed setting field %s", mapping.to)
		}
	}
	return event, nil
}

func (f *extractArrayProcessor) String() (r string) {
	return fmt.Sprintf("extract_array={field=%s, mappings=%v}", f.config.Field, f.mappings)
}

func clone(value interface{}) interface{} {
	// TODO: This is dangerous but done by most processors.
	//       Otherwise need to reflect value and deep copy lists / map types.
	switch v := value.(type) {
	case common.MapStr:
		return v.Clone()
	case map[string]interface{}:
		return common.MapStr(v).Clone()
	}
	return value
}
