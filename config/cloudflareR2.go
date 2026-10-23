package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/rs/zerolog/log"
)

func (cfg Config) LoadAwsConfig() aws.Config {
	conf, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.CloudR2.ApiKey, cfg.CloudR2.ApiSecret, "")), awsConfig.WithRegion("auto"))

	if err != nil {
		log.Fatal().Msgf("Unable load aws config %v", err)
	}

	log.Info().Msg("Success load aws config")

	return conf
}
