package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// configurable variable
var (
	log_level string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	log_level = os.Getenv("LOG_LEVEL")

	levels := logLevelMap()
	if _, ok := levels[log_level]; !ok {
		log.Fatal().Msg("there are only debug|info|warn log levels")
	}

	zerolog.SetGlobalLevel(levels[log_level])
	log.Info().Msgf("Log level is %s", levels[log_level])
}

func logLevelMap() map[string]zerolog.Level {
	return map[string]zerolog.Level{
		"debug":   zerolog.DebugLevel,
		"info":    zerolog.InfoLevel,
		"warning": zerolog.WarnLevel,
	}
}
