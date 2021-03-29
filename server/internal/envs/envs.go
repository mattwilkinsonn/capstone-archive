package envs

import "os"

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
