package config

import (
	"github.com/sonix66/animalito-bot/internal/controller/http"
	"github.com/sonix66/animalito-bot/internal/controller/tgbot"
	"github.com/sonix66/animalito-bot/internal/repository/sqlite"
	"github.com/sonix66/animalito-bot/pkg/logger"
)

type (
	Config struct {
		Logger   *logger.Config `mapstructure:"logger"`
		Telegram *tgbot.Config  `mapstructure:"telegram"`
		HTTP     *http.Config   `mapstructure:"http"`
		SQLite   *sqlite.Config `mapstructure:"sqlite"`
	}
)
