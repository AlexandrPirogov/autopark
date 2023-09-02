package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// configurable variable
var (
	redis_pwd     string
	postgres_user string
	postgres_pwd  string
	postgres_db   string

	log_level string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	redis_pwd = os.Getenv("REDIS_PASSWORD")

	postgres_user = os.Getenv("POSTGRES_USER")
	postgres_pwd = os.Getenv("POSTGRES_PASSWORD")
	postgres_db = os.Getenv("POSTGRES_DB")

	log_level = os.Getenv("LOG_LEVEL")

	levels := LogLevelMap()
	if _, ok := levels[log_level]; !ok {
		log.Fatal().Msg("there are only debug|info|warn log levels")
	}

	zerolog.SetGlobalLevel(levels[log_level])
	log.Info().Msgf("Log level is %s", levels[log_level])
}

// returns redis password
func RedisPwd() string {
	return redis_pwd
}

// returns postgres url to connect to
func PostgresURL() string {
	return "postgresql://" + postgres_user + ":" + postgres_pwd + "@postgres-booking:5432/" + postgres_db
}

func LogLevelMap() map[string]zerolog.Level {
	return map[string]zerolog.Level{
		"debug":   zerolog.DebugLevel,
		"info":    zerolog.InfoLevel,
		"warning": zerolog.WarnLevel,
	}
}
