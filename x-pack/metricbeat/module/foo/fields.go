// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package foo

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "foo", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJx8jTGugzAYg/ecwmLnAhne9k5RdUjBQVGT/FEIEty+AkoFiNajLX9fjScnDSuigOKKp0ZlRSoFtOyb7FJxEjX+FABwTMwuMBbjb3e1dFYEQdrBUwHW0be9XoYa0QRu8DllStTosgzp3Vw4jpA96GHyp7uCfQWuOd7Pkr2okSEW5sO2Cb3E7jT8cM75H01InqtRvQIAAP//onJfSA=="
}
