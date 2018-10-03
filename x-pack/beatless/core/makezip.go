package core

import (
	yaml "gopkg.in/yaml.v2"

	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/x-pack/beatless/core/bundle"
)

// Package size limits for AWS lambda, we should be a lot under this limit but
// adding a check to make sure we never go over.
const packageCompressedLimit = 50 * 1000 * 1000    // 50MB
const packageUncompressedLimit = 250 * 1000 * 1000 // 250MB

func rawYaml() ([]byte, error) {
	// Load the configuration file from disk with all the settings,
	// the function takes care of using -c.

	// TODO: changes is made in another PR
	dummyCfg := common.MustNewConfigFrom(map[string]interface{}{})
	rawConfig, err := cfgfile.Load("", dummyCfg)
	if err != nil {
		return nil, err
	}
	var config map[string]interface{}
	if err := rawConfig.Unpack(&config); err != nil {
		return nil, err
	}

	res, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// MakeZip creates a zip from the the current artifacts and the currently available configuration.
func MakeZip() ([]byte, error) {
	rawConfig, err := rawYaml()
	if err != nil {
		return nil, err
	}
	bundle := bundle.NewZipWithLimits(
		packageUncompressedLimit,
		packageCompressedLimit,
		&bundle.MemoryFile{Path: "beatless.yml", Raw: rawConfig, FileMode: 0766},
		&bundle.LocalFile{Path: "pkg/beatless", FileMode: 0755},
	)

	content, err := bundle.Bytes()
	if err != nil {
		return nil, err
	}
	return content, nil
}
