package main

import (
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/routes"
	"github.com/Zireael13/capstone-archive/server/internal/serve"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matthewhartstonge/argon2"
)

func main() {
	_ = godotenv.Load(".env")

	port := serve.GetPort()

	orm := db.CreateDatabaseClient()
	argon := argon2.DefaultConfig()

	g := gin.Default()

	routes.AttachRoutes(g, orm, &argon)

	serve.RunServer(g, port)
}
