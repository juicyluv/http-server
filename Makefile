.PHONY: build
build:
	go build -o http-server -v ./cmd/app/

run:
	go run cmd/app/main.go

.PHONY: test
test:
	go test -race -timeout 20s ./...

.SILENT: createdb
createdb:
	docker run --name=travels -e POSTGRES_PASSWORD='qwerty' -p 5437:5432 -d --rm postgres

.DEFAULT_GOAL := build