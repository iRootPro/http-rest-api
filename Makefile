.PHONY: build
build:
	go build ./cmd/apiServer
start:
	./apiServer
test:
	go test ./... -v -race -timeout 30s
.DEFAULT_GOAL := build
