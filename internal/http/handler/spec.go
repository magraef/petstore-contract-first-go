package handler

import (
	"github.com/magraef/petstore-contract-first-go/docs"
	"net/http"
)

func Spec() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		spec, err := docs.OpenApiSpec()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = writer.Write(*spec)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}
