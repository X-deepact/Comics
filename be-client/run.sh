#!/bin/bash

echo "Remove back-end-client: rm -rf ./server"
rm -rf ./server

echo "make swagger doccument"
make docs

echo "Build back-end-client: go mod tidy && go mod vendor"
go mod tidy && go mod vendor

echo "Build binary back-end-client: go build -o back-end-client main.go"
go build -o server main.go

echo "Run back-end-client"
./server -config config -env dev


