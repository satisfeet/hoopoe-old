SHELL := /bin/bash

boot:
	@go run main.go

test: test-store

test-store:
	@go test -v github.com/satisfeet/hoopoe/store

test-httpd:
	@go test -v github.com/satisfeet/hoopoe/httpd

.PHONY: test
