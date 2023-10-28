#!/usr/bin/env make

DOCKER_IMAGE=report-transaction

.PHONY: build
build:
	@cp template.env .env
	@docker build -t $(DOCKER_IMAGE):latest .

.PHONY: run
run:
	@docker-compose up -d
