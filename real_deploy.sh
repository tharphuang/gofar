#!/bin/bash

set -e

protogen() {
    rm -rf generated
    mkdir -p generated/grpc
    for file in $(find proto -name '*.proto' |cut -sd / -f 2-); do
        echo "${file}" 
        protoc -I proto "${file}" --go_out=. --go-grpc_out=. &
    done
    wait
}

deploy() {
  modify
}

case "$1" in
  generate)
    protogen "$@"
  ;;
  
  *)
    echo "Usage: sh deploy.sh [generate|build|deploy] [-f=Dockerfile]"
  ;;
esac
