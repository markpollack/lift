SHELL = /usr/bin/env bash
OUTPUT = ./bin/lift
GO_SOURCES = $(shell find pkg cmd -type f -name "*.go")
GOBIN ?= $(shell go env GOPATH)/bin
VERSION ?= $(shell cat VERSION)
GOLINT = $(GOBIN)/golangci-lint
LIFT_PACKAGE = ./cmd/lift

export GO111MODULE := on

.PHONY: all
all: build verify-goimports lint test ## Build, test, verify source formatting and lint source

.PHONY: clean
clean: ## Delete build output
	rm -f $(OUTPUT)
	rm -f lift-darwin-amd64.tgz 
	rm -f lift-linux-amd64.tgz 
	rm -f lift-windows-amd64.zip 

.PHONY: bindir
bindir:
	mkdir -p ./bin

.PHONY: build
build: bindir ## Build lift
	go build -o bin/lift $(LIFT_PACKAGE)

.PHONY: test
test: ## Run the tests
	go test ./...

.PHONY: install
install: build ## Copy build to GOPATH/bin
	cp $(OUTPUT) $(GOBIN)

.PHONY: coverage
coverage: ## Run the tests with coverage and race detection
	go test -v --race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: check-linters-installed
check-linters-installed:
ifneq ($(OS),Windows_NT)
	@$(GOLINT) --version > /dev/null || (echo "golangci-lint not installed, run script ./hack/install-linters.sh"; exit 1)
endif

.PHONY: goimports
goimports: check-linters-installed ## Runs goimports on the project
	@$(GOLINT) run --no-config --disable-all --enable goimports --fix pkg/... cmd/...

.PHONY: verify-goimports
verify-goimports: check-linters-installed ## Verifies if all source files are formatted correctly
	@$(GOLINT) run --no-config --disable-all --enable goimports pkg/... cmd/...

.PHONY: lint
lint: check-linters-installed ## Runs golangci-lint tool. This will run multiple linting tools in a single command
	@$(GOLINT) run pkg/... cmd/...

.PHONY: release
release: bindir $(GO_SOURCES) VERSION ## Cross-compile lift for various operating systems
	GOOS=darwin   GOARCH=amd64 go build -o $(OUTPUT)     $(LIFT_PACKAGE) && tar -czf lift-darwin-amd64.tgz  $(OUTPUT)     && rm -f $(OUTPUT)
	GOOS=linux    GOARCH=amd64 go build -o $(OUTPUT)     $(LIFT_PACKAGE) && tar -czf lift-linux-amd64.tgz   $(OUTPUT)     && rm -f $(OUTPUT)
	GOOS=windows  GOARCH=amd64 go build -o $(OUTPUT).exe $(LIFT_PACKAGE) && zip -mq  lift-windows-amd64.zip $(OUTPUT).exe && rm -f $(OUTPUT).exe

help: ## Print help for each make target
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
