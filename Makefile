SHELL := /bin/bash

boot:
	@go run main.go \
		--addr localhost:3000 \
		--mongo localhost/satisfeet

test: test-conf test-store

test-conf:
	@go test github.com/satisfeet/hoopoe/conf

test-store:
	@go test github.com/satisfeet/hoopoe/store

test-httpd:
	@go test github.com/satisfeet/hoopoe/httpd

.PHONY: test
