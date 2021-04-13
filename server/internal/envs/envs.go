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

func GetDatabaseUrl() (url string) {
	url = os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgresql://postgres:postgres@localhost:5432/capstone-archive"
	}
	return
}

func GetRedisAddress() (address string) {
	address = os.Getenv("REDIS_ADDRESS")
	if address == "" {
		address = "localhost:6379"
	}
	return
}
