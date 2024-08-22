package lib

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	SECRET_JWT = getEnv("SECRET_JWT")
	DATABASE   = getEnv("DATABASE")
)

func getEnv(env string) string {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv(env)
}
