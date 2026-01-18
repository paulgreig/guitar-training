# Guitar Training Application

A comprehensive guitar training application built with Go to help users learn and practice guitar through interactive exercises, lessons, and progress tracking.

## Features

- **Chord Practice**: Display chord diagrams, practice chord transitions
- **Scale Practice**: Learn and practice scales with visual fretboard display
- **Rhythm Training**: Metronome with various time signatures
- **Practice Tracking**: Log practice sessions and track progress
- **Interactive Lessons**: Structured lessons for different skill levels
- **Progress Analytics**: View practice metrics and improvement over time

## Project Structure

```
guitar-training/
├── cmd/server/          # Application entry point
├── internal/            # Private application code
│   ├── api/            # HTTP handlers and routes
│   ├── models/         # Data models
│   ├── repository/     # Data access layer
│   ├── service/        # Business logic
│   ├── utils/          # Utility functions
│   └── config/         # Configuration management
├── pkg/                # Public packages
├── migrations/         # Database migrations
├── data/               # Data files (chords, scales, etc.)
└── configs/            # Configuration files
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- SQLite (for development) or PostgreSQL (for production)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/paulgreig/guitar-training.git
cd guitar-training
```

2. Install dependencies:
```bash
go mod download
```

3. Set up configuration:
```bash
cp configs/config.example.yaml configs/config.yaml
# Edit config.yaml with your settings
```

4. Run migrations (when available):
```bash
# TBD: Migration commands
```

5. Run the server:
```bash
go run cmd/server/main.go
```

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run integration tests
go test ./tests/integration/...
```

### Building

```bash
# Build binary
go build -o bin/guitar-training cmd/server/main.go

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o bin/guitar-training-linux cmd/server/main.go
GOOS=windows GOARCH=amd64 go build -o bin/guitar-training-windows.exe cmd/server/main.go
GOOS=darwin GOARCH=amd64 go build -o bin/guitar-training-macos cmd/server/main.go
```

## API Endpoints

API documentation will be available once the server is running. Check the `/docs` endpoint for Swagger/OpenAPI documentation.

## Technology Stack

- **Backend**: Go 1.21+
- **Web Framework**: Gin/Echo/Chi
- **Database**: SQLite (dev) / PostgreSQL (prod)
- **ORM**: GORM or sqlx
- **Config**: Viper
- **Logging**: Zap or Logrus

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

TBD

## Roadmap

See [PLAN.md](PLAN.md) for detailed development phases and roadmap.
