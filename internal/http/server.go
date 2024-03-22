package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func NewServer(controller PetstoreController, apiBaseURL string, apiPort uint16) {

	httpRouter := newChiHttpRouter(controller, apiBaseURL)
	log.Info().Msgf("starting http server on port %d", apiPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", apiPort), httpRouter)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start http router")
	}
}

func newChiHttpRouter(controller PetstoreController, baseURL string) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/q/health"))

	return HandlerFromMuxWithBaseURL(NewStrictHandler(controller, nil), router, baseURL)
}
