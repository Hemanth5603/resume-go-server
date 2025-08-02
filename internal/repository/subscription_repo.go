package repository

import (
	"context"

	"github.com/Hemanth5603/resume-go-server/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SubscriptionRepository struct {
	db *pgxpool.Pool
}

func NewSubscriptionRepository(db *pgxpool.Pool) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) CreateSubscriptionTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS subscriptions (
			id UUID PRIMARY KEY,
			user_id UUID NOT NULL,
			plan INT NOT NULL,
			token TEXT NOT NULL,
			timezone TEXT NOT NULL DEFAULT 'UTC',
			created_at TIMESTAMPTZ DEFAULT now()
		);
		`
	_, err := r.db.Exec(context.Background(), query)
	return err
}

func (r *SubscriptionRepository) CreateSubscription(subscription *model.Subscription) (*model.Subscription, error) {
	query := `
		INSERT INTO subscriptions (id, user_id, plan, token, timezone) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at;
	`
	err := r.db.QueryRow(
		context.Background(),
		query,
		subscription.ID,
		subscription.UserID,
		subscription.Plan,
		subscription.Token,
		subscription.TimeZone,
	).Scan(&subscription.CreatedAt)

	if err != nil {
		return &model.Subscription{}, err
	}

	return subscription, nil
}

func (r *SubscriptionRepository) GetLatestSubscriptionOfUser(UserID string) (*model.Subscription, error) {
	subscription := &model.Subscription{}
	query := `
		SELECT id, user_id, plan, token, timezone, created_at
		FROM subscriptions
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT 1;
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		UserID,
	).Scan(
		&subscription.ID,
		&subscription.UserID,
		&subscription.Plan,
		&subscription.Token,
		&subscription.TimeZone,
		&subscription.CreatedAt,
	)
	if err != nil {
		return &model.Subscription{}, err
	}
	return subscription, nil
}
