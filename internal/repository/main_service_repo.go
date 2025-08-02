package repository

import "github.com/jackc/pgx/v5/pgxpool"

type MainServiceRepository struct {
	db *pgxpool.Pool
}

func NewMainServiceRepo(db *pgxpool.Pool) *MainServiceRepository {
	return &MainServiceRepository{db: db}
}
