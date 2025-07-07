# Makefile for mcp-server-chatgpt-app

BINARY_NAME=mcp-server-chatgpt-app-darwin-arm64
CMD_PATH=./cmd/mcp-server-chatgpt-app
MAIN_FILE=$(CMD_PATH)/main.go

# Default target
.PHONY: all
all: build

# Build the optimized binary for distribution
.PHONY: build
build:
	@echo "Building optimized $(BINARY_NAME)..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BINARY_NAME) $(CMD_PATH)
	@echo "Build complete: $(BINARY_NAME)"

# Build development binary (faster build, larger size)
.PHONY: build-dev
build-dev:
	@echo "Building development $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(CMD_PATH)
	@echo "Development build complete: $(BINARY_NAME)"

# Run the application
.PHONY: run
run:
	@echo "Running application..."
	go run $(CMD_PATH)

# Test the application
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)

# Format Go code
.PHONY: fmt
fmt:
	@echo "Formatting Go code..."
	go fmt ./...

# Run Go modules cleanup
.PHONY: tidy
tidy:
	@echo "Running go mod tidy..."
	go mod tidy

# Create DXT package
.PHONY: package
package: build
	@echo "Creating DXT package..."
	npx @anthropic-ai/dxt pack

# Show help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  build      - Build optimized binary for macOS ARM64"
	@echo "  build-dev  - Development build (faster, larger)"
	@echo "  run        - Run directly from source"
	@echo "  test       - Run tests"
	@echo "  fmt        - Format Go code"
	@echo "  tidy       - Clean up Go modules"
	@echo "  clean      - Remove build artifacts"
	@echo "  package    - Create DXT package"
	@echo "  help       - Show this help message" 