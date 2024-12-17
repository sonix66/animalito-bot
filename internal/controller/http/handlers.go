package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitHandlers(app *fiber.App, controller *Controller, cfg *Config) {
	api := app.Group("/api", cors.New())
	app.Use(logger.New())
	animals := api.Group("/animals")
	{
		animals.Post("/", controller.CreateAnimalHandler)
		animals.Get("/", controller.GetListAnimalHandler)
		animals.Get("/count", controller.GetAnimalsCountHandler)
		animals.Get("/:id", controller.GetAnimalHandler)
		animals.Put("/:id", controller.UpdateAnimalHandler)
		animals.Delete("/:id", controller.DeleteAnimalHandler)
	}

	app.Static(cfg.PhotosPrefixURL, cfg.StaticFolder)
}
