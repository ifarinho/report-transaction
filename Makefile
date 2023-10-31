#!/usr/bin/env make

PROJECT_NAME=report-transaction
TEMPLATE_ENV=template.env

.PHONY: help
help:
	@echo "Command list: \n- help: display this message \n- build: compiles the program using go build \n- build-image: compiles the program and builds the image \n- dotenv: creates .env file for local use \n- csv: creates a random csv file \n- run: runs the program \n- local: runs the program with the generated .env file"

.PHONY: build
build:
	@go build -o $(PROJECT_NAME) ./cmd/app/main.go

.PHONY: build-image
build-image:
	@docker build -t $(PROJECT_NAME) .

.PHONY: run
run:
	@docker run --rm -it $(PROJECT_NAME) $(FILENAME) $(ACCOUNT)

.PHONY: dotenv
dotenv:
	@if [ ! -f $(CURDIR)/.env ]; then cp -f $(CURDIR)/$(TEMPLATE_ENV) $(CURDIR)/.env; fi

.PHONY: csv
csv:
	@python3 scripts/generate_csv.py

.PHONY: local
local:
	@./scripts/run_local.sh $(FILENAME) $(ACCOUNT)
