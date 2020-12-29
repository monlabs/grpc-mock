#!/bin/sh

pushd $(dirname "${BASH_SOURCE[0]}")
protoc -I/usr/local/include \
       -I. \
       -I./third_party/googleapis \
       --go_out=plugins=grpc:. \
       --grpc-gateway_out=logtostderr=true:. *.proto
popd
