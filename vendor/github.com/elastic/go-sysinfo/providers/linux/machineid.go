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

package linux

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	"github.com/elastic/go-sysinfo/types"
)

var (
	// Possible (current and historic) locations of the machine-id file.
	// These will be searched in order.
	machineIDFiles = []string{"/etc/machine-id", "/var/lib/dbus/machine-id", "/var/db/dbus/machine-id"}
)

func MachineID() (string, error) {
	var contents []byte
	var err error

	for _, file := range machineIDFiles {
		contents, err = ioutil.ReadFile(file)
		if err != nil {
			if os.IsNotExist(err) {
				// Try next location
				continue
			}

			// Return with error on any other error
			return "", errors.Wrapf(err, "failed to read %v", file)
		}

		// Found it
		break
	}

	if os.IsNotExist(err) {
		// None of the locations existed
		return "", types.ErrNotImplemented
	}

	contents = bytes.TrimSpace(contents)
	return string(contents), nil
}
