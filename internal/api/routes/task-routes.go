package routes

import (
	"leafall/todo-service/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterTasks(handler *handlers.TaskHandler, router chi.Router) {
	router.Get("/v1/users/{userId}/tasks", handler.FindAllByUserIdPagination)
	router.Get("/v1/tasks/{id}", handler.FindById)
	router.Post("/v1/tasks", handler.Create)
	router.Put("/v1/tasks", handler.Update)
	router.Delete("/v1/tasks/{id}", handler.Delete)
}