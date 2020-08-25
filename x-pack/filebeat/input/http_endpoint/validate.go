// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package http_endpoint

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type validator interface {
	Validate(*http.Request) (int, error)
}

type authValidator struct {
	basicAuth          bool
	username, password string
	secretHeader       string
	secretValue        string
}

type headerValidator struct {
	method      string
	contentType string
}

type hmacValidator struct {
	hmacHeader string
	hmacToken  string
	hmacPrefix string
}

var errIncorrectUserOrPass = errors.New("Incorrect username or password")
var errIncorrectHeaderSecret = errors.New("Incorrect header or header secret")
var errIncorrectHmac = errors.New("The HMAC signature of the request body does not match with the configured secret")

func (v *authValidator) Validate(r *http.Request) (int, error) {
	if v.basicAuth {
		username, password, _ := r.BasicAuth()
		if v.username != username || v.password != password {
			return http.StatusUnauthorized, errIncorrectUserOrPass
		}
	}
	if v.secretHeader != "" && v.secretValue != "" {
		if v.secretValue != r.Header.Get(v.secretHeader) {
			return http.StatusUnauthorized, errIncorrectHeaderSecret
		}
	}

	return 0, nil
}

func (v *headerValidator) Validate(r *http.Request) (int, error) {
	if v.method != "" && v.method != r.Method {
		return http.StatusMethodNotAllowed, fmt.Errorf("Only %v requests supported", v.method)
	}

	if v.contentType != "" && r.Header.Get("Content-Type") != v.contentType {
		return http.StatusUnsupportedMediaType, fmt.Errorf("Wrong Content-Type header, expecting %v", v.contentType)
	}

	return 0, nil
}

func (v *hmacValidator) Validate(r *http.Request) (int, error) {
	if v.hmacHeader != "" && v.hmacToken == "" || v.hmacHeader == "" && v.hmacToken != "" {
		return http.StatusMethodNotAllowed, fmt.Errorf("Both hmacToken and hmacHeader has to be set")
	}

	if v.hmacToken != "" && v.hmacHeader != "" {
		if len(r.Header.Get(v.hmacHeader)) == 0 {
			return http.StatusInternalServerError, fmt.Errorf("The HMAC signature in the configured request header is empty")
		}
		s := hmac.New(sha1.New, []byte(v.hmacToken))
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return http.StatusInternalServerError, fmt.Errorf("Failed to read the request body: %v", err)
		}
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

		s.Write(b)
		h := r.Header.Get(v.hmacHeader)
		// If the header includes a prefix before the SHA-1 key, we need to only grab the signature after the prefix. Can also be 0
		hWithPrefix := make([]byte, 20)
		hex.Decode(hWithPrefix, []byte(h[len(v.hmacPrefix):]))

		if !hmac.Equal(s.Sum(nil), hWithPrefix) {
			return http.StatusUnauthorized, errIncorrectHmac
		}

	}

	return 0, nil
}
