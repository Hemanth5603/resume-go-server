package handlers

import (
	"github.com/Hemanth5603/resume-go-server/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type MainServiceHandler struct {
	service  service.MainService
	validate *validator.Validate
}

func NewMainServiceHandler(service service.MainService) *MainServiceHandler {
	return &MainServiceHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *MainServiceHandler) ForwardResumeAndDescriptionToModel(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	description := ctx.FormValue("description")
	if description == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "description is required",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.service.ForwardResumeAndDescriptionToModel(file, fileHeader.Filename, description)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "file sent successfully",
	})
}
