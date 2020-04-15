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

// +build windows

package perfmon

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/elastic/beats/v7/metricbeat/helper/windows/pdh"

	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

const (
	instanceCountLabel   = ":count"
	defaultInstanceField = "instance"
	defaultObjectField   = "object"
)

// Reader will contain the config options
type Reader struct {
	query    pdh.Query    // PDH Query
	executed bool         // Indicates if the query has been executed.
	log      *logp.Logger //
	config   Config       // Metricset configuration
	counters []PerfCounter
}

type PerfCounter struct {
	InstanceField string
	InstanceName  string
	QueryField    string
	QueryName     string
	Format        string
	ObjectName    string
	ObjectField   string
	ChildQueries  []string
}

// NewReader creates a new instance of Reader.
func NewReader(config Config) (*Reader, error) {
	var query pdh.Query
	if err := query.Open(); err != nil {
		return nil, err
	}
	r := &Reader{
		query:  query,
		log:    logp.NewLogger("perfmon"),
		config: config,
	}
	r.mapCounters(config)
	for i, counter := range r.counters {
		r.counters[i].ChildQueries = []string{}
		childQueries, err := query.GetCounterPaths(counter.QueryName)
		if err != nil {
			if config.IgnoreNECounters {
				switch err {
				case pdh.PDH_CSTATUS_NO_COUNTER, pdh.PDH_CSTATUS_NO_COUNTERNAME,
					pdh.PDH_CSTATUS_NO_INSTANCE, pdh.PDH_CSTATUS_NO_OBJECT:
					r.log.Infow("Ignoring non existent counter", "error", err,
						logp.Namespace("perfmon"), "query", counter.QueryName)
					continue
				}
			} else {
				query.Close()
				return nil, errors.Wrapf(err, `failed to expand counter (query="%v")`, counter.QueryName)
			}
		}
		// check if the pdhexpandcounterpath/pdhexpandwildcardpath functions have expanded the counter successfully.
		if len(childQueries) == 0 || (len(childQueries) == 1 && strings.Contains(childQueries[0], "*")) {
			// covering cases when PdhExpandWildCardPathW returns no counter paths or is unable to expand and the ignore_non_existent_counters flag is set
			if config.IgnoreNECounters {
				r.log.Infow("Ignoring non existent counter", "initial query", counter.QueryName,
					logp.Namespace("perfmon"), "expanded query", childQueries)
				continue
			}
			return nil, errors.Errorf(`failed to expand counter (query="%v"), no error returned`, counter.QueryName)
		}
		for _, v := range childQueries {
			if err := query.AddCounter(v, counter.InstanceName, counter.Format, len(childQueries) > 1); err != nil {
				return nil, errors.Wrapf(err, `failed to add counter (query="%v")`, counter.QueryName)
			}
			r.counters[i].ChildQueries = append(r.counters[i].ChildQueries, v)
		}
	}
	return r, nil
}

// RefreshCounterPaths will recheck for any new instances and add them to the counter list
func (re *Reader) RefreshCounterPaths() error {
	var newCounters []string
	for i, counter := range re.counters {
		re.counters[i].ChildQueries = []string{}
		childQueries, err := re.query.GetCounterPaths(counter.QueryName)
		if err != nil {
			if re.config.IgnoreNECounters {
				switch err {
				case pdh.PDH_CSTATUS_NO_COUNTER, pdh.PDH_CSTATUS_NO_COUNTERNAME,
					pdh.PDH_CSTATUS_NO_INSTANCE, pdh.PDH_CSTATUS_NO_OBJECT:
					re.log.Infow("Ignoring non existent counter", "error", err,
						logp.Namespace("perfmon"), "query", counter.QueryName)
					continue
				}
			} else {
				return errors.Wrapf(err, `failed to expand counter (query="%v")`, counter.QueryName)
			}
		}
		newCounters = append(newCounters, childQueries...)
		// there are cases when the ExpandWildCardPath will retrieve a successful status but not an expanded query so we need to check for the size of the list
		if err == nil && len(childQueries) >= 1 && !strings.Contains(childQueries[0], "*") {
			for _, v := range childQueries {
				if err := re.query.AddCounter(v, counter.InstanceName, counter.Format, len(childQueries) > 1); err != nil {
					return errors.Wrapf(err, "failed to add counter (query='%v')", counter.QueryName)
				}
				re.counters[i].ChildQueries = append(re.counters[i].ChildQueries, v)
			}
		}
	}
	err := re.query.RemoveUnusedCounters(newCounters)
	if err != nil {
		return errors.Wrap(err, "failed removing unused counter values")
	}

	return nil
}

// Read executes a query and returns those values in an event.
func (re *Reader) Read() ([]mb.Event, error) {
	// Some counters, such as rate counters, require two counter values in order to compute a displayable value. In this case we must call PdhCollectQueryData twice before calling PdhGetFormattedCounterValue.
	// For more information, see Collecting Performance Data (https://docs.microsoft.com/en-us/windows/desktop/PerfCtrs/collecting-performance-data).
	if err := re.query.CollectData(); err != nil {
		return nil, errors.Wrap(err, "failed querying counter values")
	}

	// Get the values.
	values, err := re.query.GetFormattedCounterValues()
	if err != nil {
		return nil, errors.Wrap(err, "failed formatting counter values")
	}
	var events []mb.Event
	// GroupAllCountersTo config option where counters for all instances are aggregated and instance count is added in the event under the string value provided by this option.
	if re.config.GroupAllCountersTo != "" {
		event := re.groupToSingleEvent(values)
		events = append(events, event)
	} else {
		events = re.groupToEvents(values)
	}
	re.executed = true
	return events, nil
}

// Close will close the PDH query for now.
func (re *Reader) Close() error {
	return re.query.Close()
}

func (re *Reader) getCounter(query string) (bool, PerfCounter) {
	for _, counter := range re.counters {
		for _, childQuery := range counter.ChildQueries {
			if childQuery == query {
				return true, counter
			}
		}
	}
	return false, PerfCounter{}
}

func (re *Reader) mapCounters(config Config) {
	re.counters = []PerfCounter{}
	if len(config.Counters) > 0 {
		for _, counter := range config.Counters {
			re.counters = append(re.counters, PerfCounter{
				InstanceField: counter.InstanceLabel,
				InstanceName:  counter.InstanceName,
				QueryField:    counter.MeasurementLabel,
				QueryName:     counter.Query,
				Format:        counter.Format,
				ChildQueries:  nil,
			})
		}
	}
	if len(config.Queries) > 0 {
		for _, query := range config.Queries {
			for _, counter := range query.Counters {
				// counter paths can also not contain any instances
				if len(query.Instance) == 0 {
					re.counters = append(re.counters, PerfCounter{
						InstanceField: defaultInstanceField,
						InstanceName:  "",
						QueryField:    mapCounterPathLabel(query.Namespace, counter.Field, counter.Name),
						QueryName:     mapQuery(query.Name, "", counter.Name),
						Format:        counter.Format,
						ObjectName:    query.Name,
						ObjectField:   mapObjectName(query.Field),
					})
				} else {
					for _, instance := range query.Instance {
						re.counters = append(re.counters, PerfCounter{
							InstanceField: defaultInstanceField,
							InstanceName:  instance,
							QueryField:    mapCounterPathLabel(query.Namespace, counter.Field, counter.Name),
							QueryName:     mapQuery(query.Name, instance, counter.Name),
							Format:        counter.Format,
							ObjectName:    query.Name,
							ObjectField:   mapObjectName(query.Field),
						})
					}
				}
			}
		}
	}
}

func mapObjectName(objectField string) string {
	if objectField != "" {
		return objectField
	}
	return defaultObjectField
}

func mapQuery(obj string, instance string, path string) string {
	var query string
	// trim object
	obj = strings.TrimPrefix(obj, "\\")
	obj = strings.TrimSuffix(obj, "\\")
	query = fmt.Sprintf("\\%s", obj)

	if instance != "" {
		// trim instance
		instance = strings.TrimPrefix(instance, "(")
		instance = strings.TrimSuffix(instance, ")")
		query += fmt.Sprintf("(%s)", instance)
	}

	if strings.HasPrefix(path, "\\") {
		query += path
	} else {
		query += fmt.Sprintf("\\%s", path)
	}
	return query
}

func mapCounterPathLabel(namespace string, label string, path string) string {
	if label != "" {
		return namespace + "." + label
	}
	// replace spaces with underscores
	path = strings.Replace(path, " ", "_", -1)
	// replace backslashes with "per"
	path = strings.Replace(path, "/sec", "_per_sec", -1)
	path = strings.Replace(path, "/_sec", "_per_sec", -1)
	path = strings.Replace(path, "\\", "_", -1)
	// replace actual percentage symbol with the smbol "pct"
	path = strings.Replace(path, "_%_", "_pct_", -1)
	// create an object in case of ":"
	path = strings.Replace(path, ":", "_", -1)
	// create an object in case of ":"
	path = strings.Replace(path, "_-_", "_", -1)
	// replace uppercases with underscores
	path = replaceUpperCase(path)

	//  avoid cases as this "logicaldisk_avg._disk_sec_per_transfer"
	obj := strings.Split(path, ".")
	for index := range obj {
		// in some cases a trailing "_" is found
		obj[index] = strings.TrimPrefix(obj[index], "_")
		obj[index] = strings.TrimSuffix(obj[index], "_")
	}
	path = strings.ToLower(strings.Join(obj, "_"))

	//  avoid cases as this "logicaldisk_avg__disk_sec_per_transfer"
	path = strings.Replace(path, "__", "_", -1)
	return namespace + "." + path
}

// replaceUpperCase func will replace upper case with '_'
func replaceUpperCase(src string) string {
	replaceUpperCaseRegexp := regexp.MustCompile(replaceUpperCaseRegex)
	return replaceUpperCaseRegexp.ReplaceAllStringFunc(src, func(str string) string {
		var newStr string
		for _, r := range str {
			// split into fields based on class of unicode character
			if unicode.IsUpper(r) {
				newStr += "_" + strings.ToLower(string(r))
			} else {
				newStr += string(r)
			}
		}
		return newStr
	})
}
