SHELL := /bin/bash

boot:
	@go run main.go

test: test-store

test-store:
	@go test github.com/satisfeet/hoopoe/store/customers

.PHONY: test
