package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/magraef/petstore-contract-first-go/internal"
	"github.com/rs/zerolog/log"
)

const (
	problemJsonHeaderValue = "application/problem+json"
	contentTypeHeader      = "Content-Type"
)

func RequestErrorHandler() func(w http.ResponseWriter, r *http.Request, err error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		problem := Problem{
			Detail:   err.Error(),
			Instance: r.URL.String(),
			Status:   http.StatusBadRequest,
			Title:    http.StatusText(http.StatusBadRequest),
			Type:     nil,
		}
		w.WriteHeader(problem.Status)
		w.Header().Set(contentTypeHeader, problemJsonHeaderValue)
		json.NewEncoder(w).Encode(problem)
	}
}

func ResponseErrorHandler() func(w http.ResponseWriter, r *http.Request, err error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		problem := problemForErr(err, r.URL.String())
		w.WriteHeader(problem.Status)
		w.Header().Set(contentTypeHeader, problemJsonHeaderValue)
		json.NewEncoder(w).Encode(problem)
	}
}

func problemForErr(err error, instance string) Problem {
	if errors.Is(err, internal.ErrPetNotFound) {
		return Problem{
			Detail:   err.Error(),
			Instance: instance,
			Status:   http.StatusNotFound,
			Title:    http.StatusText(http.StatusNotFound),
			Type:     nil,
		}
	}

	if errors.Is(err, internal.ErrPetAlreadyExists) {
		return Problem{
			Detail:   err.Error(),
			Instance: instance,
			Status:   http.StatusConflict,
			Title:    http.StatusText(http.StatusConflict),
			Type:     nil,
		}
	}

	log.Err(err).Msg("Handling unmapped error")

	return Problem{
		Detail:   http.StatusText(http.StatusInternalServerError),
		Instance: instance,
		Status:   http.StatusInternalServerError,
		Title:    http.StatusText(http.StatusInternalServerError),
		Type:     nil,
	}
}
