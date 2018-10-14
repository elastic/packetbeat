// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// +build mage

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"

	"github.com/elastic/beats/dev-tools/mage"
)

func init() {
	mage.BeatDescription = "Filebeat sends log files to Logstash or directly to Elasticsearch."
}

// Build builds the Beat binary.
func Build() error {
	return mage.Build(mage.DefaultBuildArgs())
}

// GolangCrossBuild build the Beat binary inside of the golang-builder.
// Do not use directly, use crossBuild instead.
func GolangCrossBuild() error {
	return mage.GolangCrossBuild(mage.DefaultGolangCrossBuildArgs())
}

// BuildGoDaemon builds the go-daemon binary (use crossBuildGoDaemon).
func BuildGoDaemon() error {
	return mage.BuildGoDaemon()
}

// CrossBuild cross-builds the beat for all target platforms.
func CrossBuild() error {
	return mage.CrossBuild()
}

// CrossBuildGoDaemon cross-builds the go-daemon binary using Docker.
func CrossBuildGoDaemon() error {
	return mage.CrossBuildGoDaemon()
}

// Clean cleans all generated files and build artifacts.
func Clean() error {
	return mage.Clean()
}

// Package packages the Beat for distribution.
// Use SNAPSHOT=true to build snapshots.
// Use PLATFORMS to control the target platforms.
// Use BEAT_VERSION_QUALIFIER to control the version qualifier.
func Package() {
	start := time.Now()
	defer func() { fmt.Println("package ran for", time.Since(start)) }()

	mage.UseCommunityBeatPackaging()
	customizePackaging()

	mg.Deps(Fields, Dashboards, Config, prepareModulePackaging)
	mg.Deps(CrossBuild, CrossBuildGoDaemon)
	mg.SerialDeps(mage.Package, TestPackages)
}

// TestPackages tests the generated packages (i.e. file modes, owners, groups).
func TestPackages() error {
	return mage.TestPackages(mage.WithModules(), mage.WithModulesD())
}

// Update updates the generated files (aka make update).
func Update() error {
	return sh.Run("make", "update")
}

// Fields generates a fields.yml and include/fields.go for the Beat.
func Fields() {
	mg.SerialDeps(fieldsYML, mage.GenerateAllInOneFieldsGo)
}

// fieldsYML generates a fields.yml.
func fieldsYML() error {
	return mage.GenerateFieldsYAML("module")
}

// Dashboards collects all the dashboards and generates index patterns.
func Dashboards() error {
	mg.Deps(Fields)
	return mage.KibanaDashboards("module")
}

// Config generates both the short and reference configs.
func Config() {
	mg.Deps(shortConfig, referenceConfig, createDirModulesD)
}

// GoTestUnit executes the Go unit tests.
// Use TEST_COVERAGE=true to enable code coverage profiling.
// Use RACE_DETECTOR=true to enable the race detector.
func GoTestUnit(ctx context.Context) error {
	return mage.GoTest(ctx, mage.DefaultGoTestUnitArgs())
}

// GoTestIntegration executes the Go integration tests.
// Use TEST_COVERAGE=true to enable code coverage profiling.
// Use RACE_DETECTOR=true to enable the race detector.
func GoTestIntegration(ctx context.Context) error {
	return mage.GoTest(ctx, mage.DefaultGoTestIntegrationArgs())
}

// -----------------------------------------------------------------------------
// Customizations specific to Filebeat.
// - Include modules directory in packages (minus _meta and test files).
// - Include modules.d directory in packages.

const (
	dirModuleGenerated   = "build/package/module"
	dirModulesDGenerated = "build/package/modules.d"
)

// customizePackaging modifies the package specs to add the modules and
// modules.d directory.
func customizePackaging() {
	var (
		moduleTarget = "module"
		module       = mage.PackageFile{
			Mode:   0644,
			Source: dirModuleGenerated,
		}
		modulesDTarget = "modules.d"
		modulesD       = mage.PackageFile{
			Mode:   0644,
			Source: dirModulesDGenerated,
			Config: true,
		}
	)

	for _, args := range mage.Packages {
		pkgType := args.Types[0]
		switch pkgType {
		case mage.TarGz, mage.Zip:
			args.Spec.Files[moduleTarget] = module
			args.Spec.Files[modulesDTarget] = modulesD
		case mage.Deb, mage.RPM:
			args.Spec.Files["/usr/share/{{.BeatName}}/"+moduleTarget] = module
			args.Spec.Files["/etc/{{.BeatName}}/"+modulesDTarget] = modulesD
		case mage.DMG:
			args.Spec.Files["/Library/Application Support/{{.BeatVendor}}/{{.BeatName}}"+moduleTarget] = module
			args.Spec.Files["/etc/{{.BeatName}}/"+modulesDTarget] = modulesD
		default:
			panic(errors.Errorf("unhandled package type: %v", pkgType))
		}
	}
}

// prepareModulePackaging copies the module dir to the build dir and excludes
// _meta and test files so that they are not included in packages.
func prepareModulePackaging() error {
	mg.Deps(createDirModulesD)

	err := mage.Clean([]string{
		dirModuleGenerated,
		dirModulesDGenerated,
	})
	if err != nil {
		return err
	}

	for _, copyAction := range []struct {
		src, dst string
	}{
		{mage.OSSBeatDir("module"), dirModuleGenerated},
		{mage.OSSBeatDir("modules.d"), dirModulesDGenerated},
	} {
		err := (&mage.CopyTask{
			Source:  copyAction.src,
			Dest:    copyAction.dst,
			Mode:    0644,
			DirMode: 0755,
			Exclude: []string{
				"/_meta",
				"/test",
				"fields.go",
			},
		}).Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

func shortConfig() error {
	var configParts = []string{
		mage.OSSBeatDir("_meta/common.p1.yml"),
		mage.OSSBeatDir("_meta/common.p2.yml"),
		"{{ elastic_beats_dir }}/libbeat/_meta/config.yml",
	}

	for i, f := range configParts {
		configParts[i] = mage.MustExpand(f)
	}

	configFile := mage.BeatName + ".yml"
	mage.MustFileConcat(configFile, 0640, configParts...)
	mage.MustFindReplace(configFile, regexp.MustCompile("beatname"), mage.BeatName)
	mage.MustFindReplace(configFile, regexp.MustCompile("beat-index-prefix"), mage.BeatIndexPrefix)
	return nil
}

func referenceConfig() error {
	const modulesConfigYml = "build/config.modules.yml"
	err := mage.GenerateModuleReferenceConfig(modulesConfigYml, mage.OSSBeatDir("module"), "module")
	if err != nil {
		return err
	}
	defer os.Remove(modulesConfigYml)

	var configParts = []string{
		mage.OSSBeatDir("_meta/common.reference.p1.yml"),
		modulesConfigYml,
		mage.OSSBeatDir("_meta/common.reference.p2.yml"),
		"{{ elastic_beats_dir }}/libbeat/_meta/config.reference.yml",
	}

	for i, f := range configParts {
		configParts[i] = mage.MustExpand(f)
	}

	configFile := mage.BeatName + ".reference.yml"
	mage.MustFileConcat(configFile, 0640, configParts...)
	mage.MustFindReplace(configFile, regexp.MustCompile("beatname"), mage.BeatName)
	mage.MustFindReplace(configFile, regexp.MustCompile("beat-index-prefix"), mage.BeatIndexPrefix)
	return nil
}

func createDirModulesD() error {
	if err := os.RemoveAll("modules.d"); err != nil {
		return err
	}

	shortConfigs, err := filepath.Glob("module/*/_meta/config.yml")
	if err != nil {
		return err
	}

	for _, f := range shortConfigs {
		parts := strings.Split(filepath.ToSlash(f), "/")
		if len(parts) < 2 {
			continue
		}
		moduleName := parts[1]

		cp := mage.CopyTask{
			Source: f,
			Dest:   filepath.Join("modules.d", moduleName+".yml.disabled"),
			Mode:   0644,
		}
		if err = cp.Execute(); err != nil {
			return err
		}
	}
	return nil
}
