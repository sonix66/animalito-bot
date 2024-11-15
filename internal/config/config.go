package config

import (
	"github.com/sonix66/animalito-bot/internal/repository/postgres"
	"github.com/sonix66/animalito-bot/pkg/logger"
)

type (
	TelegramConfig struct {
		Token                string `mapstructure:"token"`
		PollerTimeoutSeconds int    `mapstructure:"poller_timeout_seconds"`
	}

	Config struct {
		Logger   *logger.Config   `mapstructure:"logger"`
		Telegram *TelegramConfig  `mapstructure:"telegram"`
		Postgres *postgres.Config `mapstructure:"postgres"`
	}
)
