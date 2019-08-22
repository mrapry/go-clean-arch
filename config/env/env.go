package env

import (
	"os"
	"strconv"
)

type Environment struct {
	Environment string
	Development string
	HTTPPort    int
}

var environment *Environment

// GetEnv return env
func GetEnv() *Environment {
	return environment
}

// Initial Environtment
func InitEnvironment() {

	// apps environment
	DefaultPort := 8080

	// set REST port
	var port int
	if portEnv, ok := os.LookupEnv("HTTP_PORT"); ok {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			port = DefaultPort
		} else {
			port = int(portInt)
		}
	} else {
		port = DefaultPort
	}

	// set the environment
	environmentInit := &Environment{
		// apps environment
		Environment: os.Getenv("ENVIRONMENT"),
		Development: os.Getenv("DEVELOPMENT"),
		HTTPPort:    port,
	}

	// set the environment
	environment = environmentInit

}