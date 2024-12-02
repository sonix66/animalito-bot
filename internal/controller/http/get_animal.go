package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sonix66/animalito-bot/pkg/logger"
)

func (c *Controller) GetAnimalHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		logger.Error("empty id")
		return fiber.NewError(http.StatusBadRequest, "empty id")
	}

	animal, err := c.animalService.GetAnimalByID(ctx.Context(), id)
	if err != nil {
		logger.Error("failed to get animal by ID",
			"id", id,
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusInternalServerError)
	}

	err = ctx.JSON(animal)
	if err != nil {
		logger.Error("failed to send animal",
			"id", id,
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return nil
}
