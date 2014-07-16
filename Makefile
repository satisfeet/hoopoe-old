SHELL := /bin/bash

boot:
	@go run main.go \
		--host :3000 \
		--mongo localhost/satisfeet

test:
	@go test ./...

test-httpd:
	@go test -v ./httpd/...

test-store:
	@go test -v ./store/...

.PHONY: test test-httpd test-store
