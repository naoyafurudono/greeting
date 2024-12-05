# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := run
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory
BIN=$(abspath .tmp/bin)
export PATH := $(BIN):$(PATH)
export GOBIN := $(abspath $(BIN))
COPYRIGHT_YEARS := 2022-2023
LICENSE_IGNORE := --ignore /testdata/
# Set to use a different compiler. For example, `GO=go1.18rc1 make test`.
GO ?= go

.PHONY: help
help: ## Describe useful make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Delete intermediate build artifacts
	@# -X only removes untracked files, -d recurses into directories, -f actually removes files/dirs
	git clean -Xdf

.PHONY: run
run: build ## Execute the command `make run ARGS="hello --name naoya"` as `greet hello --name naoya` (default)
	$(BIN)/main $(ARGS)
	
.PHONY: /build
build: $(BIN)/build ## Build all packages

.PHONY: generate
generate: $(BIN)/buf $(BIN)/protoc-gen-go $(BIN)/protoc-gen-connect-go ## Regenerate code and licenses
	rm -rf gen
	PATH=$(BIN) $(BIN)/buf generate
	
$(BIN)/build: generate
	$(GO) build -o $(@D)/main ./cmd/server/main.go

$(BIN)/buf: Makefile
	@mkdir -p $(@D)
	$(GO) install github.com/bufbuild/buf/cmd/buf@latest

$(BIN)/protoc-gen-go: Makefile
	@mkdir -p $(@D)
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go

$(BIN)/protoc-gen-connect-go: Makefile go.mod
	@mkdir -p $(@D)
	$(GO) install connectrpc.com/connect/cmd/protoc-gen-connect-go
