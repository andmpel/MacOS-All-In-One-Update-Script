# Makefile for building `update` binaries for macOS

GOOS := darwin
GOARCH_AMD64 := amd64
GOARCH_ARM64 := arm64
BIN_DIR := bin
BIN_NAME := macup

BIN_AMD64 := $(BIN_DIR)/$(BIN_NAME)_amd64
BIN_ARM64 := $(BIN_DIR)/$(BIN_NAME)_arm64
BIN_UNIVERSAL := $(BIN_DIR)/$(BIN_NAME)

.PHONY: all clean test build universal

all: build universal

build:
	@echo "🔧 Creating Binaries"
	@GOOS=$(GOOS) GOARCH=$(GOARCH_AMD64) go build -o $(BIN_AMD64)
	@GOOS=$(GOOS) GOARCH=$(GOARCH_ARM64) go build -o $(BIN_ARM64)
	@echo "✅ Binary Created: $(BIN_AMD64), $(BIN_ARM64)"

universal:
	@command -v lipo >/dev/null 2>&1 || { echo >&2 "⚠️ lipo not found. Run 'xcode-select --install' on macOS."; exit 1; }
	@echo "🔧 Creating Universal Binary"
	@lipo -create -output $(BIN_UNIVERSAL) $(BIN_AMD64) $(BIN_ARM64)
	@echo "✅ Universal Binary Created: $(BIN_UNIVERSAL)"

test:
	@echo "🧪 Running Tests"
	@go test ./... -v

clean:
	@echo "🧹 Cleaning Up"
	@rm -rf $(BIN_DIR)
