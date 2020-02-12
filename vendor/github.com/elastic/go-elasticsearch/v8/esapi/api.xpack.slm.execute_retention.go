// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 8.0.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newSlmExecuteRetentionFunc(t Transport) SlmExecuteRetention {
	return func(o ...func(*SlmExecuteRetentionRequest)) (*Response, error) {
		var r = SlmExecuteRetentionRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SlmExecuteRetention -
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-execute-retention.html.
//
type SlmExecuteRetention func(o ...func(*SlmExecuteRetentionRequest)) (*Response, error)

// SlmExecuteRetentionRequest configures the Slm Execute Retention API request.
//
type SlmExecuteRetentionRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SlmExecuteRetentionRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_slm/_execute_retention"))
	path.WriteString("/_slm/_execute_retention")

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
func (f SlmExecuteRetention) WithContext(v context.Context) func(*SlmExecuteRetentionRequest) {
	return func(r *SlmExecuteRetentionRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SlmExecuteRetention) WithPretty() func(*SlmExecuteRetentionRequest) {
	return func(r *SlmExecuteRetentionRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SlmExecuteRetention) WithHuman() func(*SlmExecuteRetentionRequest) {
	return func(r *SlmExecuteRetentionRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SlmExecuteRetention) WithErrorTrace() func(*SlmExecuteRetentionRequest) {
	return func(r *SlmExecuteRetentionRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SlmExecuteRetention) WithFilterPath(v ...string) func(*SlmExecuteRetentionRequest) {
	return func(r *SlmExecuteRetentionRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f SlmExecuteRetention) WithHeader(h map[string]string) func(*SlmExecuteRetentionRequest) {
	return func(r *SlmExecuteRetentionRequest) {
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
func (f SlmExecuteRetention) WithOpaqueID(s string) func(*SlmExecuteRetentionRequest) {
	return func(r *SlmExecuteRetentionRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
