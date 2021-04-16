package serve

import (
	"os"

	"github.com/Zireael13/capstone-archive/server/internal/auth"
	"github.com/Zireael13/capstone-archive/server/internal/config"
	"github.com/Zireael13/capstone-archive/server/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

func CreateServer(orm *gorm.DB, argon *argon2.Config) *gin.Engine {
	g := gin.Default()

	g.Use(auth.CreateCorsMiddleware())

	g.Use(auth.GinContextToContextMiddleware())

	addr := config.GetRedisAddress()
	key := os.Getenv("REDIS_AUTH_KEY")

	g.Use(auth.CreateRedisSessionMiddleware(addr, []byte(key)))

	gqlHandler := router.GraphQLHandler(orm, argon)
	playHandler := router.PlaygroundHandler()

	router.AttachRoutes(g, gqlHandler, playHandler)

	return g

}

// TODO: hardcoded to localhost will need to fix in future
func RunServer(g *gin.Engine, port string) {
	err := g.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
