package handlers

import "net/http"

func HealthHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("{\"Status\": \"Ok\"}"))
	}
}
