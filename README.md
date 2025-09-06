# World Anvil CLI

A Go-based CLI tool and server for interacting with World Anvil, featuring both command-line interface and HTTP server capabilities.

## Project Structure

```
world-anvil-cli/
├── cmd/
│   ├── cli/main.go          # CLI application entry point
│   └── server/main.go       # Server application entry point
├── pkg/
│   ├── config/config.go     # Shared configuration management
│   ├── models/models.go     # Data models (User, World, Article)
│   └── utils/http.go        # HTTP utilities and helpers
├── internal/
│   ├── client/client.go     # HTTP client for API calls
│   └── server/server.go     # HTTP server handlers and routes
├── go.mod                   # Go module definition
├── Makefile                 # Build automation and common tasks
└── .gitignore               # Git ignore rules
```

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd world-anvil-cli
   ```

2. Install dependencies:
   ```bash
   make deps
   ```

## Usage

### Building

```bash
# Build both CLI and server binaries
make build

# Build only CLI
make cli

# Build only server
make server
```

### Running in Development

```bash
# Run CLI in development mode
make run-cli

# Run server in development mode
make run-server
```

### CLI Commands

```bash
# Show version
./bin/world-anvil-cli version

# Show configuration
./bin/world-anvil-cli config
```

### Server Endpoints

The server runs on port 8080 by default and provides the following endpoints:

- `GET /health` - Health check endpoint
- `GET /api/users` - Get users (placeholder)
- `GET /api/worlds` - Get worlds (placeholder)

### Configuration

The application uses environment variables for configuration:

- `PORT` - Server port (default: 8080)
- `DEBUG` - Enable debug mode (default: false)
- `API_KEY` - API key for authentication

Create a `.env` file in the project root:
```bash
PORT=8080
DEBUG=true
API_KEY=your_api_key_here
```

## Development

### Available Make Commands

```bash
make help          # Show all available commands
make build         # Build both CLI and server
make cli           # Build CLI binary
make server        # Build server binary
make test          # Run tests
make lint          # Run linter (requires golangci-lint)
make fmt           # Format code
make clean         # Clean build artifacts
make run-cli       # Run CLI in development
make run-server    # Run server in development
make deps          # Install and tidy dependencies
```

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
go test -cover ./...
```

### Code Quality

```bash
# Format code
make fmt

# Run linter (install golangci-lint first)
make lint
```

## License

This project is licensed under the terms specified in the LICENSE file.