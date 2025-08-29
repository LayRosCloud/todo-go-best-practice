package middleware

import (
	"leafall/todo-service/internal/services"
	"leafall/todo-service/utils/exceptions"
	"net/http"
	"strings"
)
func AuthorizationMiddleware(service services.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := w.Header().Get("Authorization")
		if authHeader == "" {
			res := exceptions.NewErrorResponse(401, "The token is invalid", "The token is missing")
			exceptions.WriteError(w, res)
			return
		}
		tokenAndType := strings.Split(authHeader, " ")
		typeToken := tokenAndType[0]
		
		if typeToken != "Bearer" {
			res := exceptions.NewErrorResponse(401, "The token is invalid", "The token is type bad")
			exceptions.WriteError(w, res)
			return
		}
		token := tokenAndType[1]
		claims, valid, err := service.ValidateAccessToken(token)
		if err != nil {
			res := exceptions.NewErrorResponse(401, "The token is invalid", err.Error())
			exceptions.WriteError(w, res)
			return
		}

		if !valid {
			res := exceptions.NewErrorResponse(401, "The token is invalid", "The token is invalid")
			exceptions.WriteError(w, res)
			return
		}

		userId := claims["user_id"].(int64)
		r.Header.Set("User-Id", string(userId))

		next.ServeHTTP(w, r)
		})
	}
}
