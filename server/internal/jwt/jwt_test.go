package jwt_test

import (
	"os"
	"testing"

	"github.com/Zireael13/capstone-archive/server/internal/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAndParseToken(t *testing.T) {

	err := os.Setenv("JWT_SECRET_KEY", "testkey")
	require.Nil(t, err)

	username := "zireael12"

	token, err := jwt.CreateToken(username)
	require.Nil(t, err)

	gotUsername, err := jwt.ParseToken(token)

	require.Nil(t, err)
	assert.Equal(t, username, gotUsername)

}
