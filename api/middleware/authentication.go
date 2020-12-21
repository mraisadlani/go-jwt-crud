package middleware

import (
	"errors"
	"github.com/vanilla/go-jwt-crud/api/payload"
	"github.com/vanilla/go-jwt-crud/api/security"
	"net/http"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		err := security.TokenValid(r)

		if err != nil {
			payload.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		next(w, r)
	}
}