# Variables
APP_NAME := testku
# testing CMD dir
CMD := testku
CMD_DIR := ./cmd/$(CMD)
MAIN_FILE := $(CMD_DIR)/main.go
BUILD_DIR := build
BINARY := $(BUILD_DIR)/$(APP_NAME)

# Default target
.PHONY: all
all: build

# Build the Go application
.PHONY: build
build:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BINARY) $(MAIN_FILE)
	@echo "Build completed: $(BINARY)"

# Run the Go application
.PHONY: run
run: build
	@echo "Running the application..."
	@$(BINARY)

# Clean the build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean completed."

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting the code..."
	@go fmt ./...

# Lint the code
.PHONY: lint
lint:
	@echo "Linting the code..."
	@golangci-lint run

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Display help
.PHONY: help
help:
	@echo "Makefile for Go project"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all      - Default target, builds the application"
	@echo "  build    - Build the Go application"
	@echo "  run      - Run the Go application"
	@echo "  clean    - Clean the build artifacts"
	@echo "  test     - Run tests"
	@echo "  fmt      - Format the code"
	@echo "  lint     - Lint the code"
	@echo "  deps     - Install dependencies"
	@echo "  help     - Display this help message"
