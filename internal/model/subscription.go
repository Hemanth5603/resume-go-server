package model

import "time"

type Subscription struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Plan      int       `json:"plan"`
	Token     string    `json:"token"`
	TimeZone  string    `json:"timezone"`
	CreatedAt time.Time `json:"created_at"`
}

type SubscriptionRequest struct {
	UserID string `json:"user_id"`
	Plan   int    `json:"plan"`
}
