#!/bin/bash

basedir=`cd $(dirname $0); pwd -P`
#export GOPATH=$basedir

go build -o $OUTPUT/test_crypto github.com/xuperchain/crypto/test

echo 'start build plugins'
go build --buildmode=plugin -o $OUTPUT/crypto-default.so.1.0.0 github.com/xuperchain/crypto/client/xchain/