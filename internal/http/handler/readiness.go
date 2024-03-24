package handler

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

type ReadinessCheck func() error
type ReadinessHandler struct {
	checks []ReadinessCheck
}

func NewReadinessHandler(checks ...ReadinessCheck) ReadinessHandler {
	return ReadinessHandler{checks: checks}
}
func (r ReadinessHandler) Handler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		for _, c := range r.checks {
			if err := c(); err != nil {
				log.Err(err).Msg("readiness check failed")
				writer.WriteHeader(http.StatusServiceUnavailable)
				return
			}
		}

		writer.WriteHeader(http.StatusOK)
		return
	}
}
