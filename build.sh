#!/bin/bash

GOPATH=/Users/pzqu/Documents/code/go/process_create_time CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o bin/process_tool src/main.go