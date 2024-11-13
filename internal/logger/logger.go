package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/sonix66/animalito-bot/internal/config"
)

var logger *slog.Logger

func Init(cfg *config.Config) error {
	var loggerLevel slog.Level
	switch cfg.LoggerLevel {
	case "debug":
		loggerLevel = slog.LevelDebug
	case "error":
		loggerLevel = slog.LevelError
	case "info":
		loggerLevel = slog.LevelInfo
	case "warn":
		loggerLevel = slog.LevelWarn
	default:
		return fmt.Errorf("unknown logger level: %s", cfg.LoggerLevel)
	}

	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: loggerLevel,
	}))

	return nil
}

func DEBUG(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}

func INFO(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

func WARN(msg string, args ...interface{}) {
	logger.Warn(msg, args...)
}

func ERROR(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}
