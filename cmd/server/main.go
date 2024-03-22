package main

import (
	"github.com/magraef/petstore-contract-first-go/internal"
	"github.com/magraef/petstore-contract-first-go/internal/http"
	"github.com/magraef/petstore-contract-first-go/internal/postgresql"
	"github.com/magraef/petstore-contract-first-go/internal/usecase"
)

func main() {
	config := internal.NewApplicationConfig()

	repository := postgresql.NewRepository(postgresql.NewPgxPool(config.Postgresql.Url, config.Postgresql.Database))
	defer repository.Close()

	petUseCase := usecase.NewPetUseCase(repository)

	controller := http.NewPetstoreController(petUseCase)

	http.NewServer(controller, config.Api.BaseUrl, config.Api.Port)
}
