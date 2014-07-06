SHELL := /bin/bash

boot:
	@go run main.go \
		--addr localhost:3000 \
		--mongo localhost/satisfeet

test: test-store

test-store:
	@go test github.com/satisfeet/hoopoe/store

test-httpd:
	@go test github.com/satisfeet/hoopoe/httpd

.PHONY: test
