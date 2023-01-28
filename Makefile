.DEFAULT_GOAL:=help

# Enforce powershell to unify execution commands
GOOS:=$(shell go env GOOS)
ifeq ($(GOOS), windows)
	SHELL:=powershell.exe
endif

#####
# go cmd list
GOCMD:=go
GORUN:=$(GOCMD) run
GOBUILD:=$(GOCMD) build
GOTEST:=$(GOCMD) test
GOLINT:=$(GOCMD) vet
GOTOOL:=$(GOCMD) tool
GOINSTALL:=$(GOCMD) install

#####
# build
BIN_DIR:=./bin
BIN_NAME:=serverd
BIN_PATH:=$(BIN_DIR)/$(BIN_NAME)

BUILD_TAGS:=develop
BUILD_RACE:=-race
BUILD_STATIC:=

ifdef RELEASE
	BUILD_TAGS:=production
	BUILD_RACE:=
	BUILD_STATIC:=-a
endif
BUILD_OPTIONS:=-ldflags '-s -w' -tags $(BUILD_TAGS) $(BUILD_RACE) $(BUILD_STATIC)

#####
# test
MODE_UNIT:=unit
MODE_INTEGRATION:=integration

COVERAGE_OUTPUT:=./coverage/
COVERAGE_EXTENTION:=txt

TEST_MODE:=$(MODE_UNIT)
TEST_LOCATE:=local

ifdef INTEGRATION
	TEST_MODE:=$(MODE_INTEGRATION)
endif

# `local` task is convert the output coverage file to html
# Skipping `local` task by makeing TEST_LOCATE = TEST_MODE
ifdef CI
	COVERAGE_OUTPUT:=./
	TEST_LOCATE:=$(TEST_MODE)
endif


.PHONY: help build clean test

help:
ifeq ($(GOOS), windows)
	@Get-Content ./makehelp.txt | Out-Host
else
	@cat ./makehelp.txt
endif

build: clean $(BIN_PATH)

$(BIN_PATH):
	$(GOBUILD) -o $@ $(BUILD_OPTIONS) ./cmd/main.go

clean:
ifeq ($(GOOS), windows)
	Remove-Item -Force -Path $(BIN_DIR)/*
else
	rm -rf $(BIN_DIR)/*
endif

docker_build:
	docker build --tag sekareco_srv:latest --build-arg tags=${RELEASE} .

docker_run: docker_build
	docker run --rm \
		-p 8000:8000 \
		-v $(CURDIR)/log:/log \
		-v $(CURDIR)/db:/db \
		-v $(CURDIR)/docs/db:/docs/db \
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
	$(GOTOOL) cover -html $(COVERAGE_OUTPUT)$^.$(COVERAGE_EXTENTION) -o $(COVERAGE_OUTPUT)$^.html

$(TEST_MODE):
	$(GOTEST) -v -p 12 -cover -tags=$@,test ./... -coverprofile=$(COVERAGE_OUTPUT)$@.$(COVERAGE_EXTENTION)

lint:
	$(GOLINT) ./...

setup:
	$(GORUN) ./tools/setup/

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
