package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVariables struct {
	APP_NAME    *string
	APP_VERSION *string
}

// NewEnvironmentVariables creates a new instance of EnvironmentVariables
func NewEnvironmentVariables() *EnvironmentVariables {
	env := &EnvironmentVariables{}
	env.initEnv()
	return env
}

func (e *EnvironmentVariables) initEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
func (e *EnvironmentVariables) GetEnv(key string) string {
	return os.Getenv(key)
}
