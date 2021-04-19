package dbtest

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/stretchr/testify/require"
)

func CreateMockDBClient(t *testing.T) (*db.Queries, sqlmock.Sqlmock) {
	t.Helper()

	conn, mock := createSqlMock(t)

	queries := db.CreateQueries(conn)

	return queries, mock
}

func createSqlMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()

	require.Nil(t, err)

	return db, mock
}
