# Go Logger

A flexible, colorful logging library for Go built on top of `slog`. The library provides a simple way to create environment-aware loggers with colorized output and custom writers.

## Features
- Environment-aware logging (development/production modes)
- Colorized console output for different log levels
- Flexible writer interface - write logs to any destination
- JSON format support for production environment
- Simple and clean API
- Built on top of Go's standard `log/slog` package

## Installation
```bash
go get github.com/ashkenazi1/go-logger
```

## Quick Start
```go
package main

import (
    "log/slog"
    logger "github.com/ashkenazi1/go-logger"
)

func main() {
    log := logger.NewLogger(logger.Config{
        Environment: "development",  // or "production"
        LogLevel:    slog.LevelDebug,
        UseColors:   true,
    })

    log.Debug("Debug message", "user", "john")
    log.Info("Server starting", "port", 8080)
    log.Warn("High CPU usage", "usage", "85%")
    log.Error("Connection failed", "err", "timeout")
}
```

## Custom Writers
The logger supports any `io.Writer` implementation, making it easy to send logs to any destination:

```go
// Write to file
file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
logger := NewLogger(Config{
    Environment: "production",
    LogLevel:    slog.LevelInfo,
    Writer:      file,
})

// Write to multiple destinations
multiWriter := io.MultiWriter(os.Stdout, file)
logger := NewLogger(Config{
    Environment: "development",
    LogLevel:    slog.LevelDebug,
    UseColors:   true,
    Writer:      multiWriter,
})
```

## Configuration
- `Environment`: "development" or "production"
  - development: Human-readable format with optional colors
  - production: JSON format for better parsing
- `LogLevel`: Standard slog levels (DEBUG, INFO, WARN, ERROR)
- `UseColors`: Enable/disable colored output (development mode only)
- `Writer`: Custom io.Writer implementation (defaults to os.Stdout)

## License
MIT License
