package main

import (
	"github.com/Zireael13/capstone-archive/server/internal/auth"
	"github.com/Zireael13/capstone-archive/server/internal/config"
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/fake"
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

	fake.AddFakeCapstonesIfEmpty(queries)

	// REMOVE IF DEPLOYING!!
	fake.AddTestAdminUserIfEmpty(queries)

	g := serve.CreateServer(queries, argon)

	serve.RunServer(g, port)

}
