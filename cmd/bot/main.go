package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sonix66/animalito-bot/internal/config"
	"github.com/sonix66/animalito-bot/internal/controller/http"
	"github.com/sonix66/animalito-bot/internal/controller/telegram"
	"github.com/sonix66/animalito-bot/internal/repository/sqlite"
	animalservice "github.com/sonix66/animalito-bot/internal/services/animal_service"

	_ "github.com/mattn/go-sqlite3"
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

	db, err := sql.Open("sqlite3", cfg.SQLite.FilePath)
	if err != nil {
		panic(err)
	}

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		panic(err)
	}

	repo := sqlite.New(db)
	animalService := animalservice.New(repo, cfg.HTTP.Host, cfg.HTTP.PhotosPrefixURL, cfg.HTTP.StaticFolder)
	tgController := telegram.New(cfg.Telegram, animalService)
	httpController := http.New(animalService, cfg.HTTP)

	bot, err := telebot.NewBot(telebot.Settings{
		Token: cfg.Telegram.Token,
		Poller: &telebot.LongPoller{
			Timeout: time.Second * time.Duration(cfg.Telegram.PollerTimeoutSeconds),
		},
	})
	if err != nil {
		panic(err)
	}

	telegram.InitHandlers(bot, tgController)

	logger.Info("bot is ready to start",
		"token", cfg.Telegram.Token,
		"logger_level", cfg.Logger.Level,
	)

	server := fiber.New()
	http.InitHandlers(server, httpController, cfg.HTTP)

	go bot.Start()
	server.Listen(fmt.Sprintf(":%s", cfg.HTTP.Port))
}

func getConfigFile() (string, error) {
	configFile := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	if configFile == nil {
		return "", errors.New("config file is required")
	}

	return *configFile, nil
}
