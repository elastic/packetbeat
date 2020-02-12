// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 6.8.5: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newXPackSQLQueryFunc(t Transport) XPackSQLQuery {
	return func(body io.Reader, o ...func(*XPackSQLQueryRequest)) (*Response, error) {
		var r = XPackSQLQueryRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// XPackSQLQuery - Execute SQL
//
type XPackSQLQuery func(body io.Reader, o ...func(*XPackSQLQueryRequest)) (*Response, error)

// XPackSQLQueryRequest configures the X PackSQL Query API request.
//
type XPackSQLQueryRequest struct {
	Body io.Reader

	Format string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r XPackSQLQueryRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_xpack/sql"))
	path.WriteString("/_xpack/sql")

	params = make(map[string]string)

	if r.Format != "" {
		params["format"] = r.Format
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

	req, err := newRequest(method, path.String(), r.Body)
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

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
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
func (f XPackSQLQuery) WithContext(v context.Context) func(*XPackSQLQueryRequest) {
	return func(r *XPackSQLQueryRequest) {
		r.ctx = v
	}
}

// WithFormat - a short version of the accept header, e.g. json, yaml.
//
func (f XPackSQLQuery) WithFormat(v string) func(*XPackSQLQueryRequest) {
	return func(r *XPackSQLQueryRequest) {
		r.Format = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f XPackSQLQuery) WithPretty() func(*XPackSQLQueryRequest) {
	return func(r *XPackSQLQueryRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f XPackSQLQuery) WithHuman() func(*XPackSQLQueryRequest) {
	return func(r *XPackSQLQueryRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f XPackSQLQuery) WithErrorTrace() func(*XPackSQLQueryRequest) {
	return func(r *XPackSQLQueryRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f XPackSQLQuery) WithFilterPath(v ...string) func(*XPackSQLQueryRequest) {
	return func(r *XPackSQLQueryRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f XPackSQLQuery) WithHeader(h map[string]string) func(*XPackSQLQueryRequest) {
	return func(r *XPackSQLQueryRequest) {
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
func (f XPackSQLQuery) WithOpaqueID(s string) func(*XPackSQLQueryRequest) {
	return func(r *XPackSQLQueryRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
