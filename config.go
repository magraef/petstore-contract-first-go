package petstore_contract_first_go

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type ApplicationConfig struct {
	Api ApiConfig
}

type ApiConfig struct {
	Port    uint16 `default:"8080"`
	BaseUrl string `split_words:"true" default:"/api"`
}

func NewApplicationConfig() ApplicationConfig {
	var c ApplicationConfig
	if err := envconfig.Process("app", &c); err != nil {
		log.Fatal().Err(err).Msg("failed to load app config")
	}
	return c
}
