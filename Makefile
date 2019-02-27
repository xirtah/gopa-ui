SHELL=/bin/bash

# Default GOPA version
GOPA_VERSION := 0.11.0_SNAPSHOT

# Get release version from environment
ifneq "$(VERSION)" ""
   GOPA_VERSION := $(VERSION)
endif

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif

#Add the GOPATH bin directory to the path file
export PATH := $(PATH):$(GOPATH)/bin

# Go environment
CURDIR := $(shell pwd)

GO        := GO15VENDOREXPERIMENT="1" go
GOBUILD  := GOPATH=$(GOPATH) CGO_ENABLED=1  $(GO) build -ldflags -s
GOTEST   := GOPATH=$(GOPATH) CGO_ENABLED=1  $(GO) test -ldflags -s

.PHONY: all build update test clean

default: build

build: config
	$(GOBUILD) -o bin/gopa-ui

clean_data:
	rm -rif dist
	rm -rif data
	rm -rif log

clean: clean_data
	rm -rif bin
	mkdir bin

config: init-version update-ui update-template-ui
	@echo "update configs"
	@# $(GO) env
	@mkdir -p bin
	@cp gopa.yml bin/gopa.yml

init-version:
	@echo building GOPA-UI $(GOPA_VERSION)

update-ui:
	@echo "generate static files"
	@$(GO) get github.com/xirtah/esc
	@(cd static&& esc -ignore="static.go|build_static.sh|.DS_Store" -o static.go -pkg static ../static )

update-template-ui:
	@echo "generate UI pages"
	@$(GO) get github.com/xirtah/ego/cmd/ego
	@cd modules/ && ego
	#@cd plugins/ && ego

test:
	$(GOTEST) -timeout 60s ./...
