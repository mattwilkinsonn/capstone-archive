package main

import (
	"github.com/Zireael13/capstone-archive/server/internal/auth"
	"github.com/Zireael13/capstone-archive/server/internal/config"
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/serve"
)

func main() {
	config.LoadEnvs()

	port := config.GetPort()

	orm := db.CreateDefaultDatabaseClient()
	argon := auth.CreateArgon()

	db.LoadSampleData(orm)

	g := serve.CreateServer(orm, argon)

	serve.RunServer(g, port)

}
