package serve

import (
	"github.com/Zireael13/capstone-archive/server/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateCorsMiddleware() gin.HandlerFunc {
	cors_config := cors.DefaultConfig()
	cors_config.AllowCredentials = true

	origin := config.GetCorsOrigin()
	cors_config.AllowOrigins = []string{origin}

	return cors.New(cors_config)
}
