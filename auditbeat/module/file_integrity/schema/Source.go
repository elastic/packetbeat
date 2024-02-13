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

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package schema

import "strconv"

type Source byte

const (
	SourceScan     Source = 0
	SourceFSNotify Source = 1
	SourceEBPF     Source = 2
)

var EnumNamesSource = map[Source]string{
	SourceScan:     "Scan",
	SourceFSNotify: "FSNotify",
	SourceEBPF:     "eBPF",
}

var EnumValuesSource = map[string]Source{
	"Scan":     SourceScan,
	"FSNotify": SourceFSNotify,
	"eBPF":     SourceEBPF,
}

func (v Source) String() string {
	if s, ok := EnumNamesSource[v]; ok {
		return s
	}
	return "Source(" + strconv.FormatInt(int64(v), 10) + ")"
}
