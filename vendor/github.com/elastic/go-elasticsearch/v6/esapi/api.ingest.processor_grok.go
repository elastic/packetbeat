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

func newIngestProcessorGrokFunc(t Transport) IngestProcessorGrok {
	return func(o ...func(*IngestProcessorGrokRequest)) (*Response, error) {
		var r = IngestProcessorGrokRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// IngestProcessorGrok returns a list of the built-in patterns.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/grok-processor.html#grok-processor-rest-get.
//
type IngestProcessorGrok func(o ...func(*IngestProcessorGrokRequest)) (*Response, error)

// IngestProcessorGrokRequest configures the Ingest Processor Grok API request.
//
type IngestProcessorGrokRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IngestProcessorGrokRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_ingest/processor/grok"))
	path.WriteString("/_ingest/processor/grok")

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
func (f IngestProcessorGrok) WithContext(v context.Context) func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IngestProcessorGrok) WithPretty() func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IngestProcessorGrok) WithHuman() func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IngestProcessorGrok) WithErrorTrace() func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IngestProcessorGrok) WithFilterPath(v ...string) func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f IngestProcessorGrok) WithHeader(h map[string]string) func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
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
func (f IngestProcessorGrok) WithOpaqueID(s string) func(*IngestProcessorGrokRequest) {
	return func(r *IngestProcessorGrokRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
