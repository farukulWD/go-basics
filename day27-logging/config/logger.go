package config

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger(env string) {
	zerolog.TimeFieldFormat = time.RFC3339

	if env == "production" {
		// JSON output for log aggregators (Datadog, Loki, CloudWatch)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	} else {
		// Pretty colored output for development
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		})
	}

	log.Info().Str("env", env).Msg("Logger initialized")
}
