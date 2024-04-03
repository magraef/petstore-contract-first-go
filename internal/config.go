package internal

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type ApplicationConfig struct {
	Api        ApiConfig
	Postgresql PostgresqlConfig
}

type ApiConfig struct {
	Port    uint16 `default:"8080"`
	BaseUrl string `split_words:"true" default:"/api"`
}

type PostgresqlConfig struct {
	Url      string `default:"postgres://postgres:admin@localhost:5432"`
	Database string `default:"petstore"`
}

func NewApplicationConfig() ApplicationConfig {
	var c ApplicationConfig
	if err := envconfig.Process("app", &c); err != nil {
		log.Fatal().Err(err).Msg("failed to load app config")
	}
	return c
}
