// Copyright 2020 The OPA Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

// Package ref implements internal helpers for references
package ref

import (
	"errors"
	"strings"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/storage"
)

// ParseDataPath returns a ref from the slash separated path s rooted at data.
// All path segments are treated as identifier strings.
func ParseDataPath(s string) (ast.Ref, error) {

	s = "/" + strings.TrimPrefix(s, "/")

	path, ok := storage.ParsePath(s)
	if !ok {
		return nil, errors.New("invalid path")
	}

	return path.Ref(ast.DefaultRootDocument), nil
}
