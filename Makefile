GOCMD:=go
GORUN:=$(GOCMD) run
GOBUILD:=$(GOCMD) build
GOTEST:=$(GOCMD) test
GOLINT:=$(GOCMD) vet
GOTOOL:=$(GOCMD) tool
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

TEST_TYPE:=local
ifdef INTEGRATION
	TEST_TYPE:=integration
endif
ifdef UNIT
	TEST_TYPE:=unit
endif

.PHONY: help build clean test

help:
	@echo Makefile Command Reference
	@echo Usage:
	@echo   make [TASK] [OPTION]...
	@echo Task:
	@echo   help                        print this view
	@echo   build [RELEASE=1]           program build
	@echo                               [RELEASE=1] release build
	@echo   clean                       cleaning bin/ directory
	@echo   test [INTEGRATION=1|UNIT=1] testing and generate test coverage html
	@echo                               no option mode local
	@echo                               [INTEGRATION=1] run to integration test only
	@echo                               [UNIT=1] run to unit test only
	@echo   lint                        lint
	@echo   swag [CI=1]                 generate swagger api document
	@echo                               [CI=1]exec swag_init task before generate
	@echo   swag_clean                  cleaning doc/api/ directory using git command
	@echo   swag_init                   for CI - install swag command at latest version

build: $(BIN_PATH)

$(BIN_PATH):
	$(GOBUILD) -o $@ $(BUILD_OPTIONS) ./cmd/main.go

clean:
	rm -rf ./bin/*

test: test_setup $(TEST_TYPE) test_clean

test_setup:
	$(GORUN) ./test/setup

test_clean:
	$(GORUN) ./test/clean

local:
	$(GOTEST) -v -cover ./... -coverprofile=cover.out
	$(GOTOOL) cover -html cover.out -o cover.html

unit:
	$(GOTEST) -v -cover -tags=$@ ./... -coverprofile=cover.txt

integration:
	$(GOTEST) -v -cover -tags=$@ ./...

lint:
	$(GOLINT) ./...

swag: $(SWAG_INIT)
	swag init \
		-o ./doc/api/ \
		-d ./cmd/,./interface/handler/ \
		--pd ./domain/ ./usecase/inputdata/ \
		--generatedTime

swag_clean:
	git checkout ./doc/api/

swag_init:
	$(GOINSTALL) github.com/swaggo/swag/cmd/swag@v1.8.4