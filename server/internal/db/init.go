package db

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v4"
)

func CreateSqlConn(db_url string) (*sql.DB, error) {
	conn, err := sql.Open("pgx", db_url)

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

func createPgxConn(db_url string) *pgx.Conn, error {
	
	conn, err := pgx.Connect(context.Background(), db_url)

	return &conn, err
}