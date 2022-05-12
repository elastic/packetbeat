// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package sysmetric

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/beats/v7/x-pack/metricbeat/module/oracle"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

type sysmetricMetric struct {
	beginTime   sql.NullString
	endTime     sql.NullString
	intsizeCsec sql.NullFloat64
	groupId     sql.NullInt64
	metricId    sql.NullInt64
	name        sql.NullString
	value       sql.NullFloat64
	metricUnit  sql.NullString
	conId       sql.NullFloat64
}

/*
 * The following function executes a query that produces the following result
 *
 * BEGIN_TIM END_TIME  INTSIZE_CSEC GROUP_ID  METRIC_ID METRIC_NAME				VALUE	   METRIC_UNIT		   CON_ID
 * 19-APR-22 19-APR-22         6042        2       2146 I/O Requests per Second 2.99569679 Requests per Second 0                                                      0
 *
 * Which is parsed into sysmetricMetric instances
 */
func (e *sysmetricExtractor) calculateQuery() string {
	if len(e.patterns) == 0 {
		e.patterns = make([]interface{}, 1)
		e.patterns[0] = "%"
	}
	query := "SELECT * FROM V$SYSMETRIC WHERE METRIC_NAME LIKE :pattern0"
	for i := 1; i < len(e.patterns); i++ {
		query = query + " OR METRIC_NAME LIKE :pattern" + strconv.Itoa(i)
	}
	return query
}

func (e *sysmetricExtractor) sysmetricMetric(ctx context.Context) ([]sysmetricMetric, error) {
	query := e.calculateQuery()
	rows, err := e.db.QueryContext(ctx, query, e.patterns...)
	if err != nil {
		return nil, fmt.Errorf("error executing query %w", err)
	}

	results := make([]sysmetricMetric, 0)

	for rows.Next() {
		dest := sysmetricMetric{}
		if err = rows.Scan(&dest.beginTime, &dest.endTime, &dest.intsizeCsec, &dest.groupId, &dest.metricId, &dest.name, &dest.value, &dest.metricUnit, &dest.conId); err != nil {
			return nil, err
		}
		results = append(results, dest)
	}
	return results, nil
}

func (m *MetricSet) addSysmetricData(bs []sysmetricMetric) map[string]mapstr.M {
	out := make(map[string]mapstr.M)
	for _, sysmetricMetric := range bs {
		key := strconv.Itoa(int(sysmetricMetric.metricId.Int64)) + strconv.Itoa(int(sysmetricMetric.groupId.Int64))

		out[key] = mapstr.M{}

		oracle.SetSqlValueWithParentKey(m.Logger(), out, key, "metrics.begin_time", &oracle.StringValue{NullString: ParseDate(sysmetricMetric.beginTime)})
		oracle.SetSqlValueWithParentKey(m.Logger(), out, key, "metrics.end_time", &oracle.StringValue{NullString: ParseDate(sysmetricMetric.endTime)})
		oracle.SetSqlValueWithParentKey(m.Logger(), out, key, "metrics.interval_size_csec", &oracle.Float64Value{NullFloat64: sysmetricMetric.intsizeCsec})
		oracle.SetSqlValueWithParentKey(m.Logger(), out, key, "metrics.group_id", &oracle.Int64Value{NullInt64: sysmetricMetric.groupId})
		oracle.SetSqlValueWithParentKey(m.Logger(), out, key, "metrics.metric_id", &oracle.Int64Value{NullInt64: sysmetricMetric.metricId})
		oracle.SetSqlValueWithParentKey(m.Logger(), out, key, "metrics.name", &oracle.StringValue{NullString: ConvertToSnakeCase(sysmetricMetric.name)})
		oracle.SetSqlValueWithParentKey(m.Logger(), out, key, "metrics.value", &oracle.Float64Value{NullFloat64: sysmetricMetric.value})
		oracle.SetSqlValueWithParentKey(m.Logger(), out, key, "metrics.metric_unit", &oracle.StringValue{NullString: sysmetricMetric.metricUnit})
		oracle.SetSqlValueWithParentKey(m.Logger(), out, key, "metrics.container_id", &oracle.Float64Value{NullFloat64: sysmetricMetric.conId})
	}
	return out
}

// ParseDate function formats date according to Elastic convention
func ParseDate(date sql.NullString) sql.NullString {
	layout := "2006-01-02T15:04:05-07:00"
	t, _ := time.Parse(layout, date.String)

	return sql.NullString{String: t.UTC().Format(time.RFC3339), Valid: date.Valid}
}

// ConvertToSnakeCase function converts a string to snake case to follow
// the Elastic naming conventions in the dynamically mapped fields
func ConvertToSnakeCase(name sql.NullString) sql.NullString {
	reg, _ := regexp.Compile("[()/]") // Regex to remove '(', ')' and '/' characters from the string
	// Convert to lowercase, replace spaces and hyphens with '_' and replace '%' with 'pct'
	str := name.String
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "_")
	str = reg.ReplaceAllString(str, "")
	str = strings.ReplaceAll(str, "%", "pct")
	str = strings.ReplaceAll(str, "-", "_")
	name.String = str
	return name
}
