package serve

import (
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort = "4000"

func GetPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	return
}

func RunServer(g *gin.Engine, port string) {
	err := g.Run("localhost:" + port)
	if err != nil {
		panic(err)
	}
}
