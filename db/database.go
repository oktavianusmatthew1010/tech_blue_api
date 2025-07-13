package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"honnef.co/go/tools/config"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase(config *config.Config) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(config.SupabaseDBURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	// Test the connection
	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	log.Println("Successfully connected to Supabase database")
	return &Database{Pool: pool}, nil
}

func (db *Database) Close() {
	db.Pool.Close()
}

// Query helper function
func (db *Database) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return db.Pool.QueryRow(ctx, sql, args...)
}

func (db *Database) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return db.Pool.Query(ctx, sql, args...)
}

func (db *Database) Exec(ctx context.Context, sql string, args ...interface{}) (pgxconn.CommandTag, error) {
	return db.Pool.Exec(ctx, sql, args...)
}
