GOCMD:=go
GOBUILD:=$(GOCMD) build
GOTEST:=$(GOCMD) test
GOLINT:=$(GOCMD) vet
GOINSTALL:=$(GOCMD) install

BIN_DIR:=bin
BIN_NAME:=server
BIN_PATH:=$(BIN_DIR)/$(BIN_NAME)

BUILD_TAGS:=debug
BUILD_RACE:=-race
BUILD_STATIC:=

ifdef RELEASE
	BUILD_TAGS:=release
	BUILD_RACE:=
	BUILD_STATIC:=-a
endif

BUILD_OPTIONS:=-ldflags '-s -w' -tags $(BUILD_TAGS) $(BUILD_RACE) $(BUILD_STATIC)

SWAG_INIT:=
ifdef CI
	SWAG_INIT:=swag_init
endif


.PHONY: help build clean test lint swag swag_init

help:
	@echo Makefile Command Reference
	@echo Usage:
	@echo   make [TASK] [OPTION]...
	@echo Task:
	@echo   build [RELEASE=1]   program build
	@echo                       [RELEASE=1] release build
	@echo   clean               cleaning bin/ directory
	@echo   test                unit testing
	@echo   lint                lint
	@echo   swag [CI=1]         generate swagger api document
	@echo                       [CI=1]exec swag_init task before generate
	@echo   swag_init           [for CI]install swag command at latest version

build: $(BIN_PATH)

$(BIN_PATH):
	$(GOBUILD) -o $@ $(BUILD_OPTIONS) ./cmd/main.go

clean:
	rm -rf ./bin/*

test:
	$(GOTEST) -v ./...

lint:
	$(GOLINT) ./...

swag: $(SWAG_INIT)
	swag init \
		-o ./doc/api/ \
		-d ./cmd/,./interface/handler/ \
		--pd .\domain\ \
		--generatedTime

swag_init:
	$(GOINSTALL) github.com/swaggo/swag/cmd/swag@latest