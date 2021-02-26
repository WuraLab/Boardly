package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadENV() error {
	//check the environment else set default
	env := "dev"
	if os.Getenv("ENV") != "" {
		env = os.Getenv("ENV")
	}
	envFile := ".env." + env + ".local"
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error Loading env file %s", envFile)
		return err
	}
	return nil
}
