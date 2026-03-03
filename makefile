# Variables
BINARY_NAME=passgen
VERSION=1.0.0
LDFLAGS=-ldflags="-s -w"

.PHONY: all clean windows linux mac

# Build for all platforms
all: windows linux mac

# Windows Build
windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY_NAME).exe main.go

# Linux Build
linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-linux main.go

# macOS Build (Intel and Apple Silicon)
mac:
	@echo "Building for macOS (Intel)..."
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-mac-amd64 main.go
	@echo "Building for macOS (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-mac-arm64 main.go

# Clean up the bin directory
clean:
	@echo "Cleaning up..."
	rm -rf bin/