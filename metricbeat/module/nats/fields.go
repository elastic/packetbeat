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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package nats

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "nats", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJy81sFy2jAQBuA7T7GTe3gADp3ptJccwqFJz0TIi1Frad3dFYzz9B3JhVBsajcE64hl/59/W8L38BObBQSjMgNQpxUu4G5pVO5mAAWKZVero7CATzMAyDPhkYpY4Qxg47AqZJGP3EMwHo/XSkObGhdQMsX6zy89V0zjJZ30ApaCGhcERI06UWcFdGsU9sgIjKaADZOH5VvEqeBUsTP8evyxT/IPTRpfDpR8wzvDwFgZxQI8KjsrJ7PPDaeO7uzLnAFSGo/t5YA2+Q7h+7eHszl9mB7QXJB3yCtXdGYeeKLsQtlzeMCYxvMWoU2Ah6+DlkD7i4rCKL7P8NTmq/OY+mrDwDKadNYgKtbpzJu1UyM7KkDfinICndehy/LoL5pcUCyR34+ykRmDgkdP3EAUU+bulp+fn6Bmsijnr3JXaImxO+vDjCH6NXJSVVQ6ayrIgbnJUydwDAIjHrSt4+0btXX8q04ZXaeSmmplKQS0KW2SanNo1bSrBQuwlcOgw1hGTzrN02csnSgyFnAptbN2pJS5674RH6YznmLQpHPBknehBI+SHvtIHEWdQkdRS/ov3bpRnLy7HDqSNnVz42xS0T6tXYkeeZJlkRLhmHjYf6oGKOTtcZC8Va1XjL9W6SNM5kykq8juop3WP9D2dT+CfmbfOk07NqRM8BScUvqThb74AXbaL19v5u5n58wr3UxRcWp4G3qlXOJapobnzCvd6Tt6Ynb+dD9T/w4AAP//NIeGQQ=="
}
