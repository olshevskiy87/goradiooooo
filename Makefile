GO = $(shell which go)
GOBIN = $(shell $(GO) env GOPATH)/bin
BINPATH = $(CURDIR)/bin

PLATFORMS := linux/386 linux/amd64
PLATFORM = $(subst /, ,$@)
OS = $(word 1, $(PLATFORM))
ARCH = $(word 2, $(PLATFORM))

BINNAME = goradiooooo
CMDPATH = main.go
VERSION = 0.1.1

LD_FLAGS = "-w -s -X main.Version=$(VERSION)"

.PHONY: deps_lint deps lint build build_all clean
.DEFAULT_GOAL := build

deps_lint:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(GOBIN) v1.21.0

deps: deps_lint
	@$(GO) get ./...

lint:
	@golangci-lint run

build:
	@echo building $(BINNAME)...
	@$(GO) build \
		-o $(BINPATH)/$(BINNAME) \
		-ldflags=$(LD_FLAGS) \
		$(CMDPATH)

$(PLATFORMS):
	@echo building $(BINNAME) for $(OS)/$(ARCH)...
	@GOOS=$(OS) GOARCH=$(ARCH) $(GO) build \
		-o $(BINPATH)/$(BINNAME)_$(OS)_$(ARCH) \
		-ldflags=$(LD_FLAGS) \
		$(CMDPATH)

build_all: $(PLATFORMS)

clean:
	@rm -f $(BINPATH)/*
