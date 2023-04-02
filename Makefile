export GO111MODULE=on

APP:=hive-backend
OS:=$(shell go env GOOS)
ARCH:=$(shell go env GOARCH)
LOCAL_BIN:=$(CURDIR)/bin
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.52.2
GOLANGCI_LINT_VERSION:=v$(GOLANGCI_TAG)
GOLANGCI_CONFIG:=.golangci.yaml
GOLANGCI_STRICT_CONFIG:=.golangci-strict.yaml

# Check local bin version
ifneq ($(wildcard $(GOLANGCI_BIN)),)
GOLANGCI_BIN_VERSION:=$(shell $(GOLANGCI_BIN) --version)
ifneq ($(GOLANGCI_BIN_VERSION),)
GOLANGCI_BIN_VERSION_SHORT:=$(shell echo "$(GOLANGCI_BIN_VERSION)" | sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
GOLANGCI_BIN_VERSION_SHORT:=0
endif
ifneq "$(GOLANGCI_TAG)" "$(word 1, $(sort $(GOLANGCI_TAG) $(GOLANGCI_BIN_VERSION_SHORT)))"
GOLANGCI_BIN:=
endif
endif

default: help

.PHONY: install-lint
install-lint: ## install golangci-lint binary
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

.PHONY: lint
lint: install-lint
ifeq ($(filter strict,$(MAKECMDGOALS)),strict)
	@echo "Running lint in strict mode..."
	$(GOLANGCI_BIN) run --new-from-rev=origin/master --config=$(GOLANGCI_STRICT_CONFIG) ./...
else
	@echo "Running lint in normal mode..."
	$(GOLANGCI_BIN) run --new-from-rev=origin/master --config=$(GOLANGCI_CONFIG) ./...
endif

.PHONY: lint-full
lint-full: install-lint
ifeq ($(filter strict,$(MAKECMDGOALS)),strict)
	@echo "Running lint-full in strict mode..."
	$(GOLANGCI_BIN) run --config=$(GOLANGCI_STRICT_CONFIG) ./...
else
	@echo "Running lint-full in normal mode..."
	$(GOLANGCI_BIN) run --config=$(GOLANGCI_CONFIG) ./...
endif

.PHONY: test
test:
	@go test -v ./...

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  lint                    Run golangci-lint with normal checks and compare changes against master branch."
	@echo "  lint strict             Same as 'lint', but with more strict checks."
	@echo "  lint-full               Run golangci-lint with normal checks for all files in the repository."
	@echo "  lint-full strict        Same as 'lint-full', but with more strict checks."
	@echo "  test                    Run tests"
	@echo "  help                    Show this help message"