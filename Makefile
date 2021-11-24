.PHONY: build
build:
	go build -o http-server -v ./cmd/app/

.DEFAULT_GOAL := build