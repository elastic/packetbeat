#!/usr/bin/env bash
BEAT_DIR=$1

if [ -z "$BEAT_DIR" ]; then
    echo "Error: Beat directory must be specified."
    exit 1
fi

echo ">>> Packaging : $BEAT_DIR"

WORKSPACE=$(pwd)

BEAT_NAME_SLUG=$(echo "$BEAT_DIR" | sed 's/x-pack\///g')

cd $BEAT_DIR
mage package
mage ironbank

cp build/distributions/* $WORKSPACE/build/distributions/$BEAT_NAME_SLUG/
ls -l $WORKSPACE/build/distributions/$BEAT_NAME_SLUG/