package middleware

import (
	"net/http"
	log "github.com/sirupsen/logrus"
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s \n", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}
