## Reference makefile
BINARY_NAME=main.out

all: build test

.PHONY: build
build:
	go build -o ${BINARY_NAME} main.go

.PHONY: test
test:
	go test -v

.PHONY: exporttest
exporttest:
	GO111MODULE=on go test -v -race -coverprofile=cover.out -covermode=atomic ./...
	go tool cover -html=cover.out -o cover.htm

.PHONY: run
run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

.PHONY: clean
clean:
	go clean
	rm ${BINARY_NAME}

.PHONY: install
install:
	make build
	mkdir -p ~/.local/bin
	cp ${BINARY_NAME} ~/.local/bin
