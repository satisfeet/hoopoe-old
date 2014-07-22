SHELL := /bin/bash

boot:
	@go run main.go \
		--host localhost:3000 \
		--auth bodokaiser:secret \
		--mongo localhost/satisfeet

test:
	@go test ./...

test-conf:
	@go test ./conf/...

test-httpd:
	@go test ./httpd/...

test-model:
	@go test ./model/...

test-store:
	@go test ./store/...

.PHONY: test test-conf test-httpd test-model test-store
