package http

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sonix66/animalito-bot/internal/entity"
	"github.com/sonix66/animalito-bot/pkg/logger"
)

func (c *Controller) CreateAnimalHandler(ctx *fiber.Ctx) error {
	animal := new(entity.Animal)

	err := ctx.BodyParser(animal)
	if err != nil {
		logger.Error("invalid body",
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusUnprocessableEntity, "invalid body")
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		logger.Error("error getting form data",
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusBadRequest, "error getting form data")
	}

	fileHeaders := form.File

	filesData := make([][]byte, 0, len(fileHeaders))
	for _, fileHeader := range fileHeaders {
		file, errOpen := fileHeader[0].Open()
		if errOpen != nil {
			logger.Error("failed to open file",
				"file", fileHeader[0].Filename,
				"err", errOpen.Error(),
			)
			return fiber.NewError(http.StatusInternalServerError)
		}
		defer file.Close()

		fileData, errRead := io.ReadAll(file)
		if errRead != nil {
			logger.Error("failed to open file",
				"file", fileHeader[0].Filename,
				"err", errRead.Error(),
			)
			return fiber.NewError(http.StatusInternalServerError)
		}
		filesData = append(filesData, fileData)
	}

	err = c.animalService.CreateAnimal(ctx.Context(), animal, filesData)
	if err != nil {
		logger.Error("failed to create animal",
			"animal", animal,
			"err", err.Error(),
		)
		return fiber.NewError(http.StatusInternalServerError)
	}
	return nil
}
