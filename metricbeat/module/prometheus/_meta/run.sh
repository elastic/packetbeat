#!/bin/sh

HOST_DOMAIN="host.docker.internal"
ping -q -c1 $HOST_DOMAIN > /dev/null 2>&1
if [ $? -ne 0 ]; then
  HOST_DOMAIN="0.0.0.0"
fi

REMOTE="$HOST_DOMAIN:9201"

sed -i "s/REMOTE/$REMOTE/g" /etc/prometheus/prometheus.yml

cat /etc/prometheus.yml

/bin/prometheus --config.file=/etc/prometheus/prometheus.yml \
 --storage.tsdb.path=/prometheus \
 --web.console.libraries=/usr/share/prometheus/console_libraries \
 --web.console.templates=/usr/share/prometheus/consoles
