// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package install

import (
	"errors"
	"runtime"

	"github.com/elastic/beats/v7/x-pack/agent/pkg/artifact"
	"github.com/elastic/beats/v7/x-pack/agent/pkg/artifact/install/snapshot"
	"github.com/elastic/beats/v7/x-pack/agent/pkg/artifact/install/tar"
	"github.com/elastic/beats/v7/x-pack/agent/pkg/artifact/install/zip"
	"github.com/elastic/beats/v7/x-pack/agent/pkg/release"
)

var (
	// ErrConfigNotProvided is returned when provided config is nil
	ErrConfigNotProvided = errors.New("config not provided")
)

// Installer is an interface allowing installation of an artifact
type Installer interface {
	// Install installs an artifact and returns
	// location of the installed program
	// error if something went wrong
	Install(programName, version, installDir string) error
}

// NewInstaller returns a correct installer associated with a
// package type:
// - rpm -> rpm installer
// - deb -> deb installer
// - binary -> zip installer on windows, tar installer on linux and mac
func NewInstaller(config *artifact.Config) (Installer, error) {
	if config == nil {
		return nil, ErrConfigNotProvided
	}

	var i Installer
	var err error

	if runtime.GOOS == "windows" {
		i, err = zip.NewInstaller(config)
	} else {
		i, err = tar.NewInstaller(config)
	}
	if err != nil {
		return nil, err
	}

	if release.Snapshot() {
		return snapshot.NewInstaller(i)
	}

	return i, nil
}
