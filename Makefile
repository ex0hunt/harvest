GOARCH ?=
GO111MODULE ?= on
GOPATH ?= $(CURDIR)

PACKAGES = $(shell $(GO) list -m all)
GO := GOPATH=$(GOPATH) GO111MODULE=$(GO111MODULE) go

.DEFAULT_GOAL := build

.PHONY: show_env
show_env:
	@echo ">> show env"
	@echo "   GOROOT            = $(GOROOT)"
	@echo "   GOPATH            = $(GOPATH)"
	@echo "   GO111MODULE       = $(GO111MODULE)"
	@echo "   VERSION           = $(VERSION)"
	@echo "   PACKAGES          = $(PACKAGES)"

.PHONY: build
build: show_env
	@echo ">> build server"
	$(GO) build -o ./bin/index_server ./src/
	@echo ">> done"

.PHONY: clean
clean:
	@echo ">> cleanup"
	rm -rf ./bin/
	@echo ">> done"