GO = $(shell which go)
GOBIN = $(shell $(GO) env GOPATH)/bin
BINPATH = $(CURDIR)/bin

PLATFORMS := linux/386 linux/amd64 darwin/386 darwin/amd64
PLATFORM = $(subst /, ,$@)
OS = $(word 1, $(PLATFORM))
ARCH = $(word 2, $(PLATFORM))

BINNAME = goradiooooo
CMDPATH = main.go
VERSION = 0.1.4

LD_FLAGS = "-w -s -X main.Version=$(VERSION)"

UPXBIN = $(shell command -v upx-ucl 2>/dev/null)

.PHONY: deps_lint deps lint build build_all clean
.DEFAULT_GOAL := build

deps_lint:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(GOBIN) v1.21.0

deps: deps_lint
	@$(GO) get -v -t -d  ./...

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
	$(eval OUTBIN := "$(BINPATH)/$(BINNAME)_$(OS)_$(ARCH)")
	@GOOS=$(OS) GOARCH=$(ARCH) $(GO) build \
		-o $(OUTBIN) \
		-ldflags=$(LD_FLAGS) \
		$(CMDPATH)
ifneq ($(UPXBIN),)
	@echo compress...
	@$(UPXBIN) -qqf $(OUTBIN)
endif

build_all: $(PLATFORMS)

clean:
	@rm -rf $(BINPATH)/*
