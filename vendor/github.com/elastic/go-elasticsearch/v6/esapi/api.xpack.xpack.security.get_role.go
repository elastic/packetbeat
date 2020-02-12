// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 6.8.5: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newXPackSecurityGetRoleFunc(t Transport) XPackSecurityGetRole {
	return func(o ...func(*XPackSecurityGetRoleRequest)) (*Response, error) {
		var r = XPackSecurityGetRoleRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// XPackSecurityGetRole - https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-role.html
//
type XPackSecurityGetRole func(o ...func(*XPackSecurityGetRoleRequest)) (*Response, error)

// XPackSecurityGetRoleRequest configures the X Pack Security Get Role API request.
//
type XPackSecurityGetRoleRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r XPackSecurityGetRoleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_xpack") + 1 + len("security") + 1 + len("role") + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString("_xpack")
	path.WriteString("/")
	path.WriteString("security")
	path.WriteString("/")
	path.WriteString("role")
	if r.Name != "" {
		path.WriteString("/")
		path.WriteString(r.Name)
	}

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f XPackSecurityGetRole) WithContext(v context.Context) func(*XPackSecurityGetRoleRequest) {
	return func(r *XPackSecurityGetRoleRequest) {
		r.ctx = v
	}
}

// WithName - role name.
//
func (f XPackSecurityGetRole) WithName(v string) func(*XPackSecurityGetRoleRequest) {
	return func(r *XPackSecurityGetRoleRequest) {
		r.Name = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f XPackSecurityGetRole) WithPretty() func(*XPackSecurityGetRoleRequest) {
	return func(r *XPackSecurityGetRoleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f XPackSecurityGetRole) WithHuman() func(*XPackSecurityGetRoleRequest) {
	return func(r *XPackSecurityGetRoleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f XPackSecurityGetRole) WithErrorTrace() func(*XPackSecurityGetRoleRequest) {
	return func(r *XPackSecurityGetRoleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f XPackSecurityGetRole) WithFilterPath(v ...string) func(*XPackSecurityGetRoleRequest) {
	return func(r *XPackSecurityGetRoleRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f XPackSecurityGetRole) WithHeader(h map[string]string) func(*XPackSecurityGetRoleRequest) {
	return func(r *XPackSecurityGetRoleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f XPackSecurityGetRole) WithOpaqueID(s string) func(*XPackSecurityGetRoleRequest) {
	return func(r *XPackSecurityGetRoleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
