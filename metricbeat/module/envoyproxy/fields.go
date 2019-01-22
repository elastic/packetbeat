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

package envoyproxy

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "envoyproxy", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzEmU9v28YSwO/+FIOcHOA5z/EDHgodCgSOgQRoiqL2oehlsdodiluTO+zuULL66YvZJWlZomQ5sWgdAlvhzvxmOH/XF3CP6xmgX9K6CfSwPgNgxxXO4N3jl+/OACxGE1zDjvwMfj4DgI1TUJNtKzwDKBxWNs7S/1+A1zVuSZcPrxucwSJQ23TfjEh/KmtTXsSwxDB8PSZvr8z8uSbP2vmY0SDbkMVCZM1x4+ltik0SU7WRMahae714grQfa5/MTbnasFui6sRH2HmwF+48467iZ6zvP7+29RwDUAGmDQE9V+tOM5yvdKjRvu8tjHtRexdoa9GehvOOWFcDCSRNcI6OSwywdDq9MmfAkC/cAijA9efb988S12Rd4SaC7pXBuQAfxRewpuVEeJ2uI+gkMJxfqL2BcbLQ7DTDuSfu4nQkPkfztHAVxnVkrF8tRYuqjSVaNV8rdvWona/2mvzgDVEVk8jgdZUhYN4WRUqMgLAKjhk9MIFOZoNtUX7Lj8p5anmvVQGpQa8K7arTRt62SR3sSkfIyoV5jiA4IyRDNAbHqLIDpgXOvtWswUl2LzPxjfSTPS/oGSMM1U2F/GZu7yLnGUoWMYPDT9SYrnPOQ1IG0f2DQjzqVHAe5mvGZwpAaL0Y/GrZX5G2CkOgifJe9IFmxrrhCFxqhoCxrRiteEB7GId5whtbYzCeqGIfQbzCgNBBFG21F9a3tbrH9enHnqTkscEI8oFaQ0sMwVlU1gWFDy7yhK7sXGidhTaivPAeB6wLaJjC+jhyT/yW9NK8X26BtrXzqj8QVW7/RxswWhMqFxn9Cab3QfDJZ+Je01FD8S+HxrqB2QTUjGkCaMPxHv4O+K7P95qB5n+hYbHEyeEo/XRF4f7Q9rGNPUGJ2+KNGxVN9qf0Hl5CPsUa8hgnT/eQ40JigkXkEfDJJnKQL48iw8mX1oQXUR5YlQeAo0Ft0M47v5gOtdf4HbDd1jX9mreLOlrFt69r4IdKt3SYoqLVlDN42hnAaJ/22qoioxmhK+BRani3wmmIJQXWizQOx1LLBF5jvdk4x320fWF2lJNGBG4KtXodlYzVlSpciKwMBpksGhcmiBfRDkk7cIng8YFBAFzhkvvmKDGUO7uFlasqSGi79eGx3J2qeHwEVyTK7obRxTQF7abnf+ASSBr4ysX9oPmNqz5STlSX+x1M19R6FpcPCjuC3dVrD2mJulGyxZ2WNGDyrwXRl7fGZwkbLUeVIe/RpKnjpNdtj2pS+pcIVNmuADSBZIwA8lBSsoZ12H9Tk+v024DPiUvwuALt7a4FB9zdNiNb+CuR9mHQJVlWJREQ0ZC3+6GWGKKj3YuPV6H6mg9DwEbi00v6dAr7RjDXEa289tvrbxBw6UZpHu9f2ZSWFqrGhVa1+45hd1fUj0gpiVUXrAobMuVpX+9GbsBTdaOtr2Rurl5tPChRWwxq2ilhM/kkhrifB6R+ZKLY9btKBwk2LrXPOTmbfWFu8r9Xs9n1IOpr3VSz2S0H1HX++dunP9SXm0+fb35Xt1//vIHz///vfv/w3elVZq48qZjkTOOOdNEVYVVikCkpg4DRVTXX5l6aq/ycmj6X1DJoDzpGMi71roz6Ae5KF4GDNvdRnmg9PjRo5AkyuTOb4e48SYA1MrTeOr3wJBk7b3ennOG+8UHVGKNeOL+Y8p7Q+aWunIWABp30wiJomTHT9cvSUZV9kGMAfuq70Je7u9/+ewWxQdN5Jk1M+YoxXTACP6gUfIdMHn/gJIbmRMgvszdyMHq+ztF/oHmSqrVPo4kkdBYw5dCvfXL61UZ2S+g+ye9sj+Q1E8ksux7SPQN/6HeHYZKMbdNQ4Jj/pC1HNUNNkYF8XyzyYSgowMfLy4tr8ux8i+kJT/7i4+WlcDTkI4Ihu3Uu9XyP3W+SZJJGrjq4VO574DQe7rRBRPTSWAP+3WLkCIbSalkEqsHSyu+pWwP2m4c0B+1j7Zg3o/rfAAAA//8UAAP+"
}
