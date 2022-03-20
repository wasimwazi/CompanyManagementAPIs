#!/bin/bash

source config/development.env

pushd migrations
goose postgres "$DBConString" up
popd

go run main.go

go test ./...