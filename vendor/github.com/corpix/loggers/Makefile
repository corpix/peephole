.DEFAULT_GOAL = all

numcpus  := $(shell cat /proc/cpuinfo | grep '^processor\s*:' | wc -l)
version  := $(shell git rev-list --count HEAD).$(shell git rev-parse --short HEAD)

name     := loggers
package  := github.com/corpix/$(name)

# XXX: Fuck you golang!
# 99% of time having vendor in a wildcard result is not what you want!
packages := $(shell go list ./... | grep -v /vendor/)

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

.PHONY: clean
clean::
	git clean -xddff
