package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Pretty console output for development
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	log.Info().Msg("Server starting")
	log.Warn().Msg("This is a warning")
	log.Error().Msg("Something went wrong")
	log.Debug().Msg("Debug details")
}
