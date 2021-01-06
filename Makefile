GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

TARGET := mirage

.PHONY: default
default: run

.PHONY: build
build: fmt
	go build -o $TARGET main.go

.PHONY: build
run: fmt
	go run main.go

.PHONY: mod-tidy
mod-tidy:
	@go mod tidy
	@go mod download

.PHONY: generate
generate:
	@go generate mirc/main.go
	@$(GOFMT) ./

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
