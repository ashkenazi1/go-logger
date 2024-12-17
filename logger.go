package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
)

var (
	instance *slog.Logger
	once     sync.Once
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
)

// Config holds logger configuration
type Config struct {
	Environment string
	LogLevel    slog.Level
	UseColors   bool
	Writer      io.Writer // New field for custom writer
}

// ColorHandler implements custom color formatting
type ColorHandler struct {
	w         io.Writer
	useColors bool
	level     slog.Level
}

// Required method for slog.Handler interface
func (h *ColorHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

// Required method for slog.Handler interface
func (h *ColorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

// Required method for slog.Handler interface
func (h *ColorHandler) WithGroup(name string) slog.Handler {
	return h
}

func (h *ColorHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()

	// Add colors based on log level
	if h.useColors {
		switch r.Level {
		case slog.LevelDebug:
			level = colorCyan + level + colorReset
		case slog.LevelInfo:
			level = colorGreen + level + colorReset
		case slog.LevelWarn:
			level = colorYellow + level + colorReset
		case slog.LevelError:
			level = colorRed + level + colorReset
		}
	}

	// Format time as desired
	timeStr := r.Time.Format("2006-01-02 15:04:05")

	// Build the basic message
	fmt.Fprintf(h.w, "[%s] %s: %s", timeStr, level, r.Message)

	// Add attributes
	if r.NumAttrs() > 0 {
		fmt.Fprint(h.w, " {")
		r.Attrs(func(a slog.Attr) bool {
			fmt.Fprintf(h.w, " %s=%v", a.Key, a.Value)
			return true
		})
		fmt.Fprint(h.w, " }")
	}
	fmt.Fprintln(h.w)

	return nil
}

// NewLogger creates a new slog.Logger with environment-specific settings
func newLogger(cfg Config) *slog.Logger {
	// If no writer is provided, default to stdout
	if cfg.Writer == nil {
		cfg.Writer = os.Stdout
	}

	var handler slog.Handler

	switch cfg.Environment {
	case "development":
		handler = &ColorHandler{
			w:         cfg.Writer,
			useColors: cfg.UseColors,
			level:     cfg.LogLevel,
		}
	case "production":
		// Production uses JSON format
		handler = slog.NewJSONHandler(cfg.Writer, &slog.HandlerOptions{
			Level: cfg.LogLevel,
		})
	default:
		handler = &ColorHandler{
			w:         cfg.Writer,
			useColors: cfg.UseColors,
			level:     cfg.LogLevel,
		}
	}

	return slog.New(handler)
}

func InitLogger(cfg Config) *slog.Logger {
	once.Do(func() {
		instance = newLogger(cfg)
	})
	return instance
}

func GetLogger() *slog.Logger {
	if instance == nil {
		panic("Logger not initialized. Call InitLogger first")
	}
	return instance
}
