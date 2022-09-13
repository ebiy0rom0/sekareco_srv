# Enforce powershell to unify execution commands
GOOS:=$(shell go env GOOS)
ifeq ($(GOOS), windows)
	SHELL:=powershell.exe
endif

GOCMD:=go
GORUN:=$(GOCMD) run
GOBUILD:=$(GOCMD) build
GOTEST:=$(GOCMD) test
GOLINT:=$(GOCMD) vet
GOTOOL:=$(GOCMD) tool
GOINSTALL:=$(GOCMD) install

BIN_DIR:=./bin
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

# `local` task is convert the output coverage file to html
# Skipping `local` task by makeing TEST_LOCATE = TEST_MODE
TEST_MODE:=unit
TEST_LOCATE:=local
ifdef INTEGRATION
	TEST_MODE:=integration
endif
ifndef LOCAL
	TEST_LOCATE:=$(TEST_MODE)
endif

.PHONY: help build clean test

help:
	@echo Makefile Command Reference
	@echo Usage:
	@echo   make [TASK] [OPTION]...
	@echo Task List:
	@echo   help                           Print this view
	@echo   build [RELEASE=1]              Build a program for ./cmd/main.go
	@echo                                  [RELEASE=1] Release build
	@echo   clean                          Cleaning bin/ directory
	@echo   test [INTEGRATION=1] [LOCAL=1] Run unit test and generate coverage file
	@echo                                  [INTEGRATION=1] Simultaneous run the integration test
	@echo                                  [LOCAL=1] Convert the coverage file to html
	@echo   lint                           Linting all code
	@echo   swag [INSTALL=1]               Generate swagger api document
	@echo                                  [INSTALL=1] Exec `swag_install` task before generate
	@echo   swag_clean                     Run `git checkout` to cleaning doc/api/ directory
	@echo   swag_install                   Install swag command at version 1.8.4

build: $(BIN_PATH)

$(BIN_PATH):
	$(GOBUILD) -o $@ $(BUILD_OPTIONS) ./cmd/main.go

clean:
ifeq ($(GOOS), windows)
	Remove-Item -Path $(BIN_DIR)/*
else
	rm -rf $(BIN_DIR)/*
endif

docker_build:
	docker build --tag sekareco_srv:latest --build-arg ENV="prod" .

docker_run: docker_build
	docker run --rm \
		-p 8000:8000 \
		-v ${CURDIR}/tmp:/log \
		-v ${CURDIR}/tmp:/db \
		-v ${CURDIR}/docs/db:/docs/db \
		-e TZ=Asia/Tokyo \
		--name sekareco_srv \
		sekareco_srv:latest

docker_clean:
	docker image rm sekareco_srv:latest

test: test_setup $(TEST_LOCATE) test_clean

test_setup:
	$(GORUN) ./test/setup

test_clean:
	$(GORUN) ./test/clean

local: $(TEST_MODE)
	$(GOTOOL) cover -html $^.txt -o $^.html

$(TEST_MODE):
	$(GOTEST) -v -p 12 -cover -tags=$@ ./... -coverprofile=$@.txt

lint:
	$(GOLINT) ./...

swag:
ifdef INSTALL
	@$(MAKE) swag_install
endif
	swag init \
		-o ./docs/api/ \
		-d ./cmd/,./interface/handler/ \
		--pd ./domain/ ./usecase/inputdata/ \
		--generatedTime

swag_clean:
	git checkout ./docs/api/

swag_install:
# Install version at 1.8.4
#   because it doesn't work properly with >= 1.8.5
	$(GOINSTALL) github.com/swaggo/swag/cmd/swag@v1.8.4
