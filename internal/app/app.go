package app

import (
	"portalnews/config"

	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()

	_, err := cfg.ConnectPostgres()
	if err != nil {
		log.Fatal().Msgf("error to connect database %v", err)
		return
	}
}
