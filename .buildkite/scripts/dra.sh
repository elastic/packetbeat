#!/usr/bin/env bash

# TODO: uncomment out below when Jenkins packaging has been stopped
# if [[ "$DRY_RUN" == "true" ]]; then
#     echo "~~~ Running in dry-run mode -- will NOT publish artifacts"
#     DRY_RUN="--dry-run"
# else
#     echo "~~~ Running in publish mode"
#     DRY_RUN=""
# fi

# TODO: delete the conditional below (and replace it with the above, uncommented out, section) after Jenkins packaging has been stopped
if [[ "$DRY_RUN" == "false" ]]; then
    echo "~~~ Running in publish mode"
    DRY_RUN=""
else
    echo "~~~ Running in dry-run mode -- will NOT publish artifacts"
    DRY_RUN="--dry-run"
fi

set -euo pipefail

# DRA_BRANCH can be used for manually testing packaging with PRs
# e.g. define `DRA_BRANCH="main"` and `RUN_SNAPSHOT="true"` under Options/Environment Variables in the Buildkite UI after clicking new Build
BRANCH="${DRA_BRANCH:="${BUILDKITE_BRANCH:=""}"}"

if [[ "${BUILDKITE_PULL_REQUEST:="false"}" != "false" ]]; then
    BRANCH=main
    DRY_RUN="--dry-run"
    echo "+++ Running in PR and setting branch main and --dry-run"
fi

BEAT_VERSION=$(make get-version)

CI_DRA_ROLE_PATH="kv/ci-shared/release/dra-role"

function release_manager_login {
  DRA_CREDS_SECRET=$(retry -t 5 -- vault kv get -field=data -format=json ${CI_DRA_ROLE_PATH})
  VAULT_ADDR_SECRET=$(echo ${DRA_CREDS_SECRET} | jq -r '.vault_addr')
  VAULT_ROLE_ID_SECRET=$(echo ${DRA_CREDS_SECRET} | jq -r '.role_id')
  VAULT_SECRET=$(echo ${DRA_CREDS_SECRET} | jq -r '.secret_id')
  export VAULT_ADDR_SECRET VAULT_ROLE_ID_SECRET VAULT_SECRET
}

set +x
release_manager_login


# required by the release-manager docker image, otherwise we hit:
# > java.io.FileNotFoundException: /artifacts/build/distributions/agentbeat/agentbeat-8.15.0-SNAPSHOT-darwin-x86_64.tar.gz.sha512 (Permission denied)
chmod -R a+r build/*
chmod -R a+w build

echo "+++ :clipboard: Listing DRA artifacts for branch [$BRANCH] using workflow [$DRA_WORKFLOW]"
set +x
docker run --rm \
        --name release-manager \
        -e VAULT_ADDR="${VAULT_ADDR_SECRET}" \
        -e VAULT_ROLE_ID="${VAULT_ROLE_ID_SECRET}" \
        -e VAULT_SECRET_ID="${VAULT_SECRET}" \
        --mount type=bind,readonly=false,src="${PWD}",target=/artifacts \
        docker.elastic.co/infra/release-manager:latest \
        cli list \
        --project "beats" \
        --branch "${BRANCH}" \
        --commit "${BUILDKITE_COMMIT}" \
        --workflow "${DRA_WORKFLOW}" \
        --version "${BEAT_VERSION}" \
        --artifact-set "main"

echo "+++ :hammer_and_pick: Publishing DRA artifacts for branch [$BRANCH] using workflow [$DRA_WORKFLOW] and DRY_RUN: [$DRY_RUN]"

set +x
docker run --rm \
        --name release-manager \
        -e VAULT_ADDR="${VAULT_ADDR_SECRET}" \
        -e VAULT_ROLE_ID="${VAULT_ROLE_ID_SECRET}" \
        -e VAULT_SECRET_ID="${VAULT_SECRET}" \
        --mount type=bind,readonly=false,src="${PWD}",target=/artifacts \
        docker.elastic.co/infra/release-manager:latest \
        cli collect \
        --project "beats" \
        --branch "${BRANCH}" \
        --commit "${BUILDKITE_COMMIT}" \
        --workflow "${DRA_WORKFLOW}" \
        --version "${BEAT_VERSION}" \
        --artifact-set "main" \
        ${DRY_RUN}
