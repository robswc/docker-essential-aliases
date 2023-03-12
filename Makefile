# use bash
SHELL := /bin/bash

PATH := $(PATH):/usr/local/go/bin

BUILD_DIR=bin
BINARY_NAME=dea

build:
	@echo "Building..."
	@mkdir -p bin
	@cd src && go build -o ../${BUILD_DIR}/${BINARY_NAME} main.go

clean:
	@echo "Cleaning up..."
	@rm -rf bin