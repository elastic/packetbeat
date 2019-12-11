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

package mysql

import (
	"github.com/go-sql-driver/mysql"
)

// Helper functions for testing used in the mysql MetricSets.

// GetMySQLEnvDSN returns the MySQL server DSN to use for testing. It
// reads the value from the MYSQL_DSN environment variable and returns
// root@tcp(127.0.0.1:3306)/ if it is not set.
func GetMySQLEnvDSN(host string) string {
	c := mysql.NewConfig()
	c.Net = "tcp"
	c.Addr = host
	c.User = "root"
	c.Passwd = "test"
	return c.FormatDSN()
}
