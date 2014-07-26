SHELL := /bin/bash

boot:
	@go run main.go \
		--host localhost:3000 \
		--username bodokaiser \
		--password secret \
		--mongo localhost/satisfeet

test:
	@go test ./...

test-conf:
	@go test ./conf/...

test-store:
	@go test ./store/...

test-email:
	@go test ./email/...

test-httpd:
	@go test ./httpd/...

.PHONY: test test-conf test-store test-email test-httpd
