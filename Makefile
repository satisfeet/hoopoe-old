SHELL := /bin/bash

boot:
	@go run main.go \
		--host :3000 \
		--mongo localhost/satisfeet

test:
	@go test ./...

test-httpd:
	@go test ./httpd/...

test-model:
	@go test ./model/...

test-store:
	@go test ./store/...

.PHONY: test test-httpd test-store
