GOCMD:=go
GOBUILD:=$(GOCMD) build

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

BUILD_OPTIONS:=-tags $(BUILD_TAGS) $(BUILD_RACE) $(BUILD_STATIC)

.PHONY: build clean

build: $(BIN_PATH)

$(BIN_PATH):
	$(GOBUILD) -o $@ $(BUILD_OPTIONS) ./cmd/main.go

clean:
	@echo $(shell go env GOOS)
#	rm -rf ./bin
