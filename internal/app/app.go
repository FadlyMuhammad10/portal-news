package app

import (
	"portalnews/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()

	_, err := cfg.ConnectPostgres()
	if err != nil {
		log.Fatal().Msgf("error to connect database %v", err)
		return
	}

	// cloudflare

	cdfR2 := cfg.LoadAwsConfig()
	_ = s3.NewFromConfig(cdfR2)
}
