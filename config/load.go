package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

var isInitializedSuccessfully = false

func Load() {
	err := godotenv.Load()

	if err != nil {
		isInitializedSuccessfully = false
		fmt.Printf("Unable to load dotenv: %s\n", err)
		return
	}

	fmt.Printf("Loading config\n")

	initPostgresConfig()

	isInitializedSuccessfully = true
}

func IsInitialized() bool {
	return isInitializedSuccessfully
}
