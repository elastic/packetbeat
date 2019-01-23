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

package haproxy

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "haproxy", asset.ModuleFieldsPri, AssetHaproxy); err != nil {
		panic(err)
	}
}

<<<<<<< HEAD
// Asset returns asset data
func Asset() string {
	return "eJzMWc2O2zYQvu9TDPbSBEj8AAs0QLFNmwBtsoe9G2NyLBGhSIUc7cZvX/BHsn5ta9dB6pslzszH+fk4Q72Hb3S4gxJrZ38cbgBYsaY7uM1Pbm8AJHnhVM3Kmjv4cAMA7Xr418pG0w3AXpGW/i6+fA8GK+orDT8+1HQHhbNNnZ/M6D0qyn+P2vbOGiYjt+Fv93ak5QtWBHYPXFInAG+sA608kyH3Fp5LJUpwJEg9kQQ0EmpnBXlPMsoJawyJoG8zRbFD8W0NiLx+FsMzevCkSXCwbKFCgwWNMIQX4Ykn90RuBlF6cTEgjZ6zTFCdkIxMJmCGZ6yxZdTbZ1SsTLFlVdG28kt2H8NiCItAGaiU1sqTsEZ68DUZhqwnvA0QntAp23j43lBDfa0pd7Q1xRTREXeEdTVMe+tmQkGecaeVL9u47JVBnR16IeLdgclvHaE8DdI01Y5cCFyUAHZofKU4p0sEp1XEXFLyoLYFKA+UVm0uBBRdFp3+cyK5CkdbYdnvvyKO70AZoRsZxB2xU5dvIVfjBc48RtfR94Y8+5YWyFGPkna0ty6wgvJgDbUeziUcDV0KrnXtz0eXLf3modB2h3olTrWCYROrhlChlI68HzP8WU4n56zbVuQ9FosmP4ZFkBeFOivC7g/w6Y+HeBIqAwJ9RBX1TffK9INnEsY2TtD84qUIldSZTeKdL+xcVZGrlMHIkJ6RF7d4b41UqUBignnfHgbKHCmmfU5GBoaZmKusXLedIABcIg9P7ZBQNTmMRfzm8f4BrINPj48Pb0+dAYu8f28NozK+oyZhG8M+eK4nDShYPXVpnPN8Gst+GwMwbH/G4JLKwaszJSesEY1zgcX62OwA1DQiIVIpMTcja3PFBjO91fVAdmG8IspMKlcAWaKROtVvv0m7ItZJQ/BSqJ6V1m1a2gH3XxFuPuVO4f3SR9qepFkO6EdNTpERrVOVP6KKMN0h1DHbVnrc157H2xV77HmW6vyz2VtXRbYD3NmG+32StBFESe2p9oLCbmGo+maKGLXCoRdr5LJl+U0+oQYLKlU4TNjZDY7l1lJtHa+3NZEaG5o4NhPL+OSdt5ZstQQ5EjlvSsmVFoYSZw1I8pwPvaVM+fO4BNQxa16REisD1cN4PlpTa6sSsG9rJHjWlwXZgci8W2an+fTrzt6/yX5+6DsbCuSSXGANDO1kJrhUrhtVJ2cP2eGr0QeoHYUJFVRqfJLij2G4VcITOlFCrZtCmdBH4BMqjTtNYdof6Gr8kCtPnePCGlaGDI8L5JT3R0VZkN2c1LMc9j6QxrA7bJW3WzHstlZDOanpEjDainHZrASxoOES446K0NW+Lh7LSi4KhuLDaxNiQcUKD7w+FU4rWoLSXfUxL3LBg6YwGKGU/eeraXaGAR352ho/BDuw/VdUAo409m5NwhQxLzwP5QwvYM2NI7kV1n5TJxv9D6OXAF/jG9RwG5T9/oS6oVugUJWgjFQizT/deJQbmRJlarKSzXZeafe0mfPaDOKSUJI72fZNIf+jPIc2MAt32sYgQDbUejzR9XFQvc1Cx8VJ2y3ECbWK1D6cJOZbxMFgOd+JM3LjX1IdIak3nUuX9SwVx2zGjtvONQk7lf3f52seVUoMe+LGGZKAbdY+Ky5BsW+39ovTNmK4MGvT2p+WtA6ft9nIVitzMkj3tqo18TBLIEj17zEDqoq4tPJdtwaNTEJP5OKo5tkpU7wUdLy+jXfwIbC24a1ExuFd/AT96nvc47cLH2Z2hH2j9ZDS34GxnNqkIBlQnBgwL9zTFbcxxJxisXe26jP8m8EWdlYe3gLumVy+rXae42eBOPS3d51rdtntUCye3I/3D/GLQmrXXzEYjb7RzH86moW7xtGksU530vxM+WoEhaA6ZXr39Uto6+nmvwAAAP//o2OPYg=="
=======
// AssetHaproxy returns asset data.
// This is the base64 encoded gzipped contents of /Users/ruflin/Dev/gopath/src/github.com/elastic/beats/filebeat/module/haproxy.
func AssetHaproxy() string {
	return "eJzMWdtu2zwSvs9TDHKzLZD6AQJsgUW2uy3w/20ucm/Q5FgalCJVcpTUb/+DB8k62lbiovWdRc7MxzkP+QG+4+EeSlE7+/NwA8DEGu/hNn+5vQFQ6KWjmsmae/h4AwDtfvjbqkbjDcCeUCt/Hxc/gBEV9pmGHx9qvIfC2abOX2b4Hhnlv0due2cNo1Hb8LdbHXH5KioEuwcusSOAd9aBJs9o0L2Hl5JkCQ4l0jMqEEZB7axE71FFOmmNQRn4baYodkJ+XwMi75/F8CI8eNQoOUi2UAkjChxhCAvhi0f3jG4GUVq4GJAWnjNNYJ2QjEQmYIZnpLFlobcvgphMsWWqcFv5JblPYTOETUAGKtKaPEprlAdfo2HIfMJqgPAsHNnGw48GG+xzTb6jrSmmiI64I6yrYdpbN2MK9Cx2mnzZ2mVPRuis0AsR7w6MfutQqNMgTVPt0AXDRQpgJ4yviLO7RHCaIuYSkwa1LYA8YNq1uRBQVFlU+q+x5CocbYRlvf8OO94BGakbFcgdsqPLj5Cj8QJlHq3r8EeDnn2bFtBhLyXtcG9dyArkwRpsNZxDOAq6FFyr2l+PLkv6l4dC253QK3HSigybsmowlVDKoffjDH82p6Nz1m0r9F4UiyI/hU2QN4U4K8LpD/D5P4+xEpIBKXxEFflNz8r4k2ccxjZO4vzmJQuV2IlN5J0u7FxUoavIiJghPQtePOKDNYpSgEQH874tBmSOKab9jkaFDDMRV1m17jiBALgUPKzawaFqdCIG8bunh0ewDj4/PT2+P1UDFvP+gzUsyPguNUnbGPZBcz1qEJLpuXPj7OdTW/bbGIBh+zMGl1gOls6EnLRGNs6FLNbHZgegphYJlkqOuRlJmws2mOmtrgeyM+MVUeakcgWQpTBKp/jtN2lXxDppCF4L1TNp3bqlHeT+K8LNVe4U3q99pG0lzXSAP2t0hEa2SiV/RBVhukOIY7Yt9bivPY+3C/bY8yzF+Rezt66K2Q7Ezjbc75OUjSBKbKvaKwK7hUH1zRSx0CSGWqwFl22W3+QKNdhQUeFEws5uUJZbSbV1vF7WhGosaKLYnFjGlXdeWpLVJsgRyXlRpFZKGFKcFaDQcy56S57y3+MWoKPXvMElVhqqh/G8tabSVjlgX9aI8KwuC7QDknm1zE7z6dfV3v+j/fLYVzYUgkt0IWuI0E7mBJfCdUN1UvYwO3wz+gC1wzChAqXGJzH+FIZbkh6FkyXUuinIhD5CPAvSYqcxTPsDXo0f5spTdVxaw2TQ8DhATml/FJQF2s1JPstm7wNpDLvDlrzdymG3tRrKSU6XgNFWjsNmJYgFDpcId1iErvZt9lhmcpExiA9vdYgFFis08HZXOM1oCUp31ce8mAseNYbBSCjV/746zc5kQIe+tsYPwQ5k/y8yAYda9G5NwhQxTzwP5UxeEDU3DtVWWvudTjb6H0eLAN/iitBwG5j9+1noBm8BQ1QCGUUyzT/deJQbmVKo1GQlme280p5pM6e1GcQlCoXuZNs3hfwXeQ5tYCbuuI1BgGqw1XhK18dB9TYTHTcnbrcQJ9QqpvbhJDHfIg4Gy/lOnAU3/jXREZx606l0mc9ScMx67LjtXOOwU9o/3l/zqFKKcCZunEEFovXaF+ISiH17tN/sthHDhV6b9v4yp3XiZZuFbDWZk0Z6sFWtkYdeAoGqf48ZUFXIpVV33R5hVCJ6RhdHNc+OTPFa0PH6Nk2qwwv4CeQnqnCgdIeVoOgb3f1Lfj26y9d55CMBL1/7ohZ1upLkF8yT8Z6c53h93tqwlfgSPbL3CNQ9jrS7u9R0fA6Z00scUs/rJb5NBIe3DW+VYHFORWvvt49vOh5NCLJ9o/Ww1N2BsZzax0AZULz9TFc8xhBzstPe2apf+d4NjrCz6vAexJ7Rje3dN/CaU3YnlIsdzdPDY3xpSWPMWwbGNNzOP6XNwlyj4LmAEFJizUOHl9r6uQFz9LL2R6D8JwAA///iRBtm"
>>>>>>> update fields.go
}
