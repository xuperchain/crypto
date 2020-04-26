#!/bin/bash

basedir=`cd $(dirname $0); pwd -P`
#export GOPATH=$basedir

go build -o $OUTPUT/test_crypto github.com/xuperchain/crypto/test
go build -o $OUTPUT/test_gm_crypto github.com/xuperchain/crypto/test/gm

echo 'start build plugins'
go build --buildmode=plugin -o $OUTPUT/crypto-default.so.1.0.0 github.com/xuperchain/crypto/client/xchain/
go build --buildmode=plugin -o $OUTPUT/crypto-gm.so.1.0.0 github.com/xuperchain/crypto/client/gm/