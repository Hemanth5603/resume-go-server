package handlers

import (
	"fmt"

	"github.com/Hemanth5603/resume-go-server/internal/model"
	"github.com/Hemanth5603/resume-go-server/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SubscriptionHandler struct {
	service  service.SubscriptionService
	validate *validator.Validate
}

func NewSubscriptionHandler(service service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *SubscriptionHandler) CreateSubscription(ctx *fiber.Ctx) error {
	var req model.SubscriptionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	if err := h.validate.Struct(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	subscription, err := h.service.CreateSubscription(&req)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create subscription",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(subscription)
}
