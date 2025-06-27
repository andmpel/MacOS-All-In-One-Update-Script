.PHONY: all

BIN_NAME=macup
BIN_DIR=./bin
BIN_PATH=$(BIN_DIR)/$(BIN_NAME)

all:

build:
	go build -o $(BIN_PATH)

run: build
	go run .
