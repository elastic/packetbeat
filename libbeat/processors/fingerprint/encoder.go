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

package fingerprint

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
)

var errEncoderUnknown = errors.New("unknown encoding method")

type encoder func([]byte) string

var encodings = map[string]encoder{
	"hex":    hex.EncodeToString,
	"base32": base32.StdEncoding.EncodeToString,
	"base64": base64.StdEncoding.EncodeToString,
}

// Unpack creates the Method enumeration value from the given string
func (e *encoder) Unpack(str string) error {
	str = strings.ToLower(str)

	m, found := encodings[str]
	if !found {
		return errEncoderUnknown
	}

	*e = m
	return nil
}
