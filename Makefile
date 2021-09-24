FILES ?= $(shell find . -type f -name '*.go' ! -path './_volumes/*' ! -path './vendor/*' )
PACKAGES ?= $(shell go list ./...)

.PHONY: all format tools

all: format test

tools:
	go install golang.org/x/tools/cmd/goimports

format:
	@ goimports -w $(FILES)
	@ go mod tidy
	@ go fmt ./...

TEST_FOLDERS = knowlib/file_controller/test_storage

test:
	go test ./...
	rm -rf $(TEST_FOLDERS)
