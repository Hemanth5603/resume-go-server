package handlers

import "github.com/gofiber/fiber/v2"

// UserHandler handles user-related requests
type UserHandler struct {
	// service service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetUser gets a user
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	return c.SendString("GetUser handler")
}
