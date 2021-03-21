package serve

import (
	"os"

	"github.com/Zireael13/capstone-archive/server/internal/auth"
	"github.com/Zireael13/capstone-archive/server/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

const defaultPort = "4000"

func GetPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	return
}

func CreateServer(orm *gorm.DB, argon *argon2.Config) *gin.Engine {
	g := gin.Default()

	g.Use(auth.CreateCorsMiddleware())

	g.Use(auth.GinContextToContextMiddleware())

	g.Use(auth.CreateSessionMiddleware())

	gqlHandler := router.GraphQLHandler(orm, argon)
	playHandler := router.PlaygroundHandler()

	router.AttachRoutes(g, gqlHandler, playHandler)

	return g

}

// TODO: hardcoded to localhost will need to fix in future
func RunServer(g *gin.Engine, port string) {
	err := g.Run("localhost:" + port)
	if err != nil {
		panic(err)
	}
}
