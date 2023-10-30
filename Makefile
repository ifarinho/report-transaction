#!/usr/bin/env make

DOCKER_IMAGE=report-transaction
TEMPLATE_ENV=template.env

.PHONY: build
build:
	@docker build -t report-transaction .

.PHONY: env
env:
	@if [ ! -f $(CURDIR)/.env ]; then cp -f $(CURDIR)/$(TEMPLATE_ENV) $(CURDIR)/.env; fi

.PHONY: db
db:
	@docker run --name some-postgres \
 		-e POSTGRES_USER=postgres \
 		-e POSTGRES_PASSWORD=postgres \
 		-e POSTGRES_DB=postgres \
 		-p 5432:5432 \
 		-d postgres

.PHONY: run
run:
	@docker run -it report-transaction
