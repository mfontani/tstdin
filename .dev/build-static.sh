#!/bin/sh

# fatal: detected dubious ownership in repository at '/__w/tstdin/tstdin'
git config --global --add safe.directory "$(pwd)"
go build -tags timetzdata --ldflags "-X 'main.Version=$(git describe --tags)' -extldflags \"-static\" -s -w" -o tstdin .
