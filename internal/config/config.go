package config

import "github.com/sonix66/animalito-bot/internal/repository/postgres"

type (
	TelegramConfig struct {
		Token                string `mapstructure:"token"`
		PollerTimeoutSeconds int    `mapstructure:"poller_timeout_seconds"`
	}

	Config struct {
		LoggerLevel string                   `mapstructure:"logger_level"`
		Telegram    *TelegramConfig          `mapstructure:"telegram"`
		Postgres    *postgres.PostgresConfig `mapstructure:"postgres"`
	}
)
