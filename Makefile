OSARCH := linux/386 linux/amd64 linux/arm linux/arm64 darwin/amd64 darwin/arm64

CLI_BUILD_VERSION ?= cli version not set
LDFLAGS := -X 'github.com/tensorleap/cli-go/pkg/version.CliVersion=$(CLI_BUILD_VERSION)'

.PHONY: build-cross
build-cross:
	echo $(LDFLAGS)
	gox \
		-output  "./dist/tensorleap-{{.OS}}-{{.Arch}}" \
		-osarch "$(OSARCH)" \
		-ldflags "$(LDFLAGS)" \
		.


