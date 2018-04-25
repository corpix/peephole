.DEFAULT_GOAL = all

version  := $(shell git rev-list --count HEAD).$(shell git rev-parse --short HEAD)

name     := peephole
package  := github.com/corpix/$(name)
packages := $(shell go list ./... | grep -v /vendor/)

build       := ./build
build_id    := 0x$(shell echo $(version) | sha1sum | awk '{print $$1}')
ldflags     := -X $(package)/cli.version=$(version) -B $(build_id)
build_flags := -a -ldflags "$(ldflags)" -o build/$(name)

.PHONY: all
all:: dependencies
all:: build

.PHONY: dependencies
dependencies::
	dep ensure

.PHONY: test
test::
	go test -v $(packages)

.PHONY: bench
bench::
	go test -bench=. -v $(packages)

.PHONY: lint
lint::
	go vet -v $(packages)

.PHONY: check
check:: lint test

.PHONY: all
all:: $(name)

.PHONY: $(name)
$(name)::
	mkdir -p $(build)
	@echo "Build id: $(build_id)"
	go build $(build_flags) -v $(package)/$(name)

.PHONY: build
build:: $(name)

.PHONY: clean
clean::
	git clean -xddff

include config.mk
include debug.mk
