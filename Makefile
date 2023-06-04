OSARCH := linux/386 linux/amd64 linux/arm linux/arm64 darwin/amd64 darwin/arm64

.PHONY: build-cross
build-multiarch:
	gox -output  "./dist/tensorleap-{{.OS}}-{{.Arch}}" -osarch "$(OSARCH)" .


