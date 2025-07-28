package handlers

import (
	"github.com/Hemanth5603/resume-go-server/internal/model"
	"github.com/Hemanth5603/resume-go-server/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// UserHandler handles user-related requests
type UserHandler struct {
	service  service.UserService
	validate *validator.Validate
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service:  service,
		validate: validator.New(),
	}
}

// GetUser gets a user
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	return c.SendString("GetUser handler")
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req model.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	if err := h.validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.service.CreateUser(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
