SHELL := /bin/bash
include .env.example

run:
	@ BS_ENV=${BS_ENV} TEST_ENV=${TEST_ENV} SECRET_KEY_API_KEY=${SECRET_KEY_API_KEY} PG_SECRET=${PG_SECRET} go run ./main.go
.PHONY: run

playground:
	@ BS_ENV=${BS_ENV} TEST_ENV=${TEST_ENV} SECRET_KEY_API_KEY=${SECRET_KEY_API_KEY} PG_SECRET=${PG_SECRET} go run playground/main.go
.PHONY: playground