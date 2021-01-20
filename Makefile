GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

TARGET := mirage
TAGS   := jsoniter

.PHONY: default
default: run

.PHONY: build
build: fmt
	go build -tags '$(TAGS)' -o $(TARGET) main.go

.PHONY: build
run:
	@go run -tags '$(TAGS)' main.go -debug

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
