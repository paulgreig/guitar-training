# Guitar Training Application - Development Plan

## Project Overview
A text-based user interface (TUI) guitar training application built with Go to help users learn and practice guitar scales and lessons through an interactive terminal interface.

## Core Features (Reduced Scope)

### 1. Scale Display
- **Scale List**: Browse available guitar scales
- **Scale Details**: View scale notes and fretboard positions
- **Text-based Fretboard**: Visual representation of scales on the fretboard

### 2. Lesson System
- **Lesson List**: Browse available lessons by level
- **Lesson Content**: Read structured lesson content
- **Level-based Organization**: Beginner, intermediate, and advanced lessons

## Technical Architecture

### Application (Golang TUI)
- **Language**: Go 1.21+
- **TUI Framework**: Bubble Tea (charmbracelet/bubbletea)
- **Styling**: Lipgloss (charmbracelet/lipgloss)
- **Data Storage**: JSON files (no database required)
- **Configuration**: Environment variables or simple config file

### Data Format
- **Scales**: JSON file with scale definitions (name, notes, positions)
- **Lessons**: JSON file with lesson content (id, title, level, content)

## Development Phases

### Phase 1: Foundation (MVP) - TUI Setup ✅
- [x] Initialize Go module and project structure
- [x] Set up TUI framework (Bubble Tea)
- [x] Create basic menu navigation
- [x] Implement data loading from JSON files
- [x] Create sample scales and lessons data

### Phase 2: Core TUI Features
- [ ] Implement scale list view
- [ ] Implement scale detail view with fretboard
- [ ] Implement lesson list view
- [ ] Implement lesson detail view
- [ ] Improve navigation and keyboard controls
- [ ] Add search/filter functionality

### Phase 3: Enhanced Display
- [ ] Improve fretboard visualization
- [ ] Add scale position highlighting
- [ ] Better text formatting for lessons
- [ ] Add color themes
- [ ] Improve layout and spacing

### Phase 4: Data Expansion
- [ ] Add more scales (pentatonic, blues, modes, etc.)
- [ ] Expand lesson library
- [ ] Add scale exercises
- [ ] Create structured lesson progression

### Phase 5: Polish & Enhancement
- [ ] Add keyboard shortcuts help
- [ ] Improve error handling
- [ ] Add data validation
- [ ] Performance optimization
- [ ] Comprehensive testing
- [ ] Documentation

## Technology Stack

### TUI Application
```
Language: Go 1.21+
TUI Framework: Bubble Tea (github.com/charmbracelet/bubbletea)
Styling: Lipgloss (github.com/charmbracelet/lipgloss)
Data: JSON files (no database)
Config: Simple environment variables
```

## File Structure

```
guitar-training/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── tui/                        # TUI components
│   │   ├── model.go               # Main TUI model
│   │   └── data.go                # Data loading logic
│   ├── models/                    # Data models
│   │   ├── scale.go
│   │   └── lesson.go
│   └── config/
│       └── config.go              # Configuration management
├── data/                          # Data files
│   ├── scales.json
│   └── lessons.json
├── go.mod
├── go.sum
└── README.md
```

## Data Models

### Scale
```go
type Scale struct {
    Name      string     `json:"name"`
    Notes     []string   `json:"notes"`
    Positions []Position `json:"positions"`
}

type Position struct {
    Fret    int   `json:"fret"`
    Strings []int `json:"strings"`
}
```

### Lesson
```go
type Lesson struct {
    ID      string `json:"id"`
    Title   string `json:"title"`
    Level   string `json:"level"` // beginner, intermediate, advanced
    Content string `json:"content"`
}
```

## User Interface Flow

1. **Main Menu**
   - View Scales
   - View Lessons
   - Quit

2. **Scales View**
   - List of available scales
   - Navigate with arrow keys
   - Enter to view details

3. **Scale Detail View**
   - Scale name and notes
   - Text-based fretboard visualization
   - Position markers

4. **Lessons View**
   - List of lessons by level
   - Navigate with arrow keys
   - Enter to view content

5. **Lesson Detail View**
   - Lesson title and level
   - Full lesson content
   - Formatted text display

## Keyboard Controls

- `↑/↓` or `j/k`: Navigate up/down
- `Enter`: Select/view details
- `Esc`: Go back to previous view
- `q` or `Ctrl+C`: Quit application

## Next Steps

1. ✅ **Initialize Go Module**: Completed
2. ✅ **Set Up TUI Framework**: Completed
3. ✅ **Create Basic Structure**: Completed
4. **Test and Refine**: Run application and improve UI
5. **Expand Data**: Add more scales and lessons
6. **Enhance Features**: Add search, filtering, better visualization

## Go-Specific Considerations

### Package Management
- Use Go modules (`go mod`)
- Version dependencies with `go get`

### Testing
- Unit tests: `*_test.go` files with `testing` package
- Test data loading and parsing
- Test TUI model updates

### Data Management
- JSON files for simplicity
- No database required for MVP
- Easy to edit and extend

### Deployment
- Single binary deployment
- Cross-platform compilation
- No external dependencies

---

*This is a living document - update as the project evolves*
