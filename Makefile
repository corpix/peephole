.DEFAULT_GOAL := all

export GOPROXY ?= https://proxy.golang.org

## parameters

NAME              ?= peephole
NAMESPACE         ?= corpix
VERSION           ?= development
ENV               ?= dev

IMAGE_TAR   = container.tar.gz
IMAGE_NAME ?= $(NAME)
IMAGE_TAG  ?= latest

PARALLEL_JOBS   ?= 8
NIX_BUILD_CORES ?= $(PARALLEL_JOBS)
NIX_OPTS        ?=

## bindings

root                := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
profile_dir         := ./profile
build_dir           := ./build
nix_dir             := ./nix
tmux                := tmux -2 -f $(PWD)/.tmux.conf -S $(PWD)/.tmux
tmux_session        := $(NAMESPACE)/$(NAME)
pkg_prefix          := github.com/$(NAMESPACE)/$(NAME)
nix                 := nix $(NIX_OPTS)
shell_volume_nix    := nix

cmds                := peephole
ports               := 1338

### reusable and long opts for commands inside rules

add_shell_opts ?=
shell_opts = -v $(shell_volume_nix):/nix:rw     \
	-v $(root):/chroot                      \
	-e COLUMNS=$(COLUMNS)                   \
	-e LINES=$(LINES)                       \
	-e TERM=$(TERM)                         \
	-e NIX_BUILD_CORES=$(NIX_BUILD_CORES)   \
	-e HOME=/chroot                         \
	-w /chroot                              \
	--hostname $(NAMESPACE).localhost       \
	$(foreach v,$(ports), -p $(v):$(v) ) $(add_shell_opts)

lint_opts := --color=always                                                        \
	--exclude='uses unkeyed fields'                                            \
	--exclude='type .* is unused'                                              \
	--exclude='should merge variable declaration with assignment on next line' \
	--deadline=120s                                                            \

## helpers

, = ,

## macro

# XXX: yes, this two empty lines here are required :)
define \n


endef

define sorted
@echo $(cmds) | sed 's|\s|\n|g' | sort
endef

define fail
{ echo "error: "$(1) 1>&2; exit 1; }
endef

define expect
{ grep $(1) > /dev/null || $(call fail,$(2)); }
endef

define required
@if [ -z $(2) ]; then $(call fail,"$(1) is required") fi
endef

####

define test_run
bash --noprofile -euxo pipefail -c "$(3) go test $(2) $(1) | grep -vF '[no test files]' | column -t"
endef

define lint_run
golangci-lint $(lint_opts) run $(1)
endef

define docker_run
docker run --rm -it --log-driver=none $(1)
endef

### releases

### development

build/%: # build specified app
	$(info == building $*)
	@mkdir -p $(build_dir)
	go build -o $@ --ldflags                          \
		"-X $(pkg_prefix)/cli.version=$(VERSION)" \
		./$*/$*.go

.PHONY: build
build: $(patsubst %,build/%,$(cmds)) # build all cmds

.PHONY: build/nix
build/nix: # build all cmds with nix
	nix-build -A all

.PHONY: repl
repl: # run nix repl
	nix repl '<nixpkgs>'

.PHONY: tidy
tidy: fmt # format code and tidy dependencies
	go mod $@

.PHONY: fmt
fmt: # go fmt repository
	go fmt ./...

config.json: config.nix # generate config.json from config.nix
	nix-instantiate --eval --expr 'builtins.toJSON (import ./$<)' | jq -r . | jq . > $@

.PHONY: list/cmd
list/cmd: # list all project cmds
	$(call sorted,$(cmds))

#### testing

.PHONY: test-race
test-race: # race-condition test on whole repository
	$(call test_run,./...,-race -short,CC=clang)

.PHONY: test-msan
test-msan: # memory sanitizer (memory leak) test on whole repository
	$(call test_run,./...,-msan -short,CC=clang)

test/%: # test specified package
	$(call test_run,./$*/...)

.PHONY: test
test: # test whole repository
	$(call test_run,./...)

lint/%: # lint specified package
	$(call lint_run,./$*/...)

.PHONY: lint
lint: # lint whole repository
	$(call lint_run,./...)

profile/%: # profile specified cmd
	$(info == profiling $*)
	@mkdir -p $(profile_dir)/$*
	go run ./cmd/$*/main.go --profile $(profile_dir)/$*
	go run ./cmd/$*/main.go --trace   $(profile_dir)/$*

.PHONY: profile
profile: $(patsubst %,profile/%,$(cmds)) # profile all cmds

#### environment management

.PHONY: dev/clean
dev/clean: # clean development environment artifacts
	docker volume rm nix

.PHONY: dev/shell
dev/shell: # run development environment shell
	@docker run --rm -it                   \
		--log-driver=none              \
		$(shell_opts) nixos/nix:latest \
		nix-shell --command "exec make dev/start-session"

.PHONY: dev/shell/raw
dev/shell/raw: # run development environment shell
	@docker run --rm -it                   \
		--log-driver=none              \
		$(shell_opts) nixos/nix:latest \
		nix-shell

.PHONY: dev/session
dev/start-session: # start development environment terminals with database, blockchain, etc... one window per app
	@$(tmux) has-session    -t $(tmux_session) && $(call fail,tmux session $(tmux_session) already exists$(,) use: '$(tmux) attach-session -t $(tmux_session)' to attach) || true
	@$(tmux) new-session    -s $(tmux_session) -n console -d
	@$(tmux) send-keys      -t $(tmux_session):0 C-z 'emacs .' Enter
	@$(tmux) select-window  -t $(tmux_session):0

	@if [ -f $(root)/.personal.tmux.conf ]; then $(tmux) source-file $(root)/.personal.tmux.conf; fi

	@$(tmux) attach-session -t $(tmux_session)

.PHONY: dev/attach-session
dev/attach-session: # attach to development session if running
	@$(tmux) attach-session -t $(tmux_session)

.PHONY: dev/stop-session
dev/stop-session: # stop development environment terminals
	@$(tmux) kill-session -t $(tmux_session)

.PHONY: run/statsd
run/statsd: # run statsd server with docker
	$(call docker_run,--net=host atlassianlabs/gostatsd --backends stdout)

##

.PHONY: container
container: container.nix # build container image archive which could be loaded into docker
	nix-build --show-trace --expr 'import ./$< { name = "$(IMAGE_NAME)"; tag = "$(IMAGE_TAG)"; }'
	cp -Lf --preserve=timestamps result ./$(IMAGE_TAR)
	rm -f result

.PHONY: container-load
container-load: # load container image archive into docker
	docker load < $(IMAGE_TAR)

.PHONY: clean
clean: # clean stored state
	rm -rf $(build_dir)
	rm -rf $(profile_dir)

.PHONY: help
help: # print defined targets and their comments
	@grep -Po '^[a-zA-Z%_/\-\s]+:+(\s.*$$|$$)' $(MAKEFILE_LIST)      \
		| sort                                                   \
		| sed 's|:.*#|#|;s|#\s*|#|'                              \
		| column -t -s '#' -o ' | '
