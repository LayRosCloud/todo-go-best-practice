package api

import (
	"leafall/todo-service/internal/api/routes"
	"leafall/todo-service/internal/handlers"
	"leafall/todo-service/internal/middleware"

	"github.com/go-chi/chi/v5"
)

type ApiSetup struct {
	UserHandler *handlers.UserHandler
}

func NewApiSetup(userHandler *handlers.UserHandler) *ApiSetup {
	return &ApiSetup{
		UserHandler: userHandler,
	}
}

func SetupRoutes(setup *ApiSetup) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.CorrelationMiddleware)
	router.Use(middleware.LoggerMiddleware)

	router.Route("/api", func(r chi.Router) {
		r.Use(middleware.JsonMiddleware)
		routes.RegisterUsers(setup.UserHandler, router)
	})
	return router
}