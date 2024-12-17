package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/sonix66/animalito-bot/internal/config"
	"github.com/sonix66/animalito-bot/internal/controller/http"
	"github.com/sonix66/animalito-bot/internal/controller/tgbot"
	"github.com/sonix66/animalito-bot/internal/repository/sqlite"
	animalservice "github.com/sonix66/animalito-bot/internal/services/animal_service"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sonix66/animalito-bot/pkg/logger"
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
	//nolint:govet // not overriding
	if err := db.Ping(); err != nil {
		panic(err)
	}

	repo := sqlite.New(db)
	animalService := animalservice.New(repo, cfg.HTTP.Host, cfg.HTTP.PhotosPrefixURL, cfg.HTTP.StaticFolder)
	tgController := tgbot.New(cfg.Telegram, animalService)
	httpController := http.New(animalService, cfg.HTTP)

	client := tg.New(cfg.Telegram.Token)
	bot := tgb.NewRouter()

	tgbot.InitRoutes(bot, tgController)

	server := fiber.New()
	http.InitHandlers(server, httpController, cfg.HTTP)

	go mustRunBot(context.Background(), bot, client)
	_ = server.Listen(fmt.Sprintf(":%s", cfg.HTTP.Port))
}

func getConfigFile() (string, error) {
	configFile := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	if configFile == nil {
		return "", errors.New("config file is required")
	}

	return *configFile, nil
}

func mustRunBot(ctx context.Context, router *tgb.Router, client *tg.Client) {
	err := tgb.NewPoller(router, client).Run(ctx)
	if err != nil {
		panic(err)
	}
}
