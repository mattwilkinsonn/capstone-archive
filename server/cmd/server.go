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

	db_url := config.GetDatabaseUrl()

	queries, err := db.CreateClient(db_url)

	if err != nil {
		panic(err)
	}

	argon := auth.CreateArgon()

	// db.LoadSampleData(orm)

	g := serve.CreateServer(queries, argon)

	serve.RunServer(g, port)

}
