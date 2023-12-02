package utils

import (
	"encoding/json"
	"net/http"

	"github.com/Milad75Rasouli/IEC61850CRUD/model"
)

type ErrorMiddleware func(http.ResponseWriter, *http.Request) error

func ErrorHandler(e ErrorMiddleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := e(w, r)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, model.ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
