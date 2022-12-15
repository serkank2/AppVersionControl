package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Env() bool {

	err := godotenv.Load("./api/.env")
	if err != nil {
		fmt.Printf("Error loading .env file %v", err)
		return false
	}
	return true

}
