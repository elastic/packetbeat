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

package apache

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "apache", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzMl8GO2zYQhu9+ioHOjZE95OJDgTQB2qKLNtjdoIeiUGhqJBOmOSyHWsNvX5CUVa1XUmzXDKLTwlr9/zfD4Qz5BrZ4WIGwQm5wAeCV17iC4n38oVgAVMjSKesVmRX8uAAASC/hl6enTx+B0T2jgx16pySjZ5CkNUqPFdSOduA3ePxij+vu/5cLAN6Q86UkU6tmBbXQHAgcahSMK2hE+B/0XpmGV/BXwayLH6DYeG+LvxcAtUJd8SoivQEjdjgIJDz+YIOOo9Z2v4wEE54v6bMvIMl4oQxH6C4k8BvhYY8OgaUT9hhXimnZiQxhhkDshW+5/3kMKjwvoz4+E7gROQlfhjxYivQ5WNH0IYyFMQxlQ+zDXy9eHgPa4mFPrjp5NxPAoJCOwstRW09e6FJIiczIo+aaTHOZ81MQBdPu1uiAakjy4PCfFtnzHMl2ffDZOLZKU9RP+6QaBzlilhZdyShHYVgKjVVZaxL+MqiHTh4sOmCUZCY4ImkmiJ9iFs4n6HKSjaLTH8fYk9ui4+W65cONKuP3viaCaO8wa68qPb47/499EJ23b61XE23htM+d4fw5qsUGdeI33p6GJGm4lKNAs9k4gys8j2nYJX1QpqvNV5xfyc3NUcbXRdr2Vovy4dPn61ZEkzgdCnDOvjwzC4HrnkQ1swCMLjNAsIiBTlPwgT3uMnI8RoOw5NMQcqN05dCU3yAnVPd2MT9nUGXP0SlXMpzYO2QMyqA2Puqv2UO95HVbKR4/MnSSdAYZBDy9VoIPRi73ToVDeQaU90F/gAKd1deItoi2FFo95+i0CSpYYAXR5JJkSU2cNVmaGKtpoplGfHUlh44L4hmdaPDiMi7uimv3eHitTFPWQnpyK7h7+/a61A0DgJpcvBlpwR52yrQepxe1ePc907/r+Geqsrj7riO4mwihH6SSHK5JvLpqXl3Mj73i8fp88anTCxe6VPnKGm5x0OvUobXTq+pQVIFg/BZ0C4yH5DBxRX6RDjQdi9Wnt6GbJCTpw0PQn8YIDTvXSPgN0aZRMO1fGS410TZLUXw0DPdRfGYhutFT/jcbMpB8SCbnjURNTZNnGN5PKB+dGyck1q3Wh7JWRvEmD8bPvQ30NtPpCFfrUmoUJkuR/Bpu7p38zKKQDYdvTTm6xh82nLM1zfWLvYhnvLIml7dU/0xGceCcVa5ZT9zLxb8BAAD//2GFvEU="
}
