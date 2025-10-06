SHELL := /usr/bin/env bash
.SHELLFLAGS += -o pipefail -O extglob
.DEFAULT_GOAL := test

GOLANGCI_LINT_VERSION := v2.4.0

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## Tool Binaries
GO_RUN := go run
GOLANGCI_LINT ?= $(GO_RUN) github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)


.PHONY: lint
lint: ## Run golangci-lint linter
	$(GOLANGCI_LINT) run

.PHONY: lint-fix
lint-fix: ## Run golangci-lint linter and perform fixes
	$(GOLANGCI_LINT) run --fix

.PHONY: lint-config
lint-config: ## Verify golangci-lint linter configuration
	$(GOLANGCI_LINT) config verify


##@ Test targets

.PHONY: test
test: ## run tests
	go test -v -race -count=1 ./...

	LOG_FORMAT=json go test -v -race -count=1 ./...
	LOG_FORMAT=text go test -v -race -count=1 ./...

	LOG_LEVEL=debug go test -v -race -count=1 ./...
	LOG_LEVEL=info go test -v -race -count=1 ./...
	LOG_LEVEL=warn go test -v -race -count=1 ./...
	LOG_LEVEL=error go test -v -race -count=1 ./...
	LOG_LEVEL=fatal go test -v -race -count=1 ./...

	LOG_TIME_FORMAT=epoch go test -v -race -count=1 ./...
	LOG_TIME_FORMAT=epoch_ms go test -v -race -count=1 ./...
	LOG_TIME_FORMAT=rfc3339 go test -v -race -count=1 ./...
	LOG_TIME_FORMAT=iso8601 go test -v -race -count=1 ./...
