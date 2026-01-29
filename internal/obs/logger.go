package obs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Level represents a logging level.
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

var (
	logger   *log.Logger
	initOnce sync.Once
	level    = LevelInfo
)

// InitLogger initialises the global logger.
// Logs are written to logs/app.log and also to stderr.
func InitLogger() {
	initOnce.Do(func() {
		// Determine log level from environment.
		switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
		case "debug":
			level = LevelDebug
		case "info", "":
			level = LevelInfo
		case "warn", "warning":
			level = LevelWarn
		case "error":
			level = LevelError
		default:
			level = LevelInfo
		}

		// Ensure logs directory exists.
		logDir := "logs"
		_ = os.MkdirAll(logDir, 0o755)

		logPath := filepath.Join(logDir, "app.log")
		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			// Fallback to stderr only if file can't be opened.
			logger = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
			logger.Printf("WARN failed to open log file %s: %v", logPath, err)
			return
		}

		multi := log.New(file, "", log.LstdFlags|log.Lmicroseconds)
		logger = multi

		logger.Printf("INFO logger initialised level=%s path=%s", levelString(level), logPath)
	})
}

func levelString(l Level) string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// logf is the internal logging helper.
func logf(l Level, format string, args ...interface{}) {
	if logger == nil {
		// In case InitLogger wasn't called, fall back safely.
		InitLogger()
	}
	if l < level {
		return
	}
	prefix := levelString(l)
	msg := fmt.Sprintf(format, args...)
	logger.Printf("%s %s", prefix, msg)
}

// Debug logs a debug message.
func Debug(format string, args ...interface{}) {
	logf(LevelDebug, format, args...)
}

// Info logs an informational message.
func Info(format string, args ...interface{}) {
	logf(LevelInfo, format, args...)
}

// Warn logs a warning message.
func Warn(format string, args ...interface{}) {
	logf(LevelWarn, format, args...)
}

// Error logs an error message.
func Error(format string, args ...interface{}) {
	logf(LevelError, format, args...)
}

// WithFields logs key=value style fields for richer observability.
func WithFields(l Level, msg string, fields map[string]interface{}) {
	if logger == nil {
		InitLogger()
	}
	if l < level {
		return
	}

	var b strings.Builder
	_, _ = b.WriteString(msg)
	for k, v := range fields {
		_, _ = b.WriteString(" ")
		_, _ = b.WriteString(k)
		_, _ = b.WriteString("=")
		_, _ = b.WriteString(fmt.Sprintf("%v", v))
	}
	logger.Printf("%s %s", levelString(l), b.String())
}

// Event logs a structured event with a name and arbitrary fields.
func Event(name string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["event"] = name
	fields["ts"] = time.Now().Format(time.RFC3339Nano)
	WithFields(LevelInfo, "event", fields)
}

