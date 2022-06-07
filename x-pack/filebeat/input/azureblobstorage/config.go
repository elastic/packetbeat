// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package azureblobstorage

import "github.com/elastic/beats/v7/filebeat/harvester"

type config struct {
	ForwaderConfig harvester.ForwarderConfig `config:",inline"`
	AccountName    string                    `config:"account_name"`
	AccountKey     string                    `config:"account_key"`
}

func defaultConfig() config {
	return config{
		ForwaderConfig: harvester.ForwarderConfig{
			Type: "azureblobstorage",
		},
		AccountName: "beatsblobstorage",
		AccountKey:  "61A0frq/mFUSw6BGivRB8jhOiElUwGcMlI5lCbXruJokvYIWUcwvpp9ln6v7MPBzwsfvprCEt2qA+AStH+iVXw==",
	}
}
