package middleware

import (
	"net/http"

	"github.com/beevik/guid"
)



func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationId := guid.New()
		if r.Header.Get("X-Correlation-Id") == "" {
			r.Header.Set("X-Correlation-Id", correlationId.String())
			w.Header().Set("X-Correlation-Id", correlationId.String())
		}
		next.ServeHTTP(w, r)
	})
}