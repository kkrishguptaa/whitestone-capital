package lib

import (
	"log"
	"os"
	"strings"
)

type EnvSpec struct {
	API_PORT          string
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
}

var Env EnvSpec

func getEnv() EnvSpec {
	keys := []string{
		"API_PORT",
		"DATABASE_HOST",
		"DATABASE_PORT",
		"DATABASE_USER",
		"DATABASE_PASSWORD",
		"DATABASE_NAME",
	}

	var env map[string]string = make(map[string]string, len(keys))

	var notDefined []string

	for _, key := range keys {
		value := os.Getenv(key)

		if value == "" {
			log.Printf("Warning: Environment variable %s is not set.", key)

			if key == "API_PORT" {
				value = "8080" // Default API port
			} else if key == "DATABASE_HOST" {
				value = "localhost" // Default database host
			} else {
				notDefined = append(notDefined, key)
			}
		}

		env[key] = value
	}

	if len(notDefined) > 0 {
		panic("Environment variables not defined: " + "\n" + strings.Join(notDefined, "\n"))
	}

	return EnvSpec{
		API_PORT:          env["API_PORT"],
		DATABASE_HOST:     env["DATABASE_HOST"],
		DATABASE_PORT:     env["DATABASE_PORT"],
		DATABASE_USER:     env["DATABASE_USER"],
		DATABASE_PASSWORD: env["DATABASE_PASSWORD"],
		DATABASE_NAME:     env["DATABASE_NAME"],
	}
}

func InitEnv() {
	Env = getEnv()
}
