#!/usr/bin/env bash

source .buildkite/scripts/common.sh

set -euo pipefail

pipelineName="pipeline.filebeat-dynamic.yml"

<<<<<<< HEAD
# TODO: steps: must be always included 
=======
# TODO: steps: must be always included
>>>>>>> 80dab50f0c (replace default images (#38583))
echo "Add the mandatory and extended tests without additional conditions into the pipeline"
if are_conditions_met_mandatory_tests; then
  cat > $pipelineName <<- YAML


steps:
<<<<<<< HEAD
  - group: "Filebeat Mandatory Testing"
    key: "mandatory-tests"
    if: build.env("GITHUB_PR_TRIGGER_COMMENT") == "filebeat" || build.env("BUILDKITE_PULL_REQUEST") != "false"

    steps:
      - label: ":ubuntu: Unit Tests"
        command:
          - ".buildkite/filebeat/scripts/unit-tests.sh"
        notify:
          - github_commit_status:
              context: "Filebeat: linux/Unit Tests"
=======
  - group: "Mandatory Testing"
    key: "mandatory-tests"

    steps:
      - label: ":ubuntu: Ubuntu Unit Tests"
        command: "cd $BEATS_PROJECT_NAME && mage build unitTest"
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Ununtu Unit Tests"
>>>>>>> 80dab50f0c (replace default images (#38583))
        agents:
          provider: "gcp"
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "${GCP_DEFAULT_MACHINE_TYPE}"
        artifact_paths:
          - "filebeat/build/*.xml"
          - "filebeat/build/*.json"

<<<<<<< HEAD
      - label: ":ubuntu: Go Integration Tests"
        command:
          - ".buildkite/filebeat/scripts/integration-gotests.sh"
        notify:
          - github_commit_status:
              context: "Filebeat: Go Integration Tests"
=======
      - label: ":ubuntu: Ubuntu Go Integration Tests"
        command: "cd $BEATS_PROJECT_NAME && mage goIntegTest"
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Go Integration Tests"
>>>>>>> 80dab50f0c (replace default images (#38583))
        agents:
          provider: "gcp"
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "${GCP_HI_PERF_MACHINE_TYPE}"
        artifact_paths:
          - "filebeat/build/*.xml"
          - "filebeat/build/*.json"

<<<<<<< HEAD
      - label: ":ubuntu: Python Integration Tests"
        command:
          - ".buildkite/filebeat/scripts/integration-pytests.sh"
        notify:
          - github_commit_status:
              context: "Filebeat: Python Integration Tests"
=======
      - label: ":ubuntu: Ubuntu Python Integration Tests"
        command: "cd $BEATS_PROJECT_NAME && mage pythonIntegTest"
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Python Integration Tests"
>>>>>>> 80dab50f0c (replace default images (#38583))
        agents:
          provider: gcp
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "${GCP_HI_PERF_MACHINE_TYPE}"
        artifact_paths:
          - "filebeat/build/*.xml"
          - "filebeat/build/*.json"

<<<<<<< HEAD
      - label: ":windows:-2016 Unit Tests"
        command: ".buildkite/scripts/win_unit_tests.ps1"
        notify:
          - github_commit_status:
              context: "Filebeat: windows/Unit Tests"
=======
      - label: ":windows: Windows 2016 Unit Tests"
        key: "windows-2016-unit-tests"
        command: |
          Set-Location -Path $BEATS_PROJECT_NAME
          mage build unitTest
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Windows 2016 Unit Tests"
>>>>>>> 80dab50f0c (replace default images (#38583))
        agents:
          provider: "gcp"
          image: "${IMAGE_WIN_2016}"
          machine_type: "${GCP_WIN_MACHINE_TYPE}"
          disk_size: 200
          disk_type: "pd-ssd"
<<<<<<< HEAD
        artifact_paths:
          - "filebeat/build/*.xml"
          - "filebeat/build/*.json"

      - label: ":windows:-2022 Unit Tests"
        command: ".buildkite/scripts/win_unit_tests.ps1"
        notify:
          - github_commit_status:
              context: "Filebeat: windows 2022/Unit Tests"
=======
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.*"

      - label: ":windows: Windows 2022 Unit Tests"
        key: "windows-2022-unit-tests"
        command: |
          Set-Location -Path $BEATS_PROJECT_NAME
          mage build unitTest
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Windows 2022 Unit Tests"
>>>>>>> 80dab50f0c (replace default images (#38583))
        agents:
          provider: "gcp"
          image: "${IMAGE_WIN_2022}"
          machine_type: "${GCP_WIN_MACHINE_TYPE}"
          disk_size: 200
<<<<<<< HEAD
          disk_type: "pd-ssd"        
        artifact_paths:
          - "filebeat/build/*.xml"
          - "filebeat/build/*.json"
=======
          disk_type: "pd-ssd"
        artifact_paths: "${BEATS_PROJECT_NAME}/build/*.*"
>>>>>>> 80dab50f0c (replace default images (#38583))

YAML
else
  echo "The conditions don't match to requirements for generating pipeline steps."
  exit 0
fi

echo "Check and add the Extended Tests into the pipeline"

<<<<<<< HEAD
if are_conditions_met_arm_tests; then
  cat >> $pipelineName <<- YAML
  - group: "Extended Tests: ARM"
      key: "extended-tests-arm"
      steps:
      - label: ":linux: ARM64 Unit Tests"
        key: "arm-extended"        
        command:
          - ".buildkite/filebeat/scripts/unit-tests.sh"
        notify:
          - github_commit_status:
              context: "Filebeat/Extended: Unit Tests ARM"
=======
if are_conditions_met_arm_tests || are_conditions_met_macos_tests; then
  cat >> $pipelineName <<- YAML

  - group: "Extended Tests"
    key: "extended-tests"
    steps:
YAML
fi

if are_conditions_met_macos_tests; then
  cat >> $pipelineName <<- YAML

      - label: ":mac: MacOS Unit Tests"
        key: "macos-unit-tests-extended"
        command: "cd $BEATS_PROJECT_NAME && mage build unitTest"
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: MacOS Unit Tests"
        agents:
          provider: "orka"
          imagePrefix: "${IMAGE_MACOS_X86_64}"
        artifact_paths:
          - "$BEATS_PROJECT_NAME/build/*.xml"
          - "$BEATS_PROJECT_NAME/build/*.json"

      - label: ":mac: MacOS ARM Unit Tests"
        key: "macos-arm64-unit-tests-extended"
        command: "cd $BEATS_PROJECT_NAME && mage build unitTest"
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: MacOS ARM Unit Tests"
        agents:
          provider: "orka"
          imagePrefix: "${IMAGE_MACOS_ARM}"
        artifact_paths:
          - "$BEATS_PROJECT_NAME/build/*.xml"
          - "$BEATS_PROJECT_NAME/build/*.json"
YAML
fi

if are_conditions_met_arm_tests; then
  cat >> $pipelineName <<- YAML

      - label: ":linux: Ubuntu ARM Unit Tests"
        key: "extended-arm64-unit-test"
        command: "cd $BEATS_PROJECT_NAME && mage build unitTest"
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Ubuntu ARM Unit Tests"
>>>>>>> 80dab50f0c (replace default images (#38583))
        agents:
          provider: "aws"
          imagePrefix: "${AWS_IMAGE_UBUNTU_ARM_64}"
          instanceType: "${AWS_ARM_INSTANCE_TYPE}"
<<<<<<< HEAD
        artifact_paths: "filebeat/build/*.xml"

=======
        artifact_paths:
          - "$BEATS_PROJECT_NAME/build/*.xml"
          - "$BEATS_PROJECT_NAME/build/*.json"
>>>>>>> 80dab50f0c (replace default images (#38583))
YAML
fi

if are_conditions_met_win_tests; then
  cat >> $pipelineName <<- YAML
<<<<<<< HEAD
  - group: "Windows Extended Testing"
    key: "extended-tests-win"
    steps:
    - label: ":windows: Win 2019 Unit Tests"
      key: "win-extended-2019"
      command: ".buildkite/scripts/win_unit_tests.ps1"
      notify:
        - github_commit_status:
            context: "Filebeat/Extended: Win-2019 Unit Tests"
      agents:
        provider: "gcp"
        image: "${IMAGE_WIN_2019}"
        machine_type: "${GCP_WIN_MACHINE_TYPE}"
        disk_size: 200
        disk_type: "pd-ssd"
      artifact_paths:
        - "filebeat/build/*.xml"
        - "filebeat/build/*.json"
=======

  - group: "Windows Extended Testing"
    key: "extended-tests-win"
    steps:
      - label: ":windows: Windows 2019 Unit Tests"
        key: "windows-extended-2019"
        command: |
          Set-Location -Path $BEATS_PROJECT_NAME
          mage build unitTest
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Windows 2019 Unit Tests"
        agents:
          provider: "gcp"
          image: "${IMAGE_WIN_2019}"
          machine_type: "${GCP_WIN_MACHINE_TYPE}"
          disk_size: 200
          disk_type: "pd-ssd"
        artifact_paths:
          - "$BEATS_PROJECT_NAME/build/*.xml"
          - "$BEATS_PROJECT_NAME/build/*.json"

      - label: ":windows: Windows 11 Unit Tests"
        key: "windows-extended-11"
        command: |
          Set-Location -Path $BEATS_PROJECT_NAME
          mage build unitTest
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Windows 11 Unit Tests"
        agents:
          provider: "gcp"
          image: "${IMAGE_WIN_11}"
          machine_type: "${GCP_WIN_MACHINE_TYPE}"
          disk_size: 200
          disk_type: "pd-ssd"
        artifact_paths:
          - "$BEATS_PROJECT_NAME/build/*.xml"
          - "$BEATS_PROJECT_NAME/build/*.json"

      - label: ":windows: Windows 10 Unit Tests"
        key: "windows-extended-10"
        command: |
          Set-Location -Path $BEATS_PROJECT_NAME
          mage build unitTest
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Windows 10 Unit Tests"
        agents:
          provider: "gcp"
          image: "${IMAGE_WIN_10}"
          machine_type: "${GCP_WIN_MACHINE_TYPE}"
          disk_size: 200
          disk_type: "pd-ssd"
        artifact_paths:
          - "$BEATS_PROJECT_NAME/build/*.xml"
          - "$BEATS_PROJECT_NAME/build/*.json"
>>>>>>> 80dab50f0c (replace default images (#38583))
YAML
fi

echo "Check and add the Packaging into the pipeline"
if are_conditions_met_packaging; then
cat >> $pipelineName <<- YAML
<<<<<<< HEAD
  - group: "Packaging"
    key: "packaging"
    if: build.env("BUILDKITE_PULL_REQUEST") != "false"
=======

  - group: "Packaging"
    key: "packaging"
>>>>>>> 80dab50f0c (replace default images (#38583))
    depends_on:
      - "mandatory-tests"

    steps:
<<<<<<< HEAD
      - label: Package pipeline
        commands: ".buildkite/scripts/packaging/package-step.sh"
        notify:
        - github_commit_status:
            context: "Filebeat: Packaging"
=======
      - label: ":linux: Packaging Linux"
        key: "packaging-linux"
        command: "cd $BEATS_PROJECT_NAME && mage package"
        agents:
          provider: "gcp"
          image: "${IMAGE_UBUNTU_X86_64}"
          machineType: "${GCP_HI_PERF_MACHINE_TYPE}"
          disk_size: 100
          disk_type: "pd-ssd"
        env:
          PLATFORMS: "${PACKAGING_PLATFORMS}"
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Packaging Linux"

      - label: ":linux: Packaging ARM"
        key: "packaging-arm"
        command: "cd $BEATS_PROJECT_NAME && mage package"
        agents:
          provider: "aws"
          imagePrefix: "${AWS_IMAGE_UBUNTU_ARM_64}"
          instanceType: "${AWS_ARM_INSTANCE_TYPE}"
        env:
          PLATFORMS: "${PACKAGING_ARM_PLATFORMS}"
          PACKAGES: "docker"
        notify:
          - github_commit_status:
              context: "$BEATS_PROJECT_NAME: Packaging Linux ARM"
>>>>>>> 80dab50f0c (replace default images (#38583))

YAML
fi

<<<<<<< HEAD
echo "--- Printing dynamic steps"     #TODO: remove if the pipeline is public
cat $pipelineName
=======
echo "+++ Printing dynamic steps"
cat $pipelineName | yq . -P
>>>>>>> 80dab50f0c (replace default images (#38583))

echo "--- Loading dynamic steps"
buildkite-agent pipeline upload $pipelineName
