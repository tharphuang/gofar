#!/bin/sh

set -e

function modify() {
    source_head=Mgoogle/protobuf/
    # target=github.com/lqs/gogoslim/types
    target=github.com/golang/protobuf/ptypes
    file_names=("any.proto" "duration.proto" "empty.proto" "struct.proto" "timestamp.proto" "wrappers.proto")
    result=""
    for name in "${file_names[@]}"; do
        tmp=(${name//./ })
        dir=${tmp[0]}
        result=${source_head}${name}=${target}/${dir},${result}
    done
    for file in `find proto -name '*.proto' |cut -sd / -f 2-`; do
        result=M${file}=$(basename $(pwd))/generated/grpc/`dirname ${file}`,${result}
    done
    echo ${result}
}

function protogen() {
    rm -rf generated
    mkdir -p generated/grpc
    proto_modify=$(modify)
    echo ${proto_modify}
    echo "Generating proto"
    for file in `find proto -name '*.proto' |cut -sd / -f 2-`; do
        # protoc -I proto "${file}" --gogoslim_out=${proto_modify},plugins=grpc:generated/grpc &
        protoc -I proto "${file}" --go_out=${proto_modify},plugins=grpc:generated/grpc &
    done
    wait
}


if [[  $# -eq 0 ]] ; then
    deploy
    echo "Deploy finished!"
    exit 0
fi

case "$1" in
  generate)
    protogen $@
  ;;
  *)
    echo "Usage: sh deploy.sh [generate|build|deploy] [-f=Dockerfile]"
  ;;
esac
