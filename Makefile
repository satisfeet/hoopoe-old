SHELL := /bin/bash

boot:
	@go run cmd/main.go \
		--host :3000 --mongo mongodb://localhost/satisfeet

test:
	@go test ./...

test-net:
	@go test -v ./net/...

test-store:
	@go test -v ./store/...

.PHONY: test test-net test-store
