package auth

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultOrigin = "http://localhost:3000"

func GetCorsOrigin() (origin string) {
	origin = os.Getenv("CORS_ORIGIN")
	if origin == "" {
		origin = defaultOrigin
	}
	return
}

func CreateCorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()

	origin := GetCorsOrigin()

	config.AllowOrigins = []string{origin}

	return cors.New(config)
}
