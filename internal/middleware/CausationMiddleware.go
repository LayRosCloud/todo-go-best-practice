package middleware

import (
	"net/http"

	"github.com/beevik/guid"
)

const HeaderCorrelationId = "X-Correlation-Id"
const HeaderCausationId = "X-Causation-Id"

func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationId := r.Header.Get(HeaderCorrelationId)
		
		if correlationId == "" {
			correlationId = guid.New().String()
			r.Header.Set(HeaderCorrelationId, correlationId)
		}
		w.Header().Set(HeaderCorrelationId, correlationId)

		causationId := r.Header.Get(HeaderCausationId)
		if causationId == "" {
			causationId = "start_todo_service"
			r.Header.Set(HeaderCausationId, causationId)
		}
		next.ServeHTTP(w, r)
	})
}