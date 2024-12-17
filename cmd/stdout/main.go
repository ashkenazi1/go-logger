package main

import (
	"log/slog"

	"github.com/ashkenazi1/go-logger"
)

func main() {
	log := logger.NewLogger(logger.Config{
		Environment: "development", // or "production"
		LogLevel:    slog.LevelDebug,
		UseColors:   true,
	})

	log.Debug("Debug message", "user", "john")
	log.Info("Server starting", "port", 8080)
	log.Warn("High CPU usage", "usage", "85%")
	log.Error("Connection failed", "err", "timeout")
}
