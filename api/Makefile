SHELL := /bin/bash

migrate_path = ~/Programs/migrate.linux-amd64
swagger_path = ~/Programs/swagger/swag

.PHONY: build
build:
	go build -v ./cmd/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: migrate_init
migrate_init:
	$(migrate_path) create -ext sql -dir ./schema -seq init

.PHONY: swag
swag:
	$(swagger_path) init -g cmd/app/main.go

.PHONY: lint
lint:
	golangci-lint run --color always

.PHONY: pipeline
pipeline:
	make lint && make swag && make test

.DEFAULT_GOAL := pipeline
