package serve

import (
	"os"

	"github.com/Zireael13/capstone-archive/server/internal/auth"
	"github.com/Zireael13/capstone-archive/server/internal/config"
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/pack"
	"github.com/Zireael13/capstone-archive/server/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

func CreateServer(queries *db.Queries, argon *argon2.Config) *gin.Engine {
	g := gin.Default()

	g.Use(CreateCorsMiddleware())

	g.Use(pack.GinContextToContextMiddleware())

	addr := config.GetRedisAddress()
	key := os.Getenv("REDIS_AUTH_KEY")

	g.Use(auth.CreateRedisSessionMiddleware(addr, []byte(key)))

	gqlHandler := router.GraphQLHandler(queries, argon)
	playHandler := router.PlaygroundHandler()

	router.AttachRoutes(g, gqlHandler, playHandler)

	return g

}

func RunServer(g *gin.Engine, port string) {
	err := g.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
