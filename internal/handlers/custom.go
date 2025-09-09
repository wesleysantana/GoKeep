package handlers

import (
	"errors"
	"net/http"

	"github.com/wesleysantana/GoKeep/internal/apperror"
)

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

func (h HandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		var statusError apperror.StatusError
		if errors.As(err, &statusError) {
			http.Error(w, err.Error(), statusError.StatusCode())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
