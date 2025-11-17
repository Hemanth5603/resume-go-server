package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system
type User struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username        string             `json:"username" bson:"username"`
	Name            string             `json:"name" bson:"name"`
	Email           string             `json:"email" bson:"email"`
	Password        string             `json:"-" bson:"password"` // Do not expose password in JSON
	ProfileImageURL string             `json:"profileImageUrl" bson:"profile_image_url,omitempty"`
	CreatedAt       time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updated_at"`
}

// CreateUserRequest defines the request body for creating a new user.
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
