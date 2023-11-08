filename ?=

.PHONY: lint
lint:
	gofmt -s -w .

.PHONY: build
build:
	mkdir -p bin && go build -o bin/ ./cmd/...

.PHONY: run-cli
run_cli: build
	bin/cli -filename=$(filename)

.PHONY: help
help: build
	bin/cli -h