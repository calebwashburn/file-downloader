#!/bin/bash -e

export GOPATH=$PWD/go
export PATH=$GOPATH/bin:$PATH

git config --global user.email "you@example.com"
git config --global user.name "Your Name"
go get -u github.com/golang/dep/cmd/dep
WORKING_DIR=$GOPATH/src/github.com/pivotalservices/file-downloader-resource
mkdir -p ${WORKING_DIR}
cp -R source/* ${WORKING_DIR}/.
cd ${WORKING_DIR}
go version
dep ensure
go test ./... -v
