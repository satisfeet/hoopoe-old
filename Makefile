SHELL := /bin/bash

boot:
	@go run main.go \
		--addr localhost:3000 \
		--mongo localhost/satisfeet

test:
	@go test ./...

test-conf:
	@go test -v ./conf

test-store:
	@go test -v ./store

test-httpd:
	@go test -v ./httpd

.PHONY: test test-conf test-store test-httpd
