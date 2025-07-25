package model

// User represents a user in the system
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // Do not expose password
}
