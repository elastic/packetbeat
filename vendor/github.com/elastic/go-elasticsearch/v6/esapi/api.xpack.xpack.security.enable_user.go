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

func newXPackSecurityEnableUserFunc(t Transport) XPackSecurityEnableUser {
	return func(username string, o ...func(*XPackSecurityEnableUserRequest)) (*Response, error) {
		var r = XPackSecurityEnableUserRequest{Username: username}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// XPackSecurityEnableUser - https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-enable-user.html
//
type XPackSecurityEnableUser func(username string, o ...func(*XPackSecurityEnableUserRequest)) (*Response, error)

// XPackSecurityEnableUserRequest configures the X Pack Security Enable User API request.
//
type XPackSecurityEnableUserRequest struct {
	Username string

	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r XPackSecurityEnableUserRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_xpack") + 1 + len("security") + 1 + len("user") + 1 + len(r.Username) + 1 + len("_enable"))
	path.WriteString("/")
	path.WriteString("_xpack")
	path.WriteString("/")
	path.WriteString("security")
	path.WriteString("/")
	path.WriteString("user")
	path.WriteString("/")
	path.WriteString(r.Username)
	path.WriteString("/")
	path.WriteString("_enable")

	params = make(map[string]string)

	if r.Refresh != "" {
		params["refresh"] = r.Refresh
	}

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
func (f XPackSecurityEnableUser) WithContext(v context.Context) func(*XPackSecurityEnableUserRequest) {
	return func(r *XPackSecurityEnableUserRequest) {
		r.ctx = v
	}
}

// WithRefresh - if `true` (the default) then refresh the affected shards to make this operation visible to search, if `wait_for` then wait for a refresh to make this operation visible to search, if `false` then do nothing with refreshes..
//
func (f XPackSecurityEnableUser) WithRefresh(v string) func(*XPackSecurityEnableUserRequest) {
	return func(r *XPackSecurityEnableUserRequest) {
		r.Refresh = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f XPackSecurityEnableUser) WithPretty() func(*XPackSecurityEnableUserRequest) {
	return func(r *XPackSecurityEnableUserRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f XPackSecurityEnableUser) WithHuman() func(*XPackSecurityEnableUserRequest) {
	return func(r *XPackSecurityEnableUserRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f XPackSecurityEnableUser) WithErrorTrace() func(*XPackSecurityEnableUserRequest) {
	return func(r *XPackSecurityEnableUserRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f XPackSecurityEnableUser) WithFilterPath(v ...string) func(*XPackSecurityEnableUserRequest) {
	return func(r *XPackSecurityEnableUserRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f XPackSecurityEnableUser) WithHeader(h map[string]string) func(*XPackSecurityEnableUserRequest) {
	return func(r *XPackSecurityEnableUserRequest) {
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
func (f XPackSecurityEnableUser) WithOpaqueID(s string) func(*XPackSecurityEnableUserRequest) {
	return func(r *XPackSecurityEnableUserRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
