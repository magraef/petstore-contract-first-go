package main

import (
	app "github.com/magraef/petstore-contract-first-go"
	"github.com/magraef/petstore-contract-first-go/http"
)

func main() {
	config := app.NewApplicationConfig()

	controller := http.NewPetstoreController()

	http.NewServer(controller, config.Api.BaseUrl, config.Api.Port)
}
