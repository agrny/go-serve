# Binary name (change to your project name)
BINARY := go-server

# Output directory
DIST := .dist

# Default target
all: build

# Build binary into .dist
build:
	@mkdir -p $(DIST)
	@echo "Building $(BINARY)..."
	GO111MODULE=on go build -o $(DIST)/$(BINARY) .
	@echo "Binary created at $(DIST)/$(BINARY)"

run: build
	  @./$(DIST)/$(BINARY)

# open all go files in editor
edit:
	nvim $(shell find . -type f -name "*.go")

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(DIST)

.PHONY: all build clean

