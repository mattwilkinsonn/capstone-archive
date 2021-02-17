package main

import (
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/serve"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	port := serve.GetPort()

	orm := db.CreateDefaultDatabaseClient()
	argon := serve.CreateArgon()

	g := serve.CreateServer(orm, argon)

	serve.RunServer(g, port)

}
