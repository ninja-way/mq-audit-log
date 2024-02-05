.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go build -o ./.bin/audit-server ./cmd/server

run: build
	./.bin/audit-server

