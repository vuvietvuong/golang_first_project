package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Setup() {
	fmt.Println("Initial configuration")
}

func PrivateKey() []byte {
	return []byte(os.Getenv("JWT_PRIVATE_KEY"))
}
