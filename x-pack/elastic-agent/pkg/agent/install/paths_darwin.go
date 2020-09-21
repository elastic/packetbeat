// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build darwin

package install

const (
	// BinaryName is the name of the installed binary.
	BinaryName = "elastic-agent"

	// InstallPath is the installation path using for install command.
	InstallPath = "/Library/Elastic/Agent"

	// SocketPath is the socket path used when installed.
	SocketPath = "unix:///run/elastic-agent.sock"

	// ServiceName is the service name when installed.
	ServiceName = "com.elastic.elastic-agent"

	// ShellWrapperPath is the path to the installed shell wrapper.
	ShellWrapperPath = "/usr/local/bin/elastic-agent"

	// ShellWrapper is the wrapper that is installed.
	ShellWrapper = `#!/bin/sh
exec /Library/Elastic/Agent/elastic-agent $@
`
)
