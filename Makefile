SHELL := /bin/bash
include .env.example

run:
	@ BS_ENV=${BS_ENV} TEST_ENV=${TEST_ENV} SECRET_KEY_API_KEY=${SECRET_KEY_API_KEY} PG_SECRET=${PG_SECRET} go run ./main.go
.PHONY: run

build:
	@ BS_ENV=${BS_ENV} TEST_ENV=${TEST_ENV} SECRET_KEY_API_KEY=${SECRET_KEY_API_KEY} PG_SECRET=${PG_SECRET} go build -o bss ./main.go
.PHONY: build