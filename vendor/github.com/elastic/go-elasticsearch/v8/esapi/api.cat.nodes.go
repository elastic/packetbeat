// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 8.0.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newCatNodesFunc(t Transport) CatNodes {
	return func(o ...func(*CatNodesRequest)) (*Response, error) {
		var r = CatNodesRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// CatNodes returns basic statistics about performance of cluster nodes.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/cat-nodes.html.
//
type CatNodes func(o ...func(*CatNodesRequest)) (*Response, error)

// CatNodesRequest configures the Cat Nodes API request.
//
type CatNodesRequest struct {
	Bytes         string
	Format        string
	FullID        *bool
	H             []string
	Help          *bool
	MasterTimeout time.Duration
	S             []string
	Time          string
	V             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r CatNodesRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_cat/nodes"))
	path.WriteString("/_cat/nodes")

	params = make(map[string]string)

	if r.Bytes != "" {
		params["bytes"] = r.Bytes
	}

	if r.Format != "" {
		params["format"] = r.Format
	}

	if r.FullID != nil {
		params["full_id"] = strconv.FormatBool(*r.FullID)
	}

	if len(r.H) > 0 {
		params["h"] = strings.Join(r.H, ",")
	}

	if r.Help != nil {
		params["help"] = strconv.FormatBool(*r.Help)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if len(r.S) > 0 {
		params["s"] = strings.Join(r.S, ",")
	}

	if r.Time != "" {
		params["time"] = r.Time
	}

	if r.V != nil {
		params["v"] = strconv.FormatBool(*r.V)
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
func (f CatNodes) WithContext(v context.Context) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.ctx = v
	}
}

// WithBytes - the unit in which to display byte values.
//
func (f CatNodes) WithBytes(v string) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.Bytes = v
	}
}

// WithFormat - a short version of the accept header, e.g. json, yaml.
//
func (f CatNodes) WithFormat(v string) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.Format = v
	}
}

// WithFullID - return the full node ID instead of the shortened version (default: false).
//
func (f CatNodes) WithFullID(v bool) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.FullID = &v
	}
}

// WithH - comma-separated list of column names to display.
//
func (f CatNodes) WithH(v ...string) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.H = v
	}
}

// WithHelp - return help information.
//
func (f CatNodes) WithHelp(v bool) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.Help = &v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
//
func (f CatNodes) WithMasterTimeout(v time.Duration) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.MasterTimeout = v
	}
}

// WithS - comma-separated list of column names or column aliases to sort by.
//
func (f CatNodes) WithS(v ...string) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.S = v
	}
}

// WithTime - the unit in which to display time values.
//
func (f CatNodes) WithTime(v string) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.Time = v
	}
}

// WithV - verbose mode. display column headers.
//
func (f CatNodes) WithV(v bool) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.V = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f CatNodes) WithPretty() func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f CatNodes) WithHuman() func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f CatNodes) WithErrorTrace() func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f CatNodes) WithFilterPath(v ...string) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f CatNodes) WithHeader(h map[string]string) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
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
func (f CatNodes) WithOpaqueID(s string) func(*CatNodesRequest) {
	return func(r *CatNodesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
