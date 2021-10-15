BINARY_NAME=lite-api-crud
.DEFAULT_GOAL := all
GOPATH=$(shell cd ../ && pwd)

all: install test build

.PHONY: build test run clean install help

build:
	export GOPATH=${GOPATH} && \
	go build -ldflags="-s -w" -o ${BINARY_NAME} main.go

test:
	export ENV="test" && \
	export GOPATH=${GOPATH} && \
	go test -timeout 30s lite-api-crud/tests
	unset ENV

run:
	make build
	./${BINARY_NAME}

clean:
	export GOPATH=${GOPATH} && \
	go clean
	rm ${BINARY_NAME}

install: ## Install dependencies
	export GOPATH=${GOPATH} && \
	go install

help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
