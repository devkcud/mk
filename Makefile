BIN_NAME=mk
BIN_VERSION=$(shell cat version.txt)

COMPILE_NAME=$(BIN_NAME)-$(BIN_VERSION)
BUILD_DIR=build

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

all: clean build

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(COMPILE_NAME) -v ./cmd/mk

clean:
	$(GOCLEAN)
	-rm -r $(BUILD_DIR)

install: build
	cp $(BUILD_DIR)/$(COMPILE_NAME) ${HOME}/.local/bin/$(BIN_NAME)

.PHONY: build clean install
