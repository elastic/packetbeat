// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build windows

package install

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasRoot(t *testing.T) {
	t.Run("check if user is admin", func(t *testing.T) {
		result, err := HasRoot()
		assert.NoError(t, err)
		assert.True(t, result)
	})
}
