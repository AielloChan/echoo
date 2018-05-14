#!/bin/bash

./package.sh

[ ! -r dist ] && mkdir dist

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build *.go
mv echoo dist/echoo_linux 
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build *.go
mv echoo dist/echoo_macos
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build *.go
mv echoo.exe dist/echoo_windows.exe