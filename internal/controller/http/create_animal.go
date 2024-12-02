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

	// callback := func() error {
	// 	for _, file := range files {
	// 		fileUUID := uuid.New().String()
	// 		if err := ctx.SaveFile(file[0], fmt.Sprintf("%s/%s", c.cfg.StaticFolder, fileUUID)); err != nil {
	// 			logger.Error("can not save file",
	// 				"file", file[0].Filename,
	// 				"err", err.Error(),
	// 			)
	// 			return fiber.NewError(http.StatusInternalServerError, "can not save file")
	// 		}
	// 		animal.PhotoURLs = append(animal.PhotoURLs, fileUUID)
	// 	}

	// 	return nil
	// }
	//

	filesData := make([][]byte, 0, len(fileHeaders))
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader[0].Open()
		if err != nil {
			logger.Error("failed to open file",
				"file", fileHeader[0].Filename,
				"err", err.Error(),
			)
			return fiber.NewError(http.StatusInternalServerError)
		}
		defer file.Close()

		fileData, err := io.ReadAll(file)
		if err != nil {
			logger.Error("failed to open file",
				"file", fileHeader[0].Filename,
				"err", err.Error(),
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
