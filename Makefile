APP=everything-template

.PHONY: build
build:
	@go build -o releases/${APP} ./cmd/app

.PHONY: windows
windows:
	@GOARCH=amd64 GOOS=windows go build -ldflags="-s" -o releases/${APP}-win ./cmd/app

.PHONY: linux
linux:
	@GOARCH=amd64 GOOS=linux go build -ldflags="-s" -o releases/${APP}-linux ./cmd/app

.PHONY: darwin
darwin:
	@GOARCH=amd64 GOOS=darwin go build -ldflags="-s" -o releases/${APP}-darwin ./cmd/app

.PHONY: lint
lint:
	@if ! command -v gofumpt &> /dev/null; then \
		echo "gofumpt not found, installing..."; \
		go install mvdan.cc/gofumpt@latest; \
	fi
	@gofumpt -l -w .

.PHONY: generate
generate:
	@go generate -x

.PHONY: clean
clean:
	@go clean -i .
	@rm -rf releases

.PHONY: help
help:
	@echo "1. make build - [go build]"
	@echo "2. make windows - [make window package]"
	@echo "3. make linux - [make linux package]"
	@echo "4. make darwin - [make darwin package]"
	@echo "5. make lint - [gofumpt -l -w .]"
	@echo "6. make generate - [go generate -x]"
	@echo "7. make clean - [remove releases files and cached files]"