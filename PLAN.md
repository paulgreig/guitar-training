# Guitar Training Application - Development Plan

## Project Overview
A comprehensive guitar training application to help users learn and practice guitar through interactive exercises, lessons, and progress tracking.

## Core Features

### 1. Practice Exercises
- **Chord Practice**: Display chord diagrams, practice chord transitions
- **Scale Practice**: Learn and practice scales with visual fretboard display
- **Rhythm Training**: Metronome with various time signatures
- **Finger Exercises**: Warm-up routines and dexterity drills
- **Song Practice**: Learn songs with tablature display

### 2. Learning Modules
- **Beginner Lessons**: Basic chords, strumming patterns, simple songs
- **Intermediate Lessons**: Barre chords, fingerpicking, scales
- **Advanced Lessons**: Complex techniques, music theory, improvisation
- **Interactive Tutorials**: Step-by-step guided lessons

### 3. Visual Tools
- **Interactive Fretboard**: Clickable fretboard for chord/scale visualization
- **Tablature Display**: Show guitar tabs with timing
- **Chord Diagrams**: Visual chord representations
- **Scale Diagrams**: Highlighted scale patterns on fretboard

### 4. Audio Features
- **Metronome**: Adjustable BPM, time signatures, subdivisions
- **Audio Playback**: Play chords, scales, and examples
- **Recording**: Record practice sessions for review
- **Tuning**: Guitar tuner functionality

### 5. Progress Tracking
- **Practice Log**: Track practice time and frequency
- **Statistics Dashboard**: View practice metrics and improvement
- **Achievement System**: Unlock achievements as skills improve
- **Goal Setting**: Set and track practice goals

### 6. Customization
- **Guitar Type**: Support for 6-string, 7-string, acoustic, electric
- **Tuning Options**: Standard, drop D, open tunings, etc.
- **Difficulty Levels**: Adjustable difficulty for exercises
- **Practice Preferences**: Customize practice sessions

## Technical Architecture

### Backend Service (Golang)
- **Language**: Go 1.21+
- **Web Framework**: Gin, Echo, or Chi (REST API)
- **Database**: SQLite (local) or PostgreSQL (production)
- **ORM/Query Builder**: GORM or sqlx
- **Configuration**: Viper for config management
- **Logging**: Logrus or Zap
- **Audio Processing**: Beep or PortAudio (Go bindings)
- **Testing**: Built-in testing package + testify

### Frontend Options

#### Option 1: Web Frontend (React/Vue/Svelte)
- **UI Framework**: React, Vue, or Svelte with TypeScript
- **Styling**: Tailwind CSS or CSS Modules
- **API Client**: HTTP client for REST API
- **Visualization**: Canvas/SVG for fretboard rendering
- **Audio**: Web Audio API or Tone.js (client-side)

#### Option 2: Desktop GUI (Wails/Fyne)
- **Wails**: Go backend + HTML/JS frontend (similar to Electron)
- **Fyne**: Native Go GUI toolkit
- **Go-qt**: Go bindings for Qt

#### Option 3: CLI + Web Interface
- **CLI Tool**: Cobra for command-line interface
- **Web UI**: Go HTML templates + HTMX or Alpine.js
- **Static Assets**: Embedded files with embed package

## Development Phases

### Phase 1: Foundation (MVP) - Go Backend Setup
- [ ] Initialize Go module and project structure
- [ ] Set up database schema and migrations
- [ ] Implement data models (Chord, Scale, PracticeSession)
- [ ] Create repository layer for data access
- [ ] Set up REST API with Gin/Echo
- [ ] Create basic chord and scale endpoints
- [ ] Set up configuration management (Viper)
- [ ] Add logging infrastructure

### Phase 2: Core API Features
- [ ] Chord API endpoints (GET, list, search)
- [ ] Scale API endpoints
- [ ] Practice session tracking endpoints
- [ ] Practice statistics and analytics endpoints
- [ ] Data seeding (chords, scales from JSON/data files)
- [ ] Input validation and error handling
- [ ] API documentation (Swagger/OpenAPI)

### Phase 3: Frontend Integration
- [ ] Set up frontend project (React/Vue or Wails)
- [ ] Create API client/service layer
- [ ] Build fretboard visualization component
- [ ] Implement chord display system
- [ ] Add practice timer and logging UI
- [ ] Create progress dashboard

### Phase 4: Enhanced Features
- [ ] Audio service (metronome, playback) using Beep
- [ ] Lesson system API endpoints
- [ ] Exercise library API
- [ ] Tablature processing and display
- [ ] Tuning functionality
- [ ] File upload for custom lessons/songs

### Phase 5: Advanced Features & Polish
- [ ] Song learning system
- [ ] Recording capabilities (if applicable)
- [ ] Achievement/badge system
- [ ] Authentication/authorization (optional)
- [ ] Performance optimization
- [ ] Comprehensive testing (unit + integration)
- [ ] Docker containerization
- [ ] Documentation and deployment guides

## Technology Stack

### Backend Service (Golang)
```
Language: Go 1.21+
Web Framework: Gin/Echo/Chi
Database: SQLite (dev) / PostgreSQL (prod)
ORM: GORM or sqlx
Config: Viper
Logging: Zap or Logrus
Testing: testing package + testify
Audio: Beep or PortAudio bindings
```

### Frontend Options

#### Option A: Web App (Go API + React Frontend)
```
Backend: Go REST API (Gin/Echo)
Frontend: React + TypeScript + Tailwind CSS
Database: SQLite/PostgreSQL
Deployment: 
  - Backend: Docker, Railway, Fly.io, or self-hosted
  - Frontend: Vercel, Netlify, or same server
```

#### Option B: Desktop App (Wails)
```
Backend: Go (embedded in Wails)
Frontend: HTML/JS/CSS (React/Vue/Svelte)
Framework: Wails v2
Packaging: Wails build system (cross-platform)
```

#### Option C: Desktop App (Fyne - Pure Go)
```
GUI: Fyne (native Go GUI)
Backend: Go (same process)
Packaging: Fyne packaging tools
Audio: Beep library
```

#### Option D: CLI Tool + Web Server
```
CLI: Cobra
Web Server: Go with HTML templates
API: REST API (Gin/Echo)
Frontend: HTMX or Alpine.js for interactivity
```

### Recommended: Option A (Go API + React Frontend)
- Separation of concerns
- Scalable architecture
- Modern development workflow
- Easy deployment options

## File Structure (Go Service + Web Frontend)

```
guitar-training/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/              # HTTP handlers
│   │   │   ├── chord_handler.go
│   │   │   ├── scale_handler.go
│   │   │   ├── practice_handler.go
│   │   │   └── lesson_handler.go
│   │   ├── middleware/            # HTTP middleware
│   │   └── router.go              # Route definitions
│   ├── models/                    # Data models
│   │   ├── chord.go
│   │   ├── scale.go
│   │   ├── practice_session.go
│   │   └── lesson.go
│   ├── repository/                # Data access layer
│   │   ├── chord_repository.go
│   │   ├── scale_repository.go
│   │   └── practice_repository.go
│   ├── service/                   # Business logic
│   │   ├── chord_service.go
│   │   ├── scale_service.go
│   │   ├── practice_service.go
│   │   └── audio_service.go
│   ├── utils/
│   │   ├── guitar.go              # Guitar-related utilities
│   │   ├── audio.go               # Audio processing
│   │   └── validation.go
│   └── config/
│       └── config.go              # Configuration management
├── pkg/                           # Public packages (if needed)
│   └── guitar/
│       ├── chord.go
│       └── scale.go
├── migrations/                    # Database migrations
│   └── *.sql
├── web/                           # Frontend (if embedding)
│   ├── static/                    # Static assets
│   └── templates/                 # HTML templates (if using server-side)
├── frontend/                      # Separate frontend (if using React/Vue)
│   ├── src/
│   ├── public/
│   └── package.json
├── data/                          # Data files (chords, scales, etc.)
│   ├── chords.json
│   ├── scales.json
│   └── lessons.json
├── configs/                       # Configuration files
│   └── config.yaml
├── scripts/                       # Build/deployment scripts
├── tests/
│   ├── integration/
│   └── unit/
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
└── README.md
```

### Go Package Structure Guidelines
- `cmd/`: Application entry points
- `internal/`: Private application code
- `pkg/`: Public library code (if creating reusable packages)
- `api/`: API handlers and routes
- `models/`: Data structures
- `repository/`: Database access (data layer)
- `service/`: Business logic (service layer)

## Data Models (Go Structs)

### Chord
```go
type Chord struct {
    ID        int      `json:"id" gorm:"primaryKey"`
    Name      string   `json:"name"`
    Notes     []string `json:"notes" gorm:"type:jsonb"`
    Frets     []int    `json:"frets" gorm:"type:jsonb"`
    Fingers   []int    `json:"fingers" gorm:"type:jsonb"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Scale
```go
type Scale struct {
    ID        int      `json:"id" gorm:"primaryKey"`
    Name      string   `json:"name"`
    Notes     []string `json:"notes" gorm:"type:jsonb"`
    Positions []Position `json:"positions" gorm:"type:jsonb"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Position struct {
    Fret  int   `json:"fret"`
    Strings []int `json:"strings"`
}
```

### PracticeSession
```go
type PracticeSession struct {
    ID           int       `json:"id" gorm:"primaryKey"`
    UserID       string    `json:"user_id"` // Optional for multi-user
    Date         time.Time `json:"date"`
    Duration     int       `json:"duration"` // in minutes
    ExerciseType string    `json:"exercise_type"`
    Score        float64   `json:"score"`
    Rating       int       `json:"rating"` // 1-5
    Notes        string    `json:"notes"`
    CreatedAt    time.Time `json:"created_at"`
}
```

### Lesson
```go
type Lesson struct {
    ID        string    `json:"id" gorm:"primaryKey"`
    Title     string    `json:"title"`
    Level     string    `json:"level"` // beginner, intermediate, advanced
    Content   string    `json:"content"` // JSON or text
    Exercises []Exercise `json:"exercises" gorm:"type:jsonb"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Exercise struct {
    Type    string `json:"type"`
    Content string `json:"content"`
}
```

## Next Steps

1. **Initialize Go Module**: Run `go mod init github.com/paulgreig/guitar-training`
2. **Set Up Project Structure**: Create directory structure following Go conventions
3. **Choose Web Framework**: Install Gin, Echo, or Chi for REST API
4. **Set Up Database**: Configure SQLite or PostgreSQL with GORM
5. **Create Data Models**: Define Chord, Scale, PracticeSession structs
6. **Implement Repository Layer**: Create data access layer
7. **Build API Endpoints**: Start with basic GET endpoints for chords/scales
8. **Add Data Seeding**: Load chords and scales from JSON files
9. **Set Up Frontend**: Choose React/Vue or Wails for UI
10. **Test & Iterate**: Build MVP with core features

## Go-Specific Considerations

### Package Management
- Use Go modules (`go mod`)
- Version dependencies with `go get`
- Use `go.work` for monorepo if frontend is separate

### Testing
- Unit tests: `*_test.go` files with `testing` package
- Integration tests: Separate `test/integration` directory
- Use `testify` for assertions and mocks
- Benchmark tests for performance-critical code

### Database
- **Development**: SQLite (zero configuration)
- **Production**: PostgreSQL (better for concurrent access)
- Migrations: Use `golang-migrate` or `goose`
- ORM: GORM (convenient) or sqlx (raw SQL control)

### Configuration
- Environment variables for secrets
- YAML/JSON config files with Viper
- Different configs for dev/staging/prod

### Audio in Go
- **Beep**: Pure Go audio library (recommended for simplicity)
- **PortAudio**: Go bindings for cross-platform audio I/O
- Consider delegating audio to frontend (Web Audio API) for web apps

### Deployment
- Single binary deployment (no runtime dependencies)
- Docker multi-stage builds for optimized images
- Cross-compilation for different platforms
- Consider `upx` for binary compression

## Questions to Consider

- What's the primary platform? (Web API, Desktop with Wails, CLI tool)
- Do you need cloud sync for progress? (Database backend)
- What's the target skill level? (Beginners, All levels)
- Any specific features you want to prioritize?
- Single-user or multi-user application? (affects auth requirements)

---

*This is a living document - update as the project evolves*
