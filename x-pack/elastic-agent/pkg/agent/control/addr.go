// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build !windows

package control

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/agent/application/paths"
)

// Address returns the address to connect to Elastic Agent daemon.
func Address() string {
	data := paths.Data()
	// entire string cannot be longer than 107 characters, this forces the
	// length to always be 88 characters (but unique per data path)
	path := filepath.Join(os.TempDir(), fmt.Sprintf("elastic-agent-%x.sock", sha256.Sum256([]byte(data))))
	return fmt.Sprintf(`unix://%s`, path)
}
