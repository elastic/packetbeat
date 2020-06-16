// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package operation

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/core/server"

	operatorCfg "github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/agent/operation/config"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/agent/stateresolver"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/artifact"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/artifact/download"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/artifact/install"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/config"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/core/logger"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/core/plugin/app"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/core/plugin/app/monitoring/noop"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/core/plugin/process"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/core/plugin/retry"
)

var downloadPath = getAbsPath("tests/downloads")
var installPath = getAbsPath("tests/scripts")

func getTestOperator(t *testing.T, downloadPath string, installPath string, p *app.Descriptor) (*Operator, *operatorCfg.Config) {
	operatorConfig := &operatorCfg.Config{
		RetryConfig: &retry.Config{
			Enabled:      true,
			RetriesCount: 2,
			Delay:        3 * time.Second,
			MaxDelay:     10 * time.Second,
		},
		ProcessConfig: &process.Config{},
		DownloadConfig: &artifact.Config{
			TargetDirectory: downloadPath,
			InstallPath:     installPath,
		},
	}

	cfg, err := config.NewConfigFrom(operatorConfig)
	if err != nil {
		t.Fatal(err)
	}

	l := getLogger()

	fetcher := &DummyDownloader{}
	verifier := &DummyVerifier{}
	installer := &DummyInstaller{}

	stateResolver, err := stateresolver.NewStateResolver(l)
	if err != nil {
		t.Fatal(err)
	}
	srv, err := server.New(l, ":0", &app.ApplicationStatusHandler{})
	if err != nil {
		t.Fatal(err)
	}
	err = srv.Start()
	if err != nil {
		t.Fatal(err)
	}

	operator, err := NewOperator(context.Background(), l, "p1", cfg, fetcher, verifier, installer, stateResolver, srv, nil, noop.NewMonitor())
	if err != nil {
		t.Fatal(err)
	}

	operator.config.DownloadConfig.OperatingSystem = "darwin"
	operator.config.DownloadConfig.Architecture = "32"

	// make the download path so the `operation_verify` can ensure the path exists
	downloadConfig := operator.config.DownloadConfig
	fullPath, err := artifact.GetArtifactPath(p.BinaryName(), p.Version(), downloadConfig.OS(), downloadConfig.Arch(), downloadConfig.TargetDirectory)
	if err != nil {
		t.Fatal(err)
	}
	createFile(t, fullPath)

	return operator, operatorConfig
}

func getLogger() *logger.Logger {
	cfg, _ := config.NewConfigFrom(map[string]interface{}{
		"logging": map[string]interface{}{
			"level": "error",
		},
	})
	l, _ := logger.NewFromConfig("", cfg)
	return l
}

func getProgram(binary, version string) *app.Descriptor {
	downloadCfg := &artifact.Config{
		InstallPath:     installPath,
		OperatingSystem: "darwin",
		Architecture:    "32",
	}
	return app.NewDescriptor(binary, version, downloadCfg, nil)
}

func getAbsPath(path string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), path)
}

func createFile(t *testing.T, path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()
	}
}

func waitFor(t *testing.T, check func() error) {
	started := time.Now()
	for {
		err := check()
		if err == nil {
			return
		}
		if time.Now().Sub(started) >= 15*time.Second {
			t.Fatalf("check timed out after 15 second: %s", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

type DummyDownloader struct{}

func (*DummyDownloader) Download(_ context.Context, p, v string) (string, error) {
	return "", nil
}

var _ download.Downloader = &DummyDownloader{}

type DummyVerifier struct{}

func (*DummyVerifier) Verify(p, v string) (bool, error) {
	return true, nil
}

var _ download.Verifier = &DummyVerifier{}

type DummyInstaller struct{}

func (*DummyInstaller) Install(p, v, _ string) error {
	return nil
}

var _ install.Installer = &DummyInstaller{}
