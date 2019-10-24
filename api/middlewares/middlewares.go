package middlewares

import (
	"errors"
	"net/http"

	"github.com/abejide001/Store_Manager_GO/api/auth"
	"github.com/abejide001/Store_Manager_GO/api/responses"
)

// SetMiddlewareJSON method
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // set the content type to json
		next(w, r)
	}
}

// SetMiddlewareAuthentication method
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.TokenValid(r); err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
