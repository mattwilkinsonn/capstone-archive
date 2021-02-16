package main

import (
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/helpers"
	"github.com/Zireael13/capstone-archive/server/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matthewhartstonge/argon2"
)

func main() {
	_ = godotenv.Load(".env")

	port := helpers.GetPort()

	orm := db.CreateDatabaseClient()
	argon := argon2.DefaultConfig()

	g := gin.Default()

	routes.AttachRoutes(g, orm, &argon)

	helpers.RunServer(g, port)
}
