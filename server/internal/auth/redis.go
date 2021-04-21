package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Creates the Redis store
func createRedisStore(address string, key []byte) redis.Store {
	// addr := os.Getenv("REDIS_ADDRESS")
	// key := os.Getenv("REDIS_AUTH_KEY")
	store, err := redis.NewStore(10, "tcp", address, "", key)

	if err != nil {
		panic(err)
	}

	return store
}

// Internal fn that creates session middleware. Consumed inside tests or the CreateSessionMiddleware() fn.
func createSessionMiddleware(store redis.Store) gin.HandlerFunc {
	return sessions.Sessions("auth", store)
}

// Creates Redis Session middleware using the default store.
func CreateRedisSessionMiddleware(address string, key []byte) gin.HandlerFunc {
	return createSessionMiddleware(createRedisStore(address, key))
}
