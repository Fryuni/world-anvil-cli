.PHONY: build clean test lint fmt help cli server

# Build both CLI and server
build: cli server

# Build CLI binary
cli:
	go build -o bin/world-anvil-cli ./cmd/cli

# Build server binary
server:
	go build -o bin/world-anvil-server ./cmd/server

# Run tests
test:
	go test ./...

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Run CLI in development
run-cli:
	go run ./cmd/cli

# Run server in development  
run-server:
	go run ./cmd/server

# Install dependencies
deps:
	go mod download
	go mod tidy

# Show help
help:
	@echo "Available commands:"
	@echo "  build     - Build both CLI and server"
	@echo "  cli       - Build CLI binary"
	@echo "  server    - Build server binary"
	@echo "  test      - Run tests"
	@echo "  lint      - Run linter"
	@echo "  fmt       - Format code"
	@echo "  clean     - Clean build artifacts"
	@echo "  run-cli   - Run CLI in development"
	@echo "  run-server - Run server in development"
	@echo "  deps      - Install dependencies"