// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package counter

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/elastic/beats/v7/libbeat/paths"
)

func TestMain(m *testing.M) {
	tmpdir, err := ioutil.TempDir("", "beats-data-dir")
	if err != nil {
		fmt.Printf("Failed to create temporal data directory: %v\n", err)
		os.Exit(1)
	}
	paths.Paths.Data = tmpdir

	result := m.Run()
	os.RemoveAll(tmpdir)

	os.Exit(result)
}
