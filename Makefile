.PHONY: install-tools lint test test-verbose format help compile-constants
.SILENT: install-tools lint test test-verbose format help compile-constants

help:
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install-tools: ## Install linting tools
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	go install mvdan.cc/gofumpt@latest
	go install github.com/segmentio/golines@latest

lint: ## Run golangci linter
	golangci-lint run -c ./golangci.yml ./...

format: ## Format code
	gofumpt -l -w -extra .
	golines . -w

test: ## Run tests
	go test -test.timeout 120s -count=1 .
	go test -test.timeout 120s -count=1 ./ciinfo/

test-verbose: ## Run tests with verbose output
	go test -test.timeout 120s -v -cover -count=1 .
	go test -test.timeout 120s -v -cover -count=1 ./ciinfo/

compile-constants: ## Generates 'constants.go' containing the list with constant values
	cp vendors.go compile-constants
	sed -i 's/ciinfo/main/' compile-constants/vendors.go
	go run ./compile-constants/main.go ./compile-constants/vendors.go
