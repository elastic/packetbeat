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

package httpcommon

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/elastic/beats/v7/libbeat/common"
)

// ProxyURI is a URL used for proxy serialized as a string.
type ProxyURI url.URL

// MarshalYAML serializes URI as a string.
func (p ProxyURI) MarshalYAML() (interface{}, error) {
	u := url.URL(p)
	return u.String(), nil
}

// MarshalJSON serializes URI as a string.
func (p ProxyURI) MarshalJSON() ([]byte, error) {
	u := url.URL(p)
	fmt.Println(u.String())
	return []byte(fmt.Sprintf("\"%s\"", u.String())), nil
}

// Unpack unpacks string into an proxy URI.
func (p *ProxyURI) Unpack(s string) error {
	uri, err := stringToProxyURI(s)
	if err != nil {
		return err
	}

	*p = uri
	return nil
}

// UnmarshalJSON unpacks string into an proxy URI.
func (p *ProxyURI) UnmarshalJSON(b []byte) error {
	unqoted := strings.Trim(string(b), `"`)
	uri, err := stringToProxyURI(unqoted)
	if err != nil {
		return err
	}

	*p = uri
	return nil
}

// UnmarshalYAML unpacks string into an proxy URI.
func (p *ProxyURI) UnmarshalYAML(unmarshal func(interface{}) error) error {
	rawURI := ""
	if err := unmarshal(&rawURI); err != nil {
		return err
	}

	uri, err := stringToProxyURI(rawURI)
	if err != nil {
		return err
	}

	*p = uri
	return nil
}

// URI returns conventional url.URL structure.
func (p *ProxyURI) URI() *url.URL {
	return (*url.URL)(p)
}

func stringToProxyURI(s string) (ProxyURI, error) {
	proxyURI, err := common.ParseURL(s)
	if err != nil {
		return ProxyURI{}, err
	}

	return ProxyURI(*proxyURI), nil
}
