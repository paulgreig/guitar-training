# Guitar Training Application

A text-based user interface (TUI) guitar training application built with Go to help users learn and practice guitar scales and lessons through an interactive terminal interface.

## Features

- **Scale Browser**: Browse and view guitar scales with text-based fretboard visualization
- **Lesson System**: Access structured lessons organized by skill level (beginner, intermediate, advanced)
- **Interactive Navigation**: Navigate through scales and lessons using keyboard controls
- **Clean TUI**: Beautiful terminal interface built with Bubble Tea

## Screenshots

*Screenshots will be added as the UI is developed*

## Installation

### Prerequisites

- Go 1.21 or higher

### Build from Source

1. Clone the repository:
```bash
git clone https://github.com/paulgreig/guitar-training.git
cd guitar-training
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o guitar-training cmd/server/main.go
```

4. Run the application:
```bash
./guitar-training
```

Or run directly:
```bash
go run cmd/server/main.go
```

## Usage

### Starting the Application

Simply run the binary or use `go run`:

```bash
go run cmd/server/main.go
```

### Navigation

- **Arrow Keys** or **j/k**: Navigate up and down
- **Enter**: Select an item or view details
- **Esc**: Go back to the previous screen
- **q** or **Ctrl+C**: Quit the application

### Main Menu

1. **View Scales**: Browse available guitar scales
2. **View Lessons**: Browse available lessons
3. **Quit**: Exit the application

### Viewing Scales

- Select a scale from the list to view its details
- See the scale notes and fretboard positions
- Text-based fretboard shows where to play the scale

### Viewing Lessons

- Browse lessons organized by level
- Select a lesson to read its full content
- Lessons include practice tips and techniques

## Project Structure

```
guitar-training/
├── cmd/server/          # Application entry point
├── internal/
│   ├── tui/             # TUI components (Bubble Tea)
│   ├── models/          # Data models
│   └── config/          # Configuration
├── data/                # JSON data files
│   ├── scales.json
│   └── lessons.json
└── README.md
```

## Data Files

The application reads from JSON files in the `data/` directory:

- `data/scales.json`: Scale definitions with notes and positions
- `data/lessons.json`: Lesson content organized by level

You can edit these files to add your own scales and lessons.

## Development

### Running Tests

```bash
go test ./...
```

### Building for Different Platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o bin/guitar-training-linux cmd/server/main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o bin/guitar-training-windows.exe cmd/server/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o bin/guitar-training-macos cmd/server/main.go
```

## Observability and performance monitoring

The app includes logging and Prometheus metrics for observability and performance:

- **Logging**: Configurable level via `LOG_LEVEL` (debug, info, warn, error). Logs are written to `logs/app.log`.
- **Prometheus metrics**: An HTTP server (default port **9090**) serves `/metrics` with:
  - **Menu selection performance**: Histogram `guitar_training_menu_selection_duration_seconds` for latency of scales (and other) menu actions.
  - **View counts**: Counters for scales/lessons list and detail views.
  - **Go runtime**: CPU, memory, and GC metrics from the Prometheus Go and process collectors.

Set `METRICS_PORT=0` to disable the metrics server. See [docs/METRICS.md](docs/METRICS.md) for Prometheus scrape config and Grafana Cloud setup.

## Technology Stack

- **Go 1.21+**: Programming language
- **Bubble Tea**: TUI framework (github.com/charmbracelet/bubbletea)
- **Lipgloss**: Styling library (github.com/charmbracelet/lipgloss)
- **Prometheus**: Metrics (client_golang) for performance and runtime monitoring
- **JSON**: Data storage format

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Roadmap

See [PLAN.md](PLAN.md) for detailed development phases and roadmap.

## License

TBD
