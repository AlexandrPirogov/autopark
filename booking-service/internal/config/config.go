package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// configurable variable
var (
	redis_pwd     string
	postgres_user string
	postgres_pwd  string
	postgres_db   string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redis_pwd = os.Getenv("REDIS_PASSWORD")
	postgres_user = os.Getenv("POSTGRES_USER")
	postgres_pwd = os.Getenv("POSTGRES_PASSWORD")
	postgres_db = os.Getenv("POSTGRES_DB")
}

// returns redis password
func RedisPwd() string {
	return redis_pwd
}

// returns postgres url to connect to
func PostgresURL() string {
	return "postgresql://" + postgres_user + ":" + postgres_pwd + "@postgres-booking:5432/" + postgres_db
}
