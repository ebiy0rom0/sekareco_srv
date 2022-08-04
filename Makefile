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

BUILD_OPTIONS:=-ldflags '-s -w' -tags $(BUILD_TAGS) $(BUILD_RACE) $(BUILD_STATIC)

SWAG_INIT:=
ifdef CI
	SWAG_INIT:=swag_init
endif


.PHONY: build clean test lint swag swag_init

build: $(BIN_PATH)

$(BIN_PATH):
	$(GOBUILD) -o $@ $(BUILD_OPTIONS) ./cmd/main.go

clean:
	rm -rf ./bin/*

test:
	go test -v ./...

lint:
	go vet ./...

swag: $(SWAG_INIT)
	swag init -o ./doc/api/ -d ./cmd/,./interface/handler/ --pd .\domain\ --generatedTime

swag_init:
	go install github.com/swaggo/swag/cmd/swag@latest