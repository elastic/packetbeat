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

func newXPackMLValidateDetectorFunc(t Transport) XPackMLValidateDetector {
	return func(body io.Reader, o ...func(*XPackMLValidateDetectorRequest)) (*Response, error) {
		var r = XPackMLValidateDetectorRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// XPackMLValidateDetector -
//
type XPackMLValidateDetector func(body io.Reader, o ...func(*XPackMLValidateDetectorRequest)) (*Response, error)

// XPackMLValidateDetectorRequest configures the X PackML Validate Detector API request.
//
type XPackMLValidateDetectorRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r XPackMLValidateDetectorRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_xpack/ml/anomaly_detectors/_validate/detector"))
	path.WriteString("/_xpack/ml/anomaly_detectors/_validate/detector")

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
func (f XPackMLValidateDetector) WithContext(v context.Context) func(*XPackMLValidateDetectorRequest) {
	return func(r *XPackMLValidateDetectorRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f XPackMLValidateDetector) WithPretty() func(*XPackMLValidateDetectorRequest) {
	return func(r *XPackMLValidateDetectorRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f XPackMLValidateDetector) WithHuman() func(*XPackMLValidateDetectorRequest) {
	return func(r *XPackMLValidateDetectorRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f XPackMLValidateDetector) WithErrorTrace() func(*XPackMLValidateDetectorRequest) {
	return func(r *XPackMLValidateDetectorRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f XPackMLValidateDetector) WithFilterPath(v ...string) func(*XPackMLValidateDetectorRequest) {
	return func(r *XPackMLValidateDetectorRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f XPackMLValidateDetector) WithHeader(h map[string]string) func(*XPackMLValidateDetectorRequest) {
	return func(r *XPackMLValidateDetectorRequest) {
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
func (f XPackMLValidateDetector) WithOpaqueID(s string) func(*XPackMLValidateDetectorRequest) {
	return func(r *XPackMLValidateDetectorRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
