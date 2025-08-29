package api

import (
	"leafall/todo-service/internal/api/routes"
	"leafall/todo-service/internal/handlers"
	"leafall/todo-service/internal/middleware"
	"leafall/todo-service/internal/services"

	"github.com/go-chi/chi/v5"
)

type ApiSetup struct {
	UserHandler *handlers.UserHandler
	TaskHandler *handlers.TaskHandler
	UserService *services.UserService
}

func NewApiSetup(userHandler *handlers.UserHandler, taskHandler *handlers.TaskHandler, userService *services.UserService) *ApiSetup {
	return &ApiSetup{
		UserHandler: userHandler,
		TaskHandler: taskHandler,
		UserService: userService,
	}
}

func SetupRoutes(setup *ApiSetup) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.CorrelationMiddleware)
	router.Use(middleware.LoggerMiddleware)
	router.Use(middleware.JsonMiddleware)
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthorizationMiddleware(*setup.UserService))
		routes.RegisterUsers(setup.UserHandler, r)
		routes.RegisterTasks(setup.TaskHandler, r)
	})

	router.Group(func(r chi.Router) {
		routes.RegisterPublicUsers(setup.UserHandler, r)
	})
	
	return router
}