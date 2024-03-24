package http

import (
	"fmt"
	"github.com/magraef/petstore-contract-first-go/internal"
	"github.com/magraef/petstore-contract-first-go/internal/http/handler"
	"github.com/swaggest/swgui/v5emb"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func NewServer(controller PetstoreController, readinessChecker internal.ReadinessCheck, apiBaseURL string, apiPort uint16) {

	httpRouter := newChiHttpRouter(controller, readinessChecker, apiBaseURL)
	log.Info().Msgf("starting http server on port %d", apiPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", apiPort), httpRouter)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start http router")
	}
}

func newChiHttpRouter(controller PetstoreController, readinessChecker internal.ReadinessCheck, baseURL string) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/q/health", handler.Health())
	router.Get("/q/readiness", handler.NewReadinessHandler(
		func() error {
			return readinessChecker.Check()
		},
	).Handler())
	router.Get("/openapi", handler.Spec())
	router.Handle("/swagger-ui*", v5emb.New("Petstore", "/openapi", "/swagger-ui"))

	return HandlerFromMuxWithBaseURL(NewStrictHandlerWithOptions(controller, nil, StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  ErrorHandler(),
		ResponseErrorHandlerFunc: ErrorHandler(),
	}), router, baseURL)
}
