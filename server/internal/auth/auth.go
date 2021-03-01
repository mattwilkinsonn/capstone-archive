package auth

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func CreateStore() redis.Store {
	addr := os.Getenv("REDIS_ADDRESS")
	key := os.Getenv("REDIS_AUTH_KEY")
	store, _ := redis.NewStore(10, "tcp", addr, "", []byte(key))

	return store
}

func CreateSessionMiddleware(store redis.Store) gin.HandlerFunc {
	return sessions.Sessions("auth", store)
}

func CreateDefaultSessionMiddleware() gin.HandlerFunc {
	return CreateSessionMiddleware(CreateStore())
}
