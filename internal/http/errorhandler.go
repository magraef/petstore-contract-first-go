package http

import (
	"encoding/json"
	"errors"
	"github.com/magraef/petstore-contract-first-go/internal"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	problemJsonHeaderValue = "application/problem+json"
	contentTypeHeader      = "Content-Type"
)

func ErrorHandler() func(w http.ResponseWriter, r *http.Request, err error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		problem := problemForErr(err, r.URL.Path)
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
			Title:    "Not Found",
			Type:     nil,
		}
	}

	if errors.Is(err, internal.ErrPetAlreadyExists) {
		return Problem{
			Detail:   err.Error(),
			Instance: instance,
			Status:   http.StatusConflict,
			Title:    "Conflict",
			Type:     nil,
		}
	}

	log.Err(err).Msg("Handling unmapped error")

	return Problem{
		Detail:   "Internal Server Error",
		Instance: instance,
		Status:   http.StatusInternalServerError,
		Title:    "Internal Server Error",
		Type:     nil,
	}
}
