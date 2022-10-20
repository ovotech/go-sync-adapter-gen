SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
TARGET ?= .

lint/golangci_lint: ## Lint using golangci-lint.
	golangci-lint run $(TARGET)/...

lint: ## Lint Go Sync.
	make lint/golangci_lint
.PHONY: lint lint/*

fix/gofmt: ## Fix formatting with gofmt.
	gofmt -w $(TARGET)

fix/gci: ## Fix imports.
	gci write $(TARGET)

fix/golangci_lint: ## Fix golangci-lint errors.
	golangci-lint run --fix $(TARGET)/...

fix: ## Fix common linter errors.
	make fix/gofmt
	make fix/gci
	make fix/golangci_lint
.PHONY: fix fix/*

generate/fixtures: ## Regenerate fixtures for use with testing.
	rm -f test/snapshots/*.txt
	go run test/snapshots.go

generate: ## Generate automated code.
	make generate/fixtures

.DEFAULT_GOAL := help
help: Makefile ## Display list of available commands.
	@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'
