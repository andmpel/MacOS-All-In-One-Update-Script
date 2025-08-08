# Makefile for building `update` binaries for macOS

# Commands
MKDIR := mkdir -p
RMDIR := rm -rf

# Go environment
## Debug
GCFLAGS := all=-N -l

## Release
GO_FLAGS := -trimpath -mod=readonly
LDFLAGS := -s -w

# Architectures
GOOS := $(shell go env GOOS)
GO_ARCH := $(shell go env GOARCH)
GO_ARCH_AMD64 := amd64
GO_ARCH_ARM64 := arm64

# Binary names and paths
BIN_DIR := bin
BIN_NAME := macup

DEBUG_BIN := $(BIN_DIR)/$(BIN_NAME)_debug
BIN_AMD64 := $(BIN_DIR)/$(BIN_NAME)_amd64
BIN_ARM64 := $(BIN_DIR)/$(BIN_NAME)_arm64
BIN_UNIVERSAL := $(BIN_DIR)/$(BIN_NAME)

# Build commands
DEBUG_CMD = go build -gcflags="$(GCFLAGS)" -x -v >/dev/null
BUILD_CMD = GOFLAGS="$(GO_FLAGS)" go build -ldflags="$(LDFLAGS)"

.PHONY: all clean test debug build run

all: clean build

debug:
	@echo "Creating Debug Binary"
	@$(MKDIR) $(BIN_DIR)
	@$(DEBUG_CMD) -o $(DEBUG_BIN)
	@echo "Debug Binary Created: $(DEBUG_BIN)"

run: debug
	@echo "Running Debug Binary"
	@$(DEBUG_BIN)

build:
	@echo "Building Binary"
	@$(MKDIR) $(BIN_DIR)
ifeq ($(shell go env GOHOSTOS), darwin)
	@echo "Detected macOS — attempting universal build"
	@if command -v lipo >/dev/null 2>&1; then \
			GOARCH=$(GO_ARCH_AMD64) $(BUILD_CMD) -o $(BIN_AMD64); \
			GOARCH=$(GO_ARCH_ARM64) $(BUILD_CMD) -o $(BIN_ARM64); \
			lipo -create -output $(UNIVERSAL_BINARY) $(BIN_AMD64) $(BIN_ARM64); \
			echo "Universal binary created at $(UNIVERSAL_BINARY)"; \
	else \
			echo "'lipo' not found. Please install Xcode Command Line Tools: xcode-select --install"; \
	fi
else
	@echo "Skipping universal binary creation (not macOS)"
	@echo "Building native binary for $(GOOS) ($(GO_ARCH))"
	@$(BUILD_CMD) -o $(BIN_UNIVERSAL)
endif

test:
	@echo "Running Tests"
	@go test ./... -v

clean:
	@echo "Cleaning Up"
	@$(RMDIR) $(BIN_DIR)
