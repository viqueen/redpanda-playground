#!/usr/bin/env bash

function spicedb() {
  go run ./cmd/spicedb/main.go "$@"
}

function kafka_session() {
  docker exec --workdir /opt/kafka/bin -it connect-output bash
}

function connect_create() {
   docker run --rm \
    docker.redpanda.com/redpandadata/connect \
    create "$@" > connect.yaml
}

function connect_run() {
  docker run --rm -it \
    --network "host" \
    --volume "${PWD}"/connect.yaml:/connect.yaml \
    docker.redpanda.com/redpandadata/connect \
    run
}


# shellcheck disable=SC2294
eval "$@"