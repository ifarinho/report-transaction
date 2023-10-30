#!/usr/bin/env make

DOCKER_IMAGE=report-transaction
TEMPLATE_ENV=template.env

.PHONY: help
help:
	@echo "Command list: \n- help: display this message \n- build: compiles the program and builds the image \n- dotenv: creates .env file for local use \n- local: runs the image with local .env \n- run: runs the program"

.PHONY: build
build:
	@docker build -t $(DOCKER_IMAGE) .

.PHONY: run
run:
	@docker run --rm -it $(DOCKER_IMAGE)

.PHONY: dotenv
dotenv:
	@if [ ! -f $(CURDIR)/.env ]; then cp -f $(CURDIR)/$(TEMPLATE_ENV) $(CURDIR)/.env; fi

.PHONY: local
local:
	@$(shell ./scripts/run_local.sh $(FILENAME) $(ACCOUNT))
