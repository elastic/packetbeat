#!/bin/bash

if [ -z "$KAFKA_ADVERTISED_HOST" ]; then
	KAFKA_ADVERTISED_HOST=$(dig +short $HOSTNAME):9092
fi

wait_for_port() {
    count=20
    port=$1
    while ! nc -z localhost $port && [[ $count -ne 0 ]]; do
        count=$(( $count - 1 ))
        [[ $count -eq 0 ]] && return 1
        sleep 0.5
    done
    # just in case, one more time
    nc -z localhost $port
}

echo "Starting ZooKeeper"
${KAFKA_HOME}/bin/zookeeper-server-start.sh ${KAFKA_HOME}/config/zookeeper.properties &
wait_for_port 2181

echo "Starting Kafka broker"
mkdir -p ${KAFKA_LOGS_DIR}
${KAFKA_HOME}/bin/kafka-server-start.sh ${KAFKA_HOME}/config/server.properties \
    --override delete.topic.enable=true \
    --override listeners=INSIDE://localhost:9091,OUTSIDE://0.0.0.0:9092 \
    --override advertised.listeners=INSIDE://localhost:9091,OUTSIDE://$KAFKA_ADVERTISED_HOST \
    --override listener.security.protocol.map=INSIDE:PLAINTEXT,OUTSIDE:PLAINTEXT \
    --override inter.broker.listener.name=INSIDE \
    --override logs.dir=${KAFKA_LOGS_DIR} &

wait_for_port 9092

echo "Kafka load status code $?"

# Make sure the container keeps running
tail -f /dev/null
