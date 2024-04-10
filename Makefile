OSARCH := linux/386 linux/amd64 linux/arm linux/arm64 darwin/amd64 darwin/arm64
CLI_BUILD_VERSION ?= cli version not set

# set variables for updating server api
NODE_SERVER_BRANCH ?= master
NODE_SERVER_BUILDER_IMAGE ?= ""

# See available flags by running: go tool link
LDFLAGS := -w -s -X 'github.com/tensorleap/leap-cli/pkg/version.CliVersion=$(CLI_BUILD_VERSION)'

export CGO_ENABLED=0

.PHONY: build-cross
build-cross:
	gox \
		-output  "./dist/leap-{{.OS}}-{{.Arch}}" \
		-osarch "$(OSARCH)" \
		-ldflags "$(LDFLAGS)" \
		.

.PHONY: build-local
build-local:
	go build .

.PHONY: docgen
docgen:
	rm -rf docs/*
	go run docgen/docgen.go

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: fmt
fmt:
	@gofmt -w -l ./

.PHONY: check-fmt
check-fmt:
	@echo "Checking code formatting..."
	@result=$$(gofmt -l ./); \
	if [ -n "$$result" ]; then \
		echo "Formatting issues found:"; \
		echo "$$result"; \
		exit 1; \
	fi

.PHONY: test
test: 
	@go test ./...

.PHONY: update-server-api
update-server-api:
	@./scripts/update_server_api.sh ${NODE_SERVER_BUILDER_IMAGE} ${NODE_SERVER_BRANCH} 

.PHONY: build-shellcheck-image
build-shellcheck-image:
	docker build -t shellcheck-image -f shellcheck.Dockerfile .

.PHONY: shellcheck
shellcheck: build-shellcheck-image
	docker run --rm -v `pwd`:/repo shellcheck-image
