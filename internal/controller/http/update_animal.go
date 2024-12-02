package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sonix66/animalito-bot/internal/entity"
	"github.com/sonix66/animalito-bot/pkg/logger"
)

func (c *Controller) UpdateAnimalHandler(ctx *fiber.Ctx) error {
	animal := new(entity.Animal)

	err := ctx.BodyParser(animal)
	if err != nil {
		logger.Error("invalid body",
			"id", animal.ID,
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusUnprocessableEntity, "invalid body")
	}

	animal.ID = ctx.Params("id")

	err = c.animalService.UpdateAnimalByID(ctx.Context(), animal)
	if err != nil {
		logger.Error("failed to update animal",
			"animal", animal,
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return nil
}
