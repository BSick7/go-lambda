SHELL := /bin/bash

.PHONY: all
all: tools dep verify

tools:
	go get -u github.com/golang/dep/cmd/dep

dep:
	dep ensure

verify:
	go fmt ./...
	go vet ./...
	go test ./...
