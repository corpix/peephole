.DEFAULT_GOAL = all

numcpus  := $(shell cat /proc/cpuinfo | grep '^processor\s*:' | wc -l)
version  := $(shell git rev-list --count HEAD).$(shell git rev-parse --short HEAD)

name     := go-boilerplate
package  := github.com/corpix/$(name)

# XXX: Fuck you golang!
# 99% of time having vendor in a wildcard result is not what you want!
packages := $(shell go list ./... | grep -v /vendor/)

build       := ./build
build_id    := 0x$(shell echo $(version) | sha1sum | awk '{print $$1}')
ldflags     := -X $(package)/cli.version=$(version) -B $(build_id)
build_flags := -a -ldflags "$(ldflags)" -o build/$(name)

.PHONY: all
all:: dependencies

.PHONY: dependencies
dependencies::
	dep ensure

.PHONY: test
test:: dependencies
	go test -v $(packages)

.PHONY: bench
bench:: dependencies
	go test -bench=. -v $(packages)

.PHONY: lint
lint:: dependencies
	go vet -v $(packages)

.PHONY: check
check:: lint test

.PHONY: all
all:: $(name)

.PHONY: $(name)
$(name):: dependencies
	mkdir -p $(build)
	@echo "Build id: $(build_id)"
	go build $(build_flags) -v $(package)/$(name)

.PHONY: build
build:: $(name)

.PHONY: clean
clean::
	git clean -xddff

include config.mk
