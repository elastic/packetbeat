// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build !aix

package azureeventhub

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/elastic/elastic-agent-libs/logp"
)

type azureInputConfig struct {
	ConnectionString string `config:"connection_string" validate:"required"`
	EventHubName     string `config:"eventhub" validate:"required"`
	ConsumerGroup    string `config:"consumer_group"`
	// Azure Storage container to store leases and checkpoints
	SAName             string `config:"storage_account" validate:"required"`
	SAKey              string `config:"storage_account_key"`               // (processor v1 only)
	SAConnectionString string `config:"storage_account_connection_string"` // (processor v2 only)
	SAContainer        string `config:"storage_account_container"`
	// by default the azure public environment is used, to override, users can provide a specific resource manager endpoint
	OverrideEnvironment string `config:"resource_manager_endpoint"`
	// cleanup the log JSON input for known issues, options: SINGLE_QUOTES, NEW_LINES
	SanitizeOptions []string `config:"sanitize_options"`
	// Processor version to use (v1 or v2). Default is v1 (processor v2 only).
	ProcessorVersion string `config:"processor_version"`
	// Controls if the input should perform the checkpoint information
	// migration from v1 to v2 (processor v2 only).
	MigrateCheckpoint bool `config:"migrate_checkpoint"`
	// Controls the start position for all partitions (processor v2 only).
	StartPosition string `config:"start_position"`
	// Processor receive timeout (processor v2 only).
	// Wait up to `ReceiveTimeout` for `ReceiveCount` events,
	// otherwise returns whatever we collected during that time.
	ReceiveTimeout time.Duration `config:"receive_timeout"`
	// Processor receive count (processor v2 only).
	// Wait up to `ReceiveTimeout` for `ReceiveCount` events,
	// otherwise returns whatever we collected during that time.
	ReceiveCount int `config:"receive_count"`
}

const ephContainerName = "filebeat"

// Validate validates the config.
func (conf *azureInputConfig) Validate() error {
	logger := logp.NewLogger("azureeventhub.config")
	if conf.ConnectionString == "" {
		return errors.New("no connection string configured")
	}
	if conf.EventHubName == "" {
		return errors.New("no event hub name configured")
	}
	if conf.SAName == "" {
		return errors.New("no storage account configured (config: storage_account)")
	}
	if conf.SAContainer == "" {
		conf.SAContainer = fmt.Sprintf("%s-%s", ephContainerName, conf.EventHubName)
	}
	if strings.Contains(conf.SAContainer, "_") {
		originalValue := conf.SAContainer
		// When a user specifies an event hub name in the input settings,
		// the configuration uses it to compose the storage account (SA) container
		// name (for example, `filebeat-<DATA-STREAM>-<EVENTHUB>`).
		//
		// The event hub allows names with underscores (_) characters, but unfortunately,
		// the SA container does not permit them.
		//
		// So instead of throwing an error to the user, we decided to replace
		// underscores (_) characters with hyphens (-).
		conf.SAContainer = strings.ReplaceAll(conf.SAContainer, "_", "-")
		logger.Warnf("replaced underscores (_) with hyphens (-) in the storage account container name (before: %s, now: %s", originalValue, conf.SAContainer)
	}
	err := storageContainerValidate(conf.SAContainer)
	if err != nil {
		return err
	}

	// log a warning for each sanitization option not supported
	for _, opt := range conf.SanitizeOptions {
		err := sanitizeOptionsValidate(opt)
		if err != nil {
			logger.Warnf("%s: %v", opt, err)
		}
	}

	if conf.ProcessorVersion == "" {
		// The default processor version is "v1".
		conf.ProcessorVersion = processorV1
	}

	switch conf.ProcessorVersion {
	case processorV1:
		if conf.SAKey == "" {
			return errors.New("no storage account key configured (config: storage_account_key)")
		}
	case processorV2:
		if conf.SAKey != "" {
			logger.Warnf("storage_account_key is not used in processor v2")
		}
		if conf.SAConnectionString == "" {
			return errors.New("no storage account connection string configured (config: storage_account_connection_string)")
		}
	default:
		return fmt.Errorf("invalid azure-eventhub processor version: %s (available versions: v1, v2)", conf.ProcessorVersion)
	}

	if conf.StartPosition == "" {
		// For backward compatibility with v1,
		// the default start position is "earliest".
		conf.StartPosition = startPositionEarliest
	}

	// Default receive timeout and count.
	if conf.ReceiveTimeout == 0 {
		// The default receive timeout is 5 second.
		conf.ReceiveTimeout = 5 * time.Second
	}
	if conf.ReceiveCount == 0 {
		// The default receive count is 100.
		conf.ReceiveCount = 100
	}

	return nil
}

// storageContainerValidate validated the storage_account_container to make sure it is conforming to all the Azure
// naming rules.
// To learn more, please check the Azure documentation visiting:
// https://docs.microsoft.com/en-us/rest/api/storageservices/naming-and-referencing-containers--blobs--and-metadata#container-names
func storageContainerValidate(name string) error {
	var previousRune rune
	runes := []rune(name)
	length := len(runes)
	if length < 3 {
		return fmt.Errorf("storage_account_container (%s) must be 3 or more characters", name)
	}
	if length > 63 {
		return fmt.Errorf("storage_account_container (%s) must be less than 63 characters", name)
	}
	if !unicode.IsLower(runes[0]) && !unicode.IsNumber(runes[0]) {
		return fmt.Errorf("storage_account_container (%s) must start with a lowercase letter or number", name)
	}
	if !unicode.IsLower(runes[length-1]) && !unicode.IsNumber(runes[length-1]) {
		return fmt.Errorf("storage_account_container (%s) must end with a lowercase letter or number", name)
	}
	for i := 0; i < length; i++ {
		if !unicode.IsLower(runes[i]) && !unicode.IsNumber(runes[i]) && !(runes[i] == '-') {
			return fmt.Errorf("rune %d of storage_account_container (%s) is not a lowercase letter, number or dash", i, name)
		}
		if runes[i] == '-' && previousRune == runes[i] {
			return fmt.Errorf("consecutive dashes ('-') are not permitted in storage_account_container (%s)", name)
		}
		previousRune = runes[i]
	}
	return nil
}
