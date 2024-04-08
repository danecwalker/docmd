# Name of the binary to be built
BINARY=docmd

# Entry point of your Go application
ENTRYPOINT=./cmd/docmd

# Target platforms
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64 arm64

# Setup linker flags option for build that interoperate with variable names in your Go code
LDFLAGS=-ldflags "-X github.com/danecwalker/docmd/internal/meta.Version=$(VERSION) -X github.com/danecwalker/docmd/internal/meta.Revision=$(REVISION)"

# Build version
VERSION=${VERSION}

# Build Revision
REVISION=$(shell git rev-parse HEAD)

# Build for all platforms
build:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES),\
	$(shell GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o bin/$(BINARY)-$(GOOS)-$(GOARCH) $(ENTRYPOINT))))

# Clean up
clean:
	rm -f bin/*

.PHONY: build clean
