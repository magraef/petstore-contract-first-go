package http

import (
	"fmt"
	"net/http"

	"github.com/magraef/petstore-contract-first-go/internal"
	"github.com/magraef/petstore-contract-first-go/internal/http/handlers"
	"github.com/magraef/petstore-contract-first-go/internal/http/middlewares"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggest/swgui/v5emb"

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
	router.Use(middleware.RequestID, middleware.Recoverer)

	//prometheus
	router.Use(middlewares.NewPatternMiddleware("petstore"))
	router.Handle("/metrics", promhttp.Handler())

	// health & ready checks
	router.Get("/q/health", handlers.HealthHandler())
	router.Get("/q/readiness", handlers.NewReadinessHandler(
		func() error {
			return readinessChecker.Check()
		},
	).Handler())

	router.Get("/openapi", handlers.OpenApiSpecHandler())
	router.Handle("/swagger-ui*", v5emb.New("Petstore", "/openapi", "/swagger-ui"))

	return HandlerFromMuxWithBaseURL(NewStrictHandlerWithOptions(controller, nil,
		StrictHTTPServerOptions{
			RequestErrorHandlerFunc:  RequestErrorHandler(),
			ResponseErrorHandlerFunc: ResponseErrorHandler(),
		}),
		router, baseURL)
}
