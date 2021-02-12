package main

import (
	"os"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/routes"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	orm := db.CreateDatabaseClient()

	g := gin.Default()

	g.POST("/graphql", routes.GraphQLHandler(orm))
	g.GET("/", routes.PlaygroundHandler())
	g.Run("localhost:4000")
}
