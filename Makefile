.PHONY: build
build:
	go build -o http-server -v ./cmd/app/

run:
	go run cmd/app/main.go

.PHONY: test
test:
	go test -v -race -timeout 20s ./...

.DEFAULT_GOAL := build