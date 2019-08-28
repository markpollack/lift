SHELL := /bin/bash
OUTPUT = ./lift
GO_SOURCES = $(shell find . -type f -name '*.go')
GOBIN ?= $(shell go env GOPATH)/bin
VERSION ?= $(shell cat VERSION)

.PHONY: all
all: build test verify-goimports verify-gofmt ## Build, test, verify source formatting and regenerate doc

.PHONY: clean
clean: ## Delete build output
	rm -f $(OUTPUT)
	rm -f lift-darwin-amd64.tgz 
	rm -f lift-linux-amd64.tgz 
	rm -f lift-windows-amd64.zip 

.PHONY: build
build: $(OUTPUT) ## Build lift

.PHONY: test
test: ## Run the tests
	go test ./...

.PHONY: install
install: build ## Copy build to GOPATH/bin
	cp $(OUTPUT) $(GOBIN)

.PHONY: coverage
coverage: ## Run the tests with coverage and race detection
	go test -v --race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: check-goimports
check-goimports: ## Checks if goimports is installed
	@which goimports > /dev/null || (echo goimports not found: issue \"go get -u golang.org/x/tools/cmd/goimports\" && false)

.PHONY: goimports
goimports: check-goimports ## Runs goimports on the project
	@goimports -w cmd pkg

.PHONY: verify-goimports
verify-goimports: check-goimports ## Verifies if all source files are formatted correctly
	@goimports -l cmd pkg | (! grep .) || (echo above files are not formatted correctly. please run \"make goimports\" && false)

.PHONY: verify-gofmt
verify-gofmt: ## Check the file format
	@gofmt -e -d cmd pkg | read; \
		if [ $$? == 0 ]; then \
			echo "gofmt check failed:"; \
			gofmt -e -d cmd pkg; \
			exit 1; \
		fi

.PHONY: gofmt
gofmt: ## run gofmt tool on sources
	@gofmt -l -w $(GO_SOURCES)

$(OUTPUT): $(GO_SOURCES) VERSION
	go build -o $(OUTPUT) .

.PHONY: release
release: $(GO_SOURCES) VERSION ## Cross-compile lift for various operating systems
	GOOS=darwin   GOARCH=amd64 go build -o $(OUTPUT)     . && tar -czf lift-darwin-amd64.tgz  $(OUTPUT)     && rm -f $(OUTPUT)
	GOOS=linux    GOARCH=amd64 go build -o $(OUTPUT)     . && tar -czf lift-linux-amd64.tgz   $(OUTPUT)     && rm -f $(OUTPUT)
	GOOS=windows  GOARCH=amd64 go build -o $(OUTPUT).exe . && zip -mq  lift-windows-amd64.zip $(OUTPUT).exe && rm -f $(OUTPUT).exe

help: ## Print help for each make target
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
