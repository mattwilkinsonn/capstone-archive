package dbtest

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateMockDBClient(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	t.Helper()

	DB, mock := CreateSQLMock(t)

	dialector := CreateMockDialector(t, DB)
	//mock.ExpectQuery("CREATE TABLE 'users'").WithArgs(2, 3)
	gdb := db.CreateDatabaseClient(dialector)

	return gdb, mock
}

func CreateSQLMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.Equal(t, nil, err)

	return db, mock
}

func CreateMockDialector(t *testing.T, db *sql.DB) gorm.Dialector {
	t.Helper()
	return postgres.New(postgres.Config{Conn: db})
}
