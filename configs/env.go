package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvString(env string) string {
	err := godotenv.Load()

	if err != nil {
		log.Println("ERROR env ", err.Error())
	}

	return os.Getenv(env)
}
