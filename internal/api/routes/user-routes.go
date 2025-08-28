package routes

import (
	"leafall/todo-service/internal/handlers"

	"github.com/go-chi/chi/v5"
)



func RegisterUsers(handler *handlers.UserHandler, router chi.Router) {
	router.Get("/v1/users", handler.FindAllPagination)
	router.Get("/v1/users/{id}", handler.FindById)
	router.Post("/v1/users/", handler.Create)
	router.Put("/v1/users/", handler.Update)
	router.Patch("/v1/users/password", handler.UpdatePassword)
	router.Delete("/v1/users/{id}", handler.Delete)
}