#!/bin/bash

source .buildkite/scripts/common.sh

set -euo pipefail

pipelineName="pipeline.metricbeat-dynamic.yml"

cat > $pipelineName <<- YAML

steps:

# YAML

# if are_conditions_met_mandatory_tests; then
#   cat >> $pipelineName <<- YAML

  - group: "Mandatory Tests"
    key: "mandatory-tests"
    steps:
      - label: ":linux: Ubuntu Unit Tests"
        key: "mandatory-linux-unit-test"
        command: ".buildkite/scripts/unit_tests.sh"
        agents:
          provider: "gcp"
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "c2-standard-16"
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.*"

      - label: ":windows: Windows Unit Tests - {{matrix.image}}"
        command: ".buildkite/scripts/win_unit_tests.ps1"
        key: "mandatory-win-unit-tests"
        agents:
          provider: "gcp"
          image: "{{matrix.image}}"
          machine_type: "n2-standard-8"
          disk_size: 100
          disk_type: "pd-ssd"
        matrix:
          setup:
            image:
              - "${IMAGE_WIN_2016}"
              - "${IMAGE_WIN_2022}"
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.*"

# YAML
# fi

# if are_conditions_met_extended_tests && are_conditions_met_macos_tests; then
#   cat >> $pipelineName <<- YAML

  - group: "Extended Tests"
    key: "extended-tests"
    steps:
      - label: ":mac: MacOS Unit Tests"
        key: "extended-macos-unit-tests"
        command: ".buildkite/scripts/unit_tests.sh"
        agents:
          provider: "orka"
          imagePrefix: "${IMAGE_MACOS_X86_64}"
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.*"

      - label: ":linux: ARM Ubuntu Unit Tests"
        key: "extended-arm64-unit-test"
        command: ".buildkite/scripts/unit_tests.sh"
        agents:
          provider: "aws"
          imagePrefix: "${IMAGE_UBUNTU_ARM_64}"
          instanceType: "t4g.xlarge"
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.*"

# YAML
# fi

# if are_conditions_met_extended_windows_tests; then
#   cat >> $pipelineName <<- YAML

  - group: "Extended Windowds Tests"
    key: "extended-win-tests"
    steps:
      - label: ":windows: Win 2019 Unit Tests"
        key: "extended-win-2019-unit-tests"
        command: ".buildkite/scripts/win_unit_tests.ps1"
        agents:
          provider: "gcp"
          image: "${IMAGE_WIN_2019}"
          machine_type: "n2-standard-8"
          disk_size: 100
          disk_type: "pd-ssd"
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.*"

# YAML
# fi

# if are_conditions_met_extended_windows_tests; then
#   cat >> $pipelineName <<- YAML

  - group: "Packaging"    # TODO: check conditions for future the main pipeline migration: https://github.com/elastic/beats/pull/28589
    key: "packaging"
    steps:
      - label: ":linux: Packaging Linux"
        key: "packaging-linux"
        command: ".buildkite/scripts/packaging.sh"
        agents:
          provider: "gcp"
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "c2-standard-16"
        env:
          PLATFORMS: "+all linux/amd64 linux/arm64 windows/amd64 darwin/amd64 darwin/arm64"

      - label: ":linux: Packaging ARM"
        key: "packaging-arm"
        command: ".buildkite/scripts/packaging.sh"
        agents:
          provider: "aws"
          imagePrefix: "${IMAGE_UBUNTU_ARM_64}"
          instanceType: "t4g.xlarge"
        env:
          PLATFORMS: "linux/arm64"
          PACKAGES: "docker"

    depends_on:
    - step: "mandatory-tests"
      allow_failure: false
    - step: "extended-tests"
      allow_failure: true
    - step: "extended-win-tests"
      allow_failure: true

YAML
# fi

cat $pipelineName       #remove after tests

buildkite-agent pipeline upload $pipelineName
