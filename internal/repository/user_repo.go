package repository

import (
	"context"

	"github.com/Hemanth5603/resume-go-server/internal/database"
	"github.com/Hemanth5603/resume-go-server/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.GetDB()}
}

func (r *UserRepository) CreateUserTable() error {
	query := `
	   CREATE TABLE IF NOT EXISTS users (
	       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	       username TEXT NOT NULL UNIQUE,
	       name TEXT NOT NULL,
	       email TEXT NOT NULL UNIQUE,
	       password TEXT NOT NULL,
	       profile_image_url TEXT,
	       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	   );
	   `
	_, err := r.db.Exec(context.Background(), query)
	return err
}

func (r *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	query := `
	   INSERT INTO users (username, name, email, password, profile_image_url)
	   VALUES ($1, $2, $3, $4, $5)
	   RETURNING id, created_at, updated_at
	   `
	err := r.db.QueryRow(context.Background(), query, user.Username, user.Name, user.Email, user.Password, user.ProfileImageURL).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
