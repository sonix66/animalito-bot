package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sonix66/animalito-bot/pkg/logger"
)

func (c *Controller) GetAnimalsCountHandler(ctx *fiber.Ctx) error {
	count, err := c.animalService.GetAnimalsCount(ctx.Context())
	if err != nil {
		logger.Error("failed to get animal list",
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusInternalServerError)
	}

	err = ctx.JSON(fiber.Map{
		"count": count,
	})
	if err != nil {
		logger.Error("failed to send animals",
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return nil
}
