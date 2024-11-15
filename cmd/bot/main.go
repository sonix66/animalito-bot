package main

import (
	"context"
	"errors"
	"flag"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonix66/animalito-bot/internal/config"
	"github.com/sonix66/animalito-bot/internal/controller/telegram"
	"github.com/sonix66/animalito-bot/internal/repository/postgres"
	announcementservice "github.com/sonix66/animalito-bot/internal/services/announcement_service"
	"github.com/sonix66/animalito-bot/pkg/logger"
	"gopkg.in/telebot.v4"
)

func main() {
	configFile, err := getConfigFile()
	if err != nil {
		panic(err)
	}

	cfg := config.MustInit(configFile)

	logger.MustInitGlobal(cfg.Logger)

	pool, err := pgxpool.New(context.Background(), cfg.Postgres.GetDSN())
	if err != nil {
		panic(err)
	}

	repo := postgres.New(pool)
	announcementService := announcementservice.New(repo)
	controller := telegram.New(announcementService)

	bot, err := telebot.NewBot(telebot.Settings{
		Token: cfg.Telegram.Token,
		Poller: &telebot.LongPoller{
			Timeout: time.Second * time.Duration(cfg.Telegram.PollerTimeoutSeconds),
		},
	})
	if err != nil {
		panic(err)
	}

	telegram.InitHandlers(bot, controller)

	logger.Info("bot is ready to start",
		"token", cfg.Telegram.Token,
		"logger_level", cfg.Logger.Level,
	)

	bot.Start()
}

func getConfigFile() (string, error) {
	configFile := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	if configFile == nil {
		return "", errors.New("config file is required")
	}

	return *configFile, nil
}
