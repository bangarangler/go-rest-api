package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVar(key string) string {
	err := godotenv.Load("database.env")
	if err != nil {
		log.Fatalf("error loading env")
	}
	return os.Getenv(key)
}

func main() {
	a := App{}
	a.Initialize(
		// os.Getenv("APP_DB_USERNAME"),
		// os.Getenv("APP_DB_PASSWORD"),
		// os.Getenv("APP_DB_NAME"))
		goDotEnvVar("POSTGRES_USER"),
		goDotEnvVar("POSTGRES_PASSWORD"),
		goDotEnvVar("POSTGRES_DB"))

	a.Run(":8010")
}
