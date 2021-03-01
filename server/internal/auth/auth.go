package auth

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Cretes the Redis store. Requires the .env variables to be set.
func createStore() redis.Store {
	addr := os.Getenv("REDIS_ADDRESS")
	key := os.Getenv("REDIS_AUTH_KEY")
	store, _ := redis.NewStore(10, "tcp", addr, "", []byte(key))

	return store
}

// Internal fn that creates session middleware. Consumed inside tests or the CreateSessionMiddleware() fn.
func createSessionMiddleware(store redis.Store) gin.HandlerFunc {
	return sessions.Sessions("auth", store)
}

// Creates Session middleware using the default store.
func CreateSessionMiddleware() gin.HandlerFunc {
	return createSessionMiddleware(createStore())
}
