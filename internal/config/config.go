package config

import (
	"github.com/sonix66/animalito-bot/internal/controller/http"
	"github.com/sonix66/animalito-bot/internal/controller/telegram"
	"github.com/sonix66/animalito-bot/internal/repository/sqlite"
	"github.com/sonix66/animalito-bot/pkg/logger"
)

type (
	Config struct {
		Logger   *logger.Config   `mapstructure:"logger"`
		Telegram *telegram.Config `mapstructure:"telegram"`
		HTTP     *http.Config     `mapstructure:"http"`
		SQLite   *sqlite.Config   `mapstructure:"sqlite"`
	}
)
