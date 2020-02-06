#!/usr/bin/env bash
set -e

readonly LOCATION="${1?Please define the path where the fix permissions should run from}"

# Change ownership of all files inside the specific folder from root/root to current user/group
docker run -v ${LOCATION}:/beat alpine:3.4 sh -c "find /beat -user 0 -exec chown -h $(id -u):$(id -g) {} \;"
