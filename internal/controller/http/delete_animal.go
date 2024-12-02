package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sonix66/animalito-bot/pkg/logger"
)

func (c *Controller) DeleteAnimalHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		logger.Error("empty id")
		return fiber.NewError(http.StatusBadRequest, "empty id")
	}

	err := c.animalService.DeleteAnimalByID(ctx.Context(), id)
	if err != nil {
		logger.Error("failed to delete animal",
			"id", id,
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusInternalServerError)
	}
	return nil
}
