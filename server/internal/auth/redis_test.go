package auth

import (
	"crypto/rand"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/require"
)

func createRedisMock() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	// defer s.Close()

	return s
}

func createRandomRedisKey() []byte {
	token := make([]byte, 32)
	rand.Read(token)
	return token
}

func TestCreateRedisStore(t *testing.T) {
	minredis := createRedisMock()

	token := createRandomRedisKey()

	redisStore := createRedisStore(minredis.Addr(), token)

	require.NotNil(t, redisStore)

}

func TestCreateSessionMiddleware(t *testing.T) {
	minredis := createRedisMock()
	token := createRandomRedisKey()
	store := createRedisStore(minredis.Addr(), token)

	handler := createSessionMiddleware(store)

	require.NotNil(t, handler)
}

func TestCreateRedisSessionMiddleware(t *testing.T) {
	minredis := createRedisMock()
	token := createRandomRedisKey()

	handler := CreateRedisSessionMiddleware(minredis.Addr(), token)

	require.NotNil(t, handler)
}
