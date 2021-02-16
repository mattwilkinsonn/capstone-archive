package main

import (
	"os"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	orm := db.CreateDatabaseClient()
	argon := argon2.DefaultConfig()

	g := gin.Default()

	g.POST("/graphql", routes.GraphQLHandler(orm, &argon))
	g.GET("/", routes.PlaygroundHandler())
	g.Run("localhost:4000")
}
