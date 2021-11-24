.PHONY: build
build:
	go build -o http-server -v ./cmd/app/

run:
	go run cmd/app/main.go

.DEFAULT_GOAL := build