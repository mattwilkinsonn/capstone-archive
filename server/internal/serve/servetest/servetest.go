package servetest

import (
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zireael13/capstone-archive/server/internal/auth"
	"github.com/Zireael13/capstone-archive/server/internal/db/dbtest"
	"github.com/Zireael13/capstone-archive/server/internal/serve"
	"github.com/gin-gonic/gin"
)

func BuildMockServer(t *testing.T) (*gin.Engine, sqlmock.Sqlmock) {
	t.Helper()

	db, mock := dbtest.CreateMockDBClient(t)

	argon := auth.CreateArgon()

	return serve.CreateServer(db, argon), mock
}

func BuildTestServer(t *testing.T, engine *gin.Engine) *httptest.Server {
	return httptest.NewServer(engine)
}
