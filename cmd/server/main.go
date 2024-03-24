package main

import (
	"github.com/magraef/petstore-contract-first-go/internal"
	"github.com/magraef/petstore-contract-first-go/internal/http"
	"github.com/magraef/petstore-contract-first-go/internal/persistence/postgresql"
)

func main() {
	config := internal.NewApplicationConfig()

	repository := postgresql.NewRepository(postgresql.NewPgxPool(config.Postgresql.Url, config.Postgresql.Database))
	defer repository.Close()

	petsResourceImpl := internal.NewPetResourceImpl(repository)

	controller := http.NewPetstoreController(petsResourceImpl)

	http.NewServer(controller, repository, config.Api.BaseUrl, config.Api.Port)
}
