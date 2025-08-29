package routes

import (
	"leafall/todo-service/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterUsers(handler *handlers.UserHandler, router chi.Router) {
	router.Get("/v1/users", handler.FindAllPagination)
	router.Get("/v1/users/{id}", handler.FindById)
	router.Put("/v1/users/", handler.Update)
	router.Patch("/v1/users/password", handler.UpdatePassword)
	router.Delete("/v1/users/{id}", handler.Delete)
}

func RegisterPublicUsers(handler *handlers.UserHandler, router chi.Router) {
	router.Post("/v1/users", handler.Create)
	router.Post("/v1/users/logout", handler.Logout)
	router.Post("/v1/users/access", handler.GenerateAccessToken)
	router.Post("/v1/users/signin", handler.Authorizate)
}