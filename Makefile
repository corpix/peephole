.DEFAULT_GOAL = all

numcpus  := $(shell cat /proc/cpuinfo | grep '^processor\s*:' | wc -l)
version  := $(shell git rev-list --count HEAD).$(shell git rev-parse --short HEAD)

name     := go-boilerplate
package  := github.com/corpix/$(name)

build    := ./build
build_id := 0x$(shell echo $(version) | sha1sum | awk '{print $$1}')
ldflags  := -X $(package)/cli.version=$(version) \
            -B $(build_id)

.PHONY: all
all:: dependencies

.PHONY: dependencies
dependencies::
	glide install

.PHONY: clean
clean::
	git clean -xddff

.PHONY: test
test:: dependencies
	go test -v \
           $(shell glide novendor)

.PHONY: bench
bench:: dependencies
	go test        \
           -bench=. -v \
           $(shell glide novendor)

.PHONY: lint
lint:: dependencies
	go vet $(shell glide novendor)

.PHONY: check
check:: lint test

.PHONY: all
all:: $(name)

.PHONY: $(name)
$(name):: dependencies
	mkdir -p $(build)
	@echo "Build id: $(build_id)"
	go build -a -ldflags "$(ldflags)" -v \
                 -o build/$(name)            \
                 $(package)/$(name)

.PHONY: build
build:: $(name)

.PHONY: clean
clean::
	rm -rf $(build)

include config.mk
