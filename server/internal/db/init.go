package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sethvargo/go-retry"
)

func CreateSqlConn(db_url string) (*sql.DB, error) {
	connConfig, err := pgx.ParseConfig(db_url)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration("60s")
	if err != nil {
		return nil, err
	}

	connConfig.ConnectTimeout = duration

	connStr := stdlib.RegisterConnConfig(connConfig)

	conn, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	// backoff handling, database probably starting up still
	ctx := context.Background()
	backoff, err := retry.NewFibonacci(1 * time.Second)

	if err != nil {
		return nil, err
	}

	backoff = retry.WithMaxRetries(10, backoff)
	retry.Do(ctx, backoff, func(ctx context.Context) error {
		if err := conn.PingContext(ctx); err != nil {
			// This marks the error as retryable
			return retry.RetryableError(err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return conn, nil

}

func CreateQueries(conn *sql.DB) *Queries {
	queries := New(conn)
	return queries
}

func CreateClient(db_url string) (*Queries, error) {
	conn, err := CreateSqlConn(db_url)
	queries := CreateQueries(conn)

	return queries, err
}

func createPgxConn(db_url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), db_url)

	return conn, err
}
