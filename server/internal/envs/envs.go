package envs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvs() {

	errLocal := godotenv.Load(".env.local")
	err := godotenv.Load(".env")

	if errLocal != nil && err != nil {

		fmt.Printf("No .env files found\n")

		fmt.Printf("%v\n%v\n", errLocal, err)

		if err != nil {
			panic(err)
		}

		panic(errLocal)
	}
}

const defaultPort = "4000"

func GetPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	return
}

const defaultEnv = "development"

func GetEnvironment() (env string) {
	env = os.Getenv("ENV")
	if env == "" {
		env = defaultEnv
	}
	return
}
