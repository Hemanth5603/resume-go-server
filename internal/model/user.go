package model

import "time"

// User represents a user in the system
type User struct {
	ID              string    `json:"id"`
	Username        string    `json:"username"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"-"` // Do not expose password
	ProfileImageURL string    `json:"profileImageUrl"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// CreateUserRequest defines the request body for creating a new user.
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
