package http

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sonix66/animalito-bot/internal/entity"
	"github.com/sonix66/animalito-bot/pkg/logger"
)

func (c *Controller) GetListAnimalHandler(ctx *fiber.Ctx) error {
	var (
		count, offset int
	)

	queryCount := ctx.Query("count")
	if queryCount != "" {
		count, _ = strconv.Atoi(queryCount)
	}

	queryOffset := ctx.Query("offset")
	if queryOffset != "" {
		offset, _ = strconv.Atoi(queryOffset)
	}

	if count == 0 {
		count = 10
	}

	animals, err := c.animalService.GetAnimalList(ctx.Context(), count, offset, entity.PhotoSourceURL)
	if err != nil {
		logger.Error("failed to get animal list",
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusInternalServerError)
	}

	err = ctx.JSON(animals)
	if err != nil {
		logger.Error("failed to send animals",
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return nil
}
