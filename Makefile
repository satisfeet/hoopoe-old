SHELL := /bin/bash

MYSQL = root@/satisfeet
HOST  = 127.0.0.1:3000
AUTH  = bodokaiser:secret

start:
	HOOPOE_MYSQL=$(MYSQL) \
	HOOPOE_HOST=$(HOST) \
	HOOPOE_AUTH=$(AUTH) \
	go run main.go
