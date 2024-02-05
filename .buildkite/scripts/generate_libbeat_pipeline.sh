#!/usr/bin/env bash

source .buildkite/scripts/common.sh

set -euo pipefail

pipelineName="pipeline.libbeat-dynamic.yml"

cat > $pipelineName <<- YAML

steps:

YAML

if are_conditions_met_mandatory_tests; then
  cat >> $pipelineName <<- YAML

  - group: "Mandatory Tests"
    key: "mandatory-tests"
    steps:
      - label: ":linux: Ubuntu Unit Tests"
        key: "mandatory-linux-unit-test"
        command: ".buildkite/scripts/unit_tests.sh ${PIPELINE_NAME}"
        agents:
          provider: "gcp"
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "c2-standard-16"
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.xml"

      - label: ":go: Go Intergration Tests"
        key: "mandatory-int-test"
        command: ".buildkite/scripts/go_int_tests.sh ${PIPELINE_NAME}"
        agents:
          provider: "gcp"
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "c2-standard-16"
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.xml"

      - label: ":python: Python Integration Tests"
        key: "mandatory-python-int-test"
        command: ".buildkite/scripts/py_int_tests.sh ${PIPELINE_NAME}"
        agents:
          provider: "gcp"
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "c2-standard-16"
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.xml"

      - label: ":negative_squared_cross_mark: Cross compile"
        key: "mandatory-cross-compile"
        command: ".buildkite/scripts/crosscompile.sh ${PIPELINE_NAME}"
        agents:
          provider: "gcp"
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "c2-standard-16"
        artifact_paths: " ${BEATS_PROJECT_NAME}/build/*.xml"

      - label: ":testengine: Stress Tests"
        key: "mandatory-stress-test"
        command: ".buildkite/scripts/stress_tests.sh ${PIPELINE_NAME}"
        agents:
          provider: "gcp"
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "c2-standard-16"
        artifact_paths: "${BEATS_PROJECT_NAME}/libbeat-stress-test.xml"

YAML
fi

if are_conditions_met_extended_tests && are_conditions_met_arm_tests; then
  cat >> $pipelineName <<- YAML

  - group: "Extended Tests"
    key: "extended-tests"
    steps:
      - label: ":linux: Arm64 Unit Tests"
        key: "extended-arm64-unit-tests"
        command: ".buildkite/scripts/unit_tests.sh ${PIPELINE_NAME}"
        agents:
          provider: "aws"
          imagePrefix: "${IMAGE_UBUNTU_ARM_64}"
          instanceType: "t4g.xlarge"
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.xml"

YAML
fi

echo "--- Printing dynamic steps"     #TODO: remove if the pipeline is public
cat $pipelineName

echo "--- Loading dynamic steps"
buildkite-agent pipeline upload $pipelineName
