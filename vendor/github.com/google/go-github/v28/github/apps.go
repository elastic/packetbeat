// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
	"time"
)

// AppsService provides access to the installation related functions
// in the GitHub API.
//
// GitHub API docs: https://developer.github.com/v3/apps/
type AppsService service

// App represents a GitHub App.
type App struct {
	ID          *int64     `json:"id,omitempty"`
	NodeID      *string    `json:"node_id,omitempty"`
	Owner       *User      `json:"owner,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	ExternalURL *string    `json:"external_url,omitempty"`
	HTMLURL     *string    `json:"html_url,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
}

// InstallationToken represents an installation token.
type InstallationToken struct {
	Token        *string                  `json:"token,omitempty"`
	ExpiresAt    *time.Time               `json:"expires_at,omitempty"`
	Permissions  *InstallationPermissions `json:"permissions,omitempty"`
	Repositories []*Repository            `json:"repositories,omitempty"`
}

// InstallationTokenOptions allow restricting a token's access to specific repositories.
type InstallationTokenOptions struct {
	// The IDs of the repositories that the installation token can access.
	// Providing repository IDs restricts the access of an installation token to specific repositories.
	RepositoryIDs []int64 `json:"repository_ids,omitempty"`

	// The permissions granted to the access token.
	// The permissions object includes the permission names and their access type.
	Permissions *InstallationPermissions `json:"permissions,omitempty"`
}

// InstallationPermissions lists the repository and organization permissions for an installation.
//
// Permission names taken from:
//   https://developer.github.com/v3/apps/permissions/
//   https://developer.github.com/enterprise/v3/apps/permissions/
type InstallationPermissions struct {
	Administration              *string `json:"administration,omitempty"`
	Checks                      *string `json:"checks,omitempty"`
	Contents                    *string `json:"contents,omitempty"`
	ContentReferences           *string `json:"content_references,omitempty"`
	Deployments                 *string `json:"deployments,omitempty"`
	Issues                      *string `json:"issues,omitempty"`
	Metadata                    *string `json:"metadata,omitempty"`
	Members                     *string `json:"members,omitempty"`
	OrganizationAdministration  *string `json:"organization_administration,omitempty"`
	OrganizationHooks           *string `json:"organization_hooks,omitempty"`
	OrganizationPlan            *string `json:"organization_plan,omitempty"`
	OrganizationPreReceiveHooks *string `json:"organization_pre_receive_hooks,omitempty"`
	OrganizationProjects        *string `json:"organization_projects,omitempty"`
	OrganizationUserBlocking    *string `json:"organization_user_blocking,omitempty"`
	Packages                    *string `json:"packages,omitempty"`
	Pages                       *string `json:"pages,omitempty"`
	PullRequests                *string `json:"pull_requests,omitempty"`
	RepositoryHooks             *string `json:"repository_hooks,omitempty"`
	RepositoryProjects          *string `json:"repository_projects,omitempty"`
	RepositoryPreReceiveHooks   *string `json:"repository_pre_receive_hooks,omitempty"`
	SingleFile                  *string `json:"single_file,omitempty"`
	Statuses                    *string `json:"statuses,omitempty"`
	TeamDiscussions             *string `json:"team_discussions,omitempty"`
	VulnerabilityAlerts         *string `json:"vulnerability_alerts,omitempty"`
}

// Installation represents a GitHub Apps installation.
type Installation struct {
	ID                  *int64                   `json:"id,omitempty"`
	AppID               *int64                   `json:"app_id,omitempty"`
	TargetID            *int64                   `json:"target_id,omitempty"`
	Account             *User                    `json:"account,omitempty"`
	AccessTokensURL     *string                  `json:"access_tokens_url,omitempty"`
	RepositoriesURL     *string                  `json:"repositories_url,omitempty"`
	HTMLURL             *string                  `json:"html_url,omitempty"`
	TargetType          *string                  `json:"target_type,omitempty"`
	SingleFileName      *string                  `json:"single_file_name,omitempty"`
	RepositorySelection *string                  `json:"repository_selection,omitempty"`
	Events              []string                 `json:"events,omitempty"`
	Permissions         *InstallationPermissions `json:"permissions,omitempty"`
	CreatedAt           *Timestamp               `json:"created_at,omitempty"`
	UpdatedAt           *Timestamp               `json:"updated_at,omitempty"`
}

// Attachment represents a GitHub Apps attachment.
type Attachment struct {
	ID    *int64  `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Body  *string `json:"body,omitempty"`
}

func (i Installation) String() string {
	return Stringify(i)
}

// Get a single GitHub App. Passing the empty string will get
// the authenticated GitHub App.
//
// Note: appSlug is just the URL-friendly name of your GitHub App.
// You can find this on the settings page for your GitHub App
// (e.g., https://github.com/settings/apps/:app_slug).
//
// GitHub API docs: https://developer.github.com/v3/apps/#get-a-single-github-app
func (s *AppsService) Get(ctx context.Context, appSlug string) (*App, *Response, error) {
	var u string
	if appSlug != "" {
		u = fmt.Sprintf("apps/%v", appSlug)
	} else {
		u = "app"
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launches.
	req.Header.Set("Accept", mediaTypeIntegrationPreview)

	app := new(App)
	resp, err := s.client.Do(ctx, req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, nil
}

// ListInstallations lists the installations that the current GitHub App has.
//
// GitHub API docs: https://developer.github.com/v3/apps/#find-installations
func (s *AppsService) ListInstallations(ctx context.Context, opt *ListOptions) ([]*Installation, *Response, error) {
	u, err := addOptions("app/installations", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launches.
	req.Header.Set("Accept", mediaTypeIntegrationPreview)

	var i []*Installation
	resp, err := s.client.Do(ctx, req, &i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// GetInstallation returns the specified installation.
//
// GitHub API docs: https://developer.github.com/v3/apps/#get-a-single-installation
func (s *AppsService) GetInstallation(ctx context.Context, id int64) (*Installation, *Response, error) {
	return s.getInstallation(ctx, fmt.Sprintf("app/installations/%v", id))
}

// ListUserInstallations lists installations that are accessible to the authenticated user.
//
// GitHub API docs: https://developer.github.com/v3/apps/#list-installations-for-user
func (s *AppsService) ListUserInstallations(ctx context.Context, opt *ListOptions) ([]*Installation, *Response, error) {
	u, err := addOptions("user/installations", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launches.
	req.Header.Set("Accept", mediaTypeIntegrationPreview)

	var i struct {
		Installations []*Installation `json:"installations"`
	}
	resp, err := s.client.Do(ctx, req, &i)
	if err != nil {
		return nil, resp, err
	}

	return i.Installations, resp, nil
}

// CreateInstallationToken creates a new installation token.
//
// GitHub API docs: https://developer.github.com/v3/apps/#create-a-new-installation-token
func (s *AppsService) CreateInstallationToken(ctx context.Context, id int64, opt *InstallationTokenOptions) (*InstallationToken, *Response, error) {
	u := fmt.Sprintf("app/installations/%v/access_tokens", id)

	req, err := s.client.NewRequest("POST", u, opt)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launches.
	req.Header.Set("Accept", mediaTypeIntegrationPreview)

	t := new(InstallationToken)
	resp, err := s.client.Do(ctx, req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil
}

// Create a new attachment on user comment containing a url.
//
// GitHub API docs: https://developer.github.com/v3/apps/#create-a-content-attachment
func (s *AppsService) CreateAttachment(ctx context.Context, contentReferenceID int64, title, body string) (*Attachment, *Response, error) {
	u := fmt.Sprintf("content_references/%v/attachments", contentReferenceID)
	payload := &Attachment{Title: String(title), Body: String(body)}
	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept headers when APIs fully launch.
	req.Header.Set("Accept", mediaTypeReactionsPreview)

	m := &Attachment{}
	resp, err := s.client.Do(ctx, req, m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}

// FindOrganizationInstallation finds the organization's installation information.
//
// GitHub API docs: https://developer.github.com/v3/apps/#find-organization-installation
func (s *AppsService) FindOrganizationInstallation(ctx context.Context, org string) (*Installation, *Response, error) {
	return s.getInstallation(ctx, fmt.Sprintf("orgs/%v/installation", org))
}

// FindRepositoryInstallation finds the repository's installation information.
//
// GitHub API docs: https://developer.github.com/v3/apps/#find-repository-installation
func (s *AppsService) FindRepositoryInstallation(ctx context.Context, owner, repo string) (*Installation, *Response, error) {
	return s.getInstallation(ctx, fmt.Sprintf("repos/%v/%v/installation", owner, repo))
}

// FindRepositoryInstallationByID finds the repository's installation information.
//
// Note: FindRepositoryInstallationByID uses the undocumented GitHub API endpoint /repositories/:id/installation.
func (s *AppsService) FindRepositoryInstallationByID(ctx context.Context, id int64) (*Installation, *Response, error) {
	return s.getInstallation(ctx, fmt.Sprintf("repositories/%d/installation", id))
}

// FindUserInstallation finds the user's installation information.
//
// GitHub API docs: https://developer.github.com/v3/apps/#find-repository-installation
func (s *AppsService) FindUserInstallation(ctx context.Context, user string) (*Installation, *Response, error) {
	return s.getInstallation(ctx, fmt.Sprintf("users/%v/installation", user))
}

func (s *AppsService) getInstallation(ctx context.Context, url string) (*Installation, *Response, error) {
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launches.
	req.Header.Set("Accept", mediaTypeIntegrationPreview)

	i := new(Installation)
	resp, err := s.client.Do(ctx, req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}
