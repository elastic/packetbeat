// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package cloudfoundry

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "cloudfoundry", asset.ModuleFieldsPri, AssetCloudfoundry); err != nil {
		panic(err)
	}
}

// AssetCloudfoundry returns asset data.
// This is the base64 encoded gzipped contents of module/cloudfoundry.
func AssetCloudfoundry() string {
	return "eJzclbFy2zwQhHs+xY178wFY/MXvTGbcpXDqGASWEkYgwAAHKXz7DCjRJilIcUIpRVioOIC7n/Zw4CPt0FckjYuqcdEq3xdErNmgoodp+aEgUgjS6461sxX9VxARPaUt9Pm4h1qnokFB5GEgAiqqwaIgajSMCtXwyiNZ0eLMND3cd6ho413sTpWM5Vxtqii67q2WE7soeHxetkgSRkuRlom3oBbstSQdSITgpBYMRQfN23Ly6pJnyqTVrDxi7dAfnF+uXYEbAZ8/kWsGtAlqeRaFdJaFtvBrAnl9U3kdBcO78imbQI137fwgTMNZHoUPhWYDCyvxTVuFH9kAjbOb30rvOUmN2Y0G0x7XSJqB2JVZKNnFspOcpWmME8uVXzTz6ctXikFsQB28hGWxQd64Ret8X9Y9I9wmC/o/aaUwYoAaDa6Zf4+OxX0QxF5oI2qDqxxKh90dIwjs/MUGDN5/JYEzjPd5jpZXT/OgMZvloXLHSU6/t70Ak+I4xif+C12DmWHSilYlY6WbBh7p1qjBB+D4hTAiMLFuMSUi7GGZnJTRe6g8IDsW5oaAgx7thYmgxvl8QqP5sG3VcRoU/qHDNPyffKei1fl7/4+Nk+IHjJddojVfnJfRa2H8MwAA//+0z6RB"
}
