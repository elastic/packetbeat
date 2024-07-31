// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Package okta provides Okta API support.
package okta

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"

	"github.com/elastic/elastic-agent-libs/logp"
)

// ISO8601 is the time format accepted by Okta queries.
const ISO8601 = "2006-01-02T15:04:05.000Z"

// User is an Okta user's details.
//
// See https://developer.okta.com/docs/reference/api/users/#user-properties for details.
type User struct {
	ID                    string         `json:"id"`
	Status                string         `json:"status"`
	Created               time.Time      `json:"created"`
	Activated             time.Time      `json:"activated"`
	StatusChanged         *time.Time     `json:"statusChanged,omitempty"`
	LastLogin             *time.Time     `json:"lastLogin,omitempty"`
	LastUpdated           time.Time      `json:"lastUpdated"`
	PasswordChanged       *time.Time     `json:"passwordChanged,omitempty"`
	Type                  map[string]any `json:"type"`
	TransitioningToStatus *string        `json:"transitioningToStatus,omitempty"`
	Profile               map[string]any `json:"profile"`
	Credentials           *Credentials   `json:"credentials,omitempty"`
	Links                 HAL            `json:"_links,omitempty"` // See https://developer.okta.com/docs/reference/api/users/#links-object for details.
	Embedded              HAL            `json:"_embedded,omitempty"`
}

// Credentials is a redacted Okta user's credential details. Only the credential provider is retained.
//
// See https://developer.okta.com/docs/reference/api/users/#credentials-object for details.
type Credentials struct {
	Password         *struct{} `json:"password,omitempty"` // Contains "value"; omit but mark.
	RecoveryQuestion *struct {
		// Also contains "answer""; intentionally omitted.
		Question string `json:"question"`
	} `json:"recovery_question,omitempty"`
	Provider Provider `json:"provider"`
}

// Provider is an Okta credential provider.
//
// See https://developer.okta.com/docs/reference/api/users/#provider-object for details.
type Provider struct {
	Type string  `json:"type"`
	Name *string `json:"name,omitempty"`
}

// Group is an Okta user group.
//
// See https://developer.okta.com/docs/reference/api/users/#request-parameters-8 (no anchor exists on the page for this endpoint) for details.
type Group struct {
	ID      string         `json:"id"`
	Profile map[string]any `json:"profile"`
}

// Device is an Okta device's details.
//
// See https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Device/#tag/Device/operation/listDevices for details
type Device struct {
	Created             time.Time         `json:"created"`
	ID                  string            `json:"id"`
	LastUpdated         time.Time         `json:"lastUpdated"`
	Profile             map[string]any    `json:"profile"`
	ResourceAlternateID string            `json:"resourceAlternateID"`
	ResourceDisplayName DeviceDisplayName `json:"resourceDisplayName"`
	ResourceID          string            `json:"resourceID"`
	ResourceType        string            `json:"resourceType"`
	Status              string            `json:"status"`
	Links               HAL               `json:"_links,omitempty"` // See https://developer.okta.com/docs/reference/api/users/#links-object for details.

	// Users is the set of users associated with the device.
	// It is not part of the list devices API return, but can
	// be populated by a call to GetDeviceUsers.
	Users []User `json:"users,omitempty"`
}

// DeviceDisplayName is an Okta device's annotated display name.
//
// See https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Device/#tag/Device/operation/listDevices for details
type DeviceDisplayName struct {
	Sensitive bool   `json:"sensitive"`
	Value     string `json:"value"`
}

// HAL is a JSON Hypertext Application Language object.
//
// See https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06 for details.
type HAL map[string]any

// Response is a set of omit options specifying a part of the response to omit.
//
// See https://developer.okta.com/docs/reference/api/users/#content-type-header-fields-2 for details.
type Response uint8

const (
	// Omit the credentials sub-object from the response.
	OmitCredentials Response = 1 << iota

	// Omit the following HAL links from the response:
	// Change Password, Change Recovery Question, Forgot Password, Reset Password, Reset Factors, Unlock.
	OmitCredentialsLinks

	// Omit the transitioningToStatus field from the response.
	OmitTransitioningToStatus

	OmitNone Response = 0
)

var oktaResponse = [...]string{
	"omitCredentials",
	"omitCredentialsLinks",
	"omitTransitioningToStatus",
}

func (o Response) String() string {
	if o == OmitNone {
		return ""
	}
	var buf strings.Builder
	buf.WriteString("okta-response=")
	var n int
	for i, s := range &oktaResponse {
		if o&(1<<i) != 0 {
			if n != 0 {
				buf.WriteByte(',')
			}
			buf.WriteString(s)
			n++
		}
	}
	return buf.String()
}

// GetUserDetails returns Okta user details using the list users API endpoint. host is the
// Okta user domain and key is the API token to use for the query. If user is not empty,
// details for the specific user are returned, otherwise a list of all users is returned.
// The query parameter holds queries as described in https://developer.okta.com/docs/reference/user-query/
// with the query syntax described at https://developer.okta.com/docs/reference/core-okta-api/#filter.
// Parts of the response may be omitted using the omit parameter.
//
// The provided rate limiter must allow at least request and will be updated with the
// response's X-Rate-Limit headers considering the rate limit window time. Details
// for rate limits are available at https://help.okta.com/en-us/Content/Topics/Security/API-rate-limits.htm
// and account rate limits and windows can be seen on the Okta admin dashboard at
// https://${yourOktaDomain}/reports/rate-limit.
//
// See https://developer.okta.com/docs/reference/api/users/#list-users for details.
func GetUserDetails(ctx context.Context, cli *http.Client, host, key, user string, query url.Values, omit Response, lim *rate.Limiter, window time.Duration, log *logp.Logger) ([]User, http.Header, error) {
	const endpoint = "/api/v1/users"

	u := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     path.Join(endpoint, user),
		RawQuery: query.Encode(),
	}
	return getDetails[User](ctx, cli, u, key, user == "", omit, lim, window, log)
}

// GetUserGroupDetails returns Okta group details using the users API endpoint. host is the
// Okta user domain and key is the API token to use for the query. user must not be empty.
//
// See GetUserDetails for details of the query and rate limit parameters.
//
// See https://developer.okta.com/docs/reference/api/users/#request-parameters-8 (no anchor exists on the page for this endpoint) for details.
func GetUserGroupDetails(ctx context.Context, cli *http.Client, host, key, user string, lim *rate.Limiter, window time.Duration, log *logp.Logger) ([]Group, http.Header, error) {
	const endpoint = "/api/v1/users"

	if user == "" {
		return nil, nil, errors.New("no user specified")
	}

	u := &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path.Join(endpoint, user, "groups"),
	}
	return getDetails[Group](ctx, cli, u, key, true, OmitNone, lim, window, log)
}

// GetDeviceDetails returns Okta device details using the list devices API endpoint. host is the
// Okta user domain and key is the API token to use for the query. If device is not empty,
// details for the specific device are returned, otherwise a list of all devices is returned.
//
// See GetUserDetails for details of the query and rate limit parameters.
//
// See https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Device/#tag/Device/operation/listDevices for details.
func GetDeviceDetails(ctx context.Context, cli *http.Client, host, key, device string, query url.Values, lim *rate.Limiter, window time.Duration, log *logp.Logger) ([]Device, http.Header, error) {
	const endpoint = "/api/v1/devices"

	u := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     path.Join(endpoint, device),
		RawQuery: query.Encode(),
	}
	return getDetails[Device](ctx, cli, u, key, device == "", OmitNone, lim, window, log)
}

// GetDeviceUsers returns Okta user details for users associated with the provided device identifier
// using the list device users API. host is the Okta user domain and key is the API token to use for
// the query. If device is empty, a nil User slice and header is returned, without error.
//
// See GetUserDetails for details of the query and rate limit parameters.
//
// See https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Device/#tag/Device/operation/listDeviceUsers for details.
func GetDeviceUsers(ctx context.Context, cli *http.Client, host, key, device string, query url.Values, omit Response, lim *rate.Limiter, window time.Duration, log *logp.Logger) ([]User, http.Header, error) {
	if device == "" {
		// No user associated with a null device. Not an error.
		return nil, nil, nil
	}

	const endpoint = "/api/v1/devices"

	u := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     path.Join(endpoint, device, "users"),
		RawQuery: query.Encode(),
	}
	du, h, err := getDetails[devUser](ctx, cli, u, key, true, omit, lim, window, log)
	if err != nil {
		return nil, h, err
	}
	users := make([]User, len(du))
	for i, du := range du {
		users[i] = du.User
	}
	return users, h, nil
}

// entity is an Okta entity analytics entity.
type entity interface {
	User | Group | Device | devUser
}

type devUser struct {
	User `json:"user"`
}

// getDetails returns Okta details using the API endpoint in u. host is the Okta
// user domain and key is the API token to use for the query. If all is false, details
// for the specific user are returned, otherwise a list of all users is returned.
//
// See GetUserDetails for details of the query and rate limit parameters.
func getDetails[E entity](ctx context.Context, cli *http.Client, u *url.URL, key string, all bool, omit Response, lim *rate.Limiter, window time.Duration, log *logp.Logger) ([]E, http.Header, error) {
	url := u.String()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", "application/json")
	contentType := "application/json"
	if omit != OmitNone {
		contentType += "; " + omit.String()
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", fmt.Sprintf("SSWS %s", key))

	log.Debugw("rate limit", "limit", lim.Limit(), "burst", lim.Burst(), "url", url)
	err = lim.Wait(ctx)
	if err != nil {
		return nil, nil, err
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	err = oktaRateLimit(resp.Header, window, lim, log)
	if err != nil {
		io.Copy(io.Discard, resp.Body)
		return nil, nil, err
	}

	var body bytes.Buffer
	n, err := io.Copy(&body, resp.Body)
	if n == 0 || err != nil {
		return nil, nil, err
	}

	if all {
		// List all entities.
		var e []E
		err = json.Unmarshal(body.Bytes(), &e)
		if err != nil {
			err = recoverError(body.Bytes())
		}
		return e, resp.Header, err
	}
	// Get single entity's details.
	var e [1]E
	err = json.Unmarshal(body.Bytes(), &e[0])
	if err != nil {
		err = recoverError(body.Bytes())
	}
	return e[:], resp.Header, err
}

// recoverError returns an error based on the returned Okta API error. Error
// detection here depends on Okta errors being a JSON object while we are
// requesting a JSON array.
func recoverError(msg []byte) error {
	var e Error
	err := json.Unmarshal(msg, &e)
	if err != nil {
		return err
	}
	return &e
}

// Error is an Okta API error value.
type Error struct {
	Code    string  `json:"errorCode,omitempty"`
	Summary string  `json:"errorSummary,omitempty"`
	Link    string  `json:"errorLink,omitempty"`
	ID      string  `json:"errorId,omitempty"`
	Causes  []Error `json:"errorCauses,omitempty"`
}

func (e *Error) Error() string {
	summary := strings.ToLower(strings.TrimRight(e.Summary, "."))
	if len(e.Causes) == 0 {
		return summary
	}
	causes := make([]string, len(e.Causes))
	for i, c := range e.Causes {
		causes[i] = c.Error()
	}
	return fmt.Sprintf("%s: %s", summary, strings.Join(causes, ","))
}

// oktaRateLimit implements the Okta rate limit policy translation.
//
// See https://developer.okta.com/docs/reference/rl-best-practices/ for details.
func oktaRateLimit(h http.Header, window time.Duration, limiter *rate.Limiter, log *logp.Logger) error {
	limit := h.Get("X-Rate-Limit-Limit")
	remaining := h.Get("X-Rate-Limit-Remaining")
	reset := h.Get("X-Rate-Limit-Reset")
	log.Debugw("rate limit header", "X-Rate-Limit-Limit", limit, "X-Rate-Limit-Remaining", remaining, "X-Rate-Limit-Reset", reset)
	if limit == "" || remaining == "" || reset == "" {
		return nil
	}

	lim, err := strconv.ParseFloat(limit, 64)
	if err != nil {
		return err
	}
	rem, err := strconv.ParseFloat(remaining, 64)
	if err != nil {
		return err
	}
	rst, err := strconv.ParseInt(reset, 10, 64)
	if err != nil {
		return err
	}
	resetTime := time.Unix(rst, 0)
	per := time.Until(resetTime).Seconds()

	// Be conservative here; the docs don't exactly specify burst rates.
	// Make sure we can make at least one new request, even if we fail
	// to get a non-zero rate.Limit. We could set to zero for the case
	// that limit=rate.Inf, but that detail is not important.
	burst := 1

	rateLimit := rate.Limit(rem / per)

	// Process reset if we need to wait until reset to avoid a request against a zero quota.
	if rateLimit == 0 {
		waitUntil := resetTime.UTC()
		// next gives us a sane next window estimate, but the
		// estimate will be overwritten when we make the next
		// permissible API request.
		next := rate.Limit(lim / window.Seconds())
		limiter.SetLimitAt(waitUntil, next)
		limiter.SetBurstAt(waitUntil, burst)
		log.Debugw("rate limit adjust", "reset_time", waitUntil, "next_rate", next, "next_burst", burst)
		return nil
	}
	limiter.SetLimit(rateLimit)
	limiter.SetBurst(burst)
	log.Debugw("rate limit adjust", "set_rate", rateLimit, "set_burst", burst)
	return nil
}

// Next returns the next URL query for a pagination sequence. If no further
// page is available, Next returns io.EOF.
func Next(h http.Header) (query url.Values, err error) {
	for _, v := range h.Values("link") {
		f := strings.Split(v, ";")
		if len(f) == 1 {
			continue
		}
		for _, p := range f[1:] {
			_, rel, ok := strings.Cut(p, "rel")
			if !ok {
				continue
			}
			_, rel, ok = strings.Cut(rel, "=")
			if !ok {
				continue
			}
			if strings.HasPrefix(strings.TrimSpace(rel), `"next"`) {
				u, err := url.Parse(strings.TrimFunc(f[0], func(r rune) bool { return r == '<' || r == '>' }))
				if err != nil {
					return nil, err
				}
				return u.Query(), nil
			}
		}
	}
	return nil, io.EOF
}
