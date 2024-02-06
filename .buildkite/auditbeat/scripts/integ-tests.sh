#!/usr/bin/env bash

set -euo pipefail

source .buildkite/env-scripts/linux-env.sh

echo "--- Running Integration Tests"
sudo chmod -R go-w auditbeat/

cd auditbeat
umask 0022
mage build integTest
