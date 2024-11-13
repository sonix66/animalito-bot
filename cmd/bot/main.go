package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonix66/animalito-bot/internal/config"
	"github.com/sonix66/animalito-bot/internal/controller/telegram"
	"github.com/sonix66/animalito-bot/internal/logger"
	"github.com/sonix66/animalito-bot/internal/repository/postgres"
	announcementservice "github.com/sonix66/animalito-bot/internal/services/announcement_service"
	"gopkg.in/telebot.v4"
)

func main() {
	configFile, err := getConfigFile()
	if err != nil {
		panic(err)
	}

	cfg, err := config.InitConfig(configFile)
	if err != nil {
		panic(err)
	}

	err = logger.Init(cfg)
	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.New(context.Background(), getDSNFromConfig(cfg))
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

	logger.INFO("bot is ready to start",
		"token", cfg.Telegram.Token,
		"logger_level", cfg.LoggerLevel,
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

func getDSNFromConfig(config *config.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.Database)
}

func initLogger(cfg *config.Config) *slog.Logger {
	loggerLevel := slog.LevelInfo
	switch cfg.LoggerLevel {
	case "debug":
		loggerLevel = slog.LevelDebug
	case "error":
		loggerLevel = slog.LevelError
	case "info":
		loggerLevel = slog.LevelInfo
	case "warn":
		loggerLevel = slog.LevelWarn
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: loggerLevel,
	}))
}
