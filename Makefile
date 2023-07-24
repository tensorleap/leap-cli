OSARCH := linux/386 linux/amd64 linux/arm linux/arm64 darwin/amd64 darwin/arm64

CLI_BUILD_VERSION ?= cli version not set

# See available flags by running: go tool link
LDFLAGS := -w -s -X 'github.com/tensorleap/cli-go/pkg/version.CliVersion=$(CLI_BUILD_VERSION)'

export CGO_ENABLED=0

.PHONY: build-cross
build-cross:
	gox \
		-output  "./dist/leap-{{.OS}}-{{.Arch}}" \
		-osarch "$(OSARCH)" \
		-ldflags "$(LDFLAGS)" \
		.

.PHONY: docgen
docgen:
	rm -rf docs/*
	go run docgen/docgen.go

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: check-fmt
check-fmt:
	@if test -z "$$(go fmt ./... 2>&1)" ; then \
		echo "No formatting issues found."; \
	else \
		echo "[ERROR] Fix formatting issues with 'make fmt'"; \
		exit 1; \
	fi

.PHONY: test
test: 
	@go test ./...