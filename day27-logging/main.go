package main

import (
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().
		Str("method", "POST").
		Str("path", "/api/auth/login").
		Int("status", 200).
		Dur("duration", 3*time.Millisecond).
		Msg("Request handled")
}
