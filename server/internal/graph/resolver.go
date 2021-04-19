package graph

import (
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/matthewhartstonge/argon2"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Queries *db.Queries
	Argon   *argon2.Config
}
