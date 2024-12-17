package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/ashkenazi1/go-logger"
)

// Example custom writer
type TelegramWriter struct {
	botToken string
	chatID   string
}

func (tw *TelegramWriter) Write(p []byte) (n int, err error) {
	// Send to Telegram
	// Implementation here
	return len(p), nil
}

func main() {
	// Use it
	telegramWriter := &TelegramWriter{
		botToken: "your-bot-token",
		chatID:   "your-chat-id",
	}

	multiWriter := io.MultiWriter(os.Stdout, telegramWriter)

	logger := logger.NewLogger(logger.Config{
		Environment: "production",
		LogLevel:    slog.LevelInfo,
		Writer:      multiWriter,
	})

	logger.Debug("Debug message", "user", "john")
	logger.Info("Server starting", "port", 8080)
	logger.Warn("High CPU usage", "usage", "85%")
	logger.Error("Connection failed", "err", "timeout")
}
