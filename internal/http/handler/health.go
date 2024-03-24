package handler

import "net/http"

func Health() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("{\"Status\": \"Ok\"}"))
	}
}
