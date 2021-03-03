

_GIT_VERSION := $(shell git describe --tags --dirty)
_GIT_COMMIT := $(shell git rev-parse --short HEAD)
_GO_BUILD_FLAGS = -ldflags "-X main.Version=$(_GIT_VERSION)"
GO_BIN ?= go
DOCKER_BIN ?= docker
DOCKER_TAG ?= tmpl:$(_GIT_COMMIT)
DOCKER_NAME ?= tmpl
DIST ?= dist/tmpl

$(DIST):
	$(GO_BIN) build $(GO_BUILD_FLAGS) $(_GO_BUILD_FLAGS) -o $(DIST) cmd/main.go

.PHONY: clean
clean:
	if [ -f $(DIST) ]; then rm -f $(DIST); fi

.PHONY: build
build: clean $(DIST)

.PHONY: test
test:
	$(GO_BIN) test -v $$($(GO_BIN) list ./... | grep -v /example) && echo && echo "ALL TESTS OK"

.PHONY: docker-build
docker-build:
	$(DOCKER_BIN) build -t $(DOCKER_TAG) .

.PHONY: docker-run
docker-run:
	@echo Execute:
	@echo $(DOCKER_BIN) run $(DOCKER_TAG)
	@echo " for example:"
	@echo "  $(DOCKER_BIN) run --rm --name $(DOCKER_NAME) $(DOCKER_TAG) --help"
	@echo "  $(DOCKER_BIN) run $(DOCKER_TAG) "