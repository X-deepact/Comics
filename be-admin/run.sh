#!/bin/bash

echo "Remove back-end-admin: rm -rf ./server"
rm -rf ./server

echo "make swagger doccument"
make docs

echo "Build back-end-admin: go mod tidy && go mod vendor"
go mod tidy && go mod vendor

echo "Build binary back-end-admin: go build -o back-end-admin main.go"
go build -o server main.go

echo "Run back-end-admin"
./server -config config -env dev


