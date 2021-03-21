package auth

import "github.com/matthewhartstonge/argon2"

func CreateArgon() *argon2.Config {
	argon := argon2.DefaultConfig()
	return &argon
}
