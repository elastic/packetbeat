// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package artifact

import (
	"fmt"
	"path/filepath"
)

var packageArchMap = map[string]string{
	"linux-binary-32":   "linux-x86.tar.gz",
	"linux-binary-64":   "linux-x86_64.tar.gz",
	"windows-binary-32": "windows-x86.zip",
	"windows-binary-64": "windows-x86_64.zip",
	"darwin-binary-32":  "darwin-x86.tar.gz",
	"darwin-binary-64":  "darwin-x86.tar.gz",
}

// GetArtifactName constructs a path to a downloaded artifact
func GetArtifactName(program, version, operatingSystem, arch string) (string, error) {
	key := fmt.Sprintf("%s-binary-%s", operatingSystem, arch)
	suffix, found := packageArchMap[key]
	if !found {
		return "", fmt.Errorf("'%s' is not a valid combination for a package", key)
	}

	return fmt.Sprintf("%s-%s-%s", program, version, suffix), nil
}

// GetArtifactPath returns a full path of artifact for a program in specific version
func GetArtifactPath(programName, version, operatingSystem, arch, targetDir string) (string, error) {
	artifactName, err := GetArtifactName(programName, version, operatingSystem, arch)
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(targetDir, artifactName)
	return fullPath, nil
}
