// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by "stringer -type=stateChange -linecomment=true"; DO NOT EDIT.

package stateresolver

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[startState-1]
	_ = x[updateState-2]
	_ = x[unchangedState-3]
}

const _stateChange_name = "STARTUPDATEUNCHANGE"

var _stateChange_index = [...]uint8{0, 5, 11, 19}

func (i stateChange) String() string {
	i -= 1
	if i >= stateChange(len(_stateChange_index)-1) {
		return "stateChange(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _stateChange_name[_stateChange_index[i]:_stateChange_index[i+1]]
}
