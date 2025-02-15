package cmd

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "main.go/internal/sqlc/generate"
)

var QueriesPool *sqlc.Queries

// InitDB initializes the database connections as required
// It currently pgxpool and sqlc for PostgreSQL, but can be changed as per need.
func InitDB() error {

	ctx := context.Background()

	dbConnStr, exists := os.LookupEnv("PGDBLoginCredentials")
	if !exists {
		return errors.New("PGDBLoginCredentials not found")
	}

	pgPool, err := pgxpool.New(ctx, dbConnStr)
	if err != nil {
		return err
	}

	QueriesPool = sqlc.New(pgPool)

	return nil
}