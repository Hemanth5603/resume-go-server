package database

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db   *pgxpool.Pool
	once sync.Once
)

// Connect initializes the database connection pool.
// It uses a singleton pattern to ensure only one pool is created.
func Connect(databaseURL string) (*pgxpool.Pool, error) {
	var err error
	once.Do(func() {
		db, err = pgxpool.New(context.Background(), databaseURL)
		if err != nil {
			// Wrapping the error for more context is a good practice,
			// but for now, we'll keep it simple.
			return
		}
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetDB returns the existing database connection pool.
// It's recommended to call Connect once at the application start.
func GetDB() *pgxpool.Pool {
	return db
}
