

_GIT_VERSION := $(shell git describe --tags --dirty)
_GIT_COMMIT := $(shell git rev-parse --short HEAD)
GO_BUILD_FLAGS ?= -ldflags "-X main.Version=$(_GIT_VERSION) -X main.BuildCommit=$(_GIT_COMMIT)"
GO_BIN ?= go

dist/tmpl:
	$(GO_BIN) build $(GO_BUILD_FLAGS) -o dist/tmpl cmd/main.go

.PHONY: clean
clean:
	if [ -f dist/tmpl ]; then rm -f dist/tmpl; fi

.PHONY: build
build: clean dist/tmpl

.PHONY: test
test:
	$(GO_BIN) test -v $$($(GO_BIN) list ./... | grep -v /example) && echo && echo "ALL TESTS OK"