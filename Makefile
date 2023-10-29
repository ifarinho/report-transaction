#!/usr/bin/env make

DOCKER_IMAGE=report-transaction
TEMPLATE_ENV=template.env

.PHONY: env
build:
	@if [ ! -f $(CURDIR)/.env ]; then cp -f $(CURDIR)/$(TEMPLATE_ENV) $(CURDIR)/.env; fi

.PHONY: run
run:
	@docker-compose --env-file .env  up -d
