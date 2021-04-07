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

func CreateArgon() *argon2.Config {
	argon := argon2.DefaultConfig()
	return &argon
}

func CreateServer(orm *gorm.DB, argon *argon2.Config) *gin.Engine {
	g := gin.Default()

	g.Use(auth.GinContextToContextMiddleware())

	addr := os.Getenv("REDIS_ADDRESS")
	key := os.Getenv("REDIS_AUTH_KEY")

	g.Use(auth.CreateRedisSessionMiddleware(addr, []byte(key)))

	gqlHandler := router.GraphQLHandler(orm, argon)
	playHandler := router.PlaygroundHandler()

	router.AttachRoutes(g, gqlHandler, playHandler)

	return g

}

func RunServer(g *gin.Engine, port string) {
	err := g.Run("localhost:" + port)
	if err != nil {
		panic(err)
	}
}
