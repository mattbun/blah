.PHONY: all
all: tidy fmt build

.PHONY: build
build:
	go build -o bin/blah main.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt
