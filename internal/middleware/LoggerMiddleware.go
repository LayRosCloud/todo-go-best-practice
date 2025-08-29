package middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		causationId := r.Header.Get(HeaderCausationId)
		correlationId := r.Header.Get(HeaderCorrelationId)
		log.Printf("correlationId: %s,causationId: %s, %s %s", correlationId, causationId, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}