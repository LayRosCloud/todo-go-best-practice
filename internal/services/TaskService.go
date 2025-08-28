package services

import (
	"leafall/todo-service/internal/repositories"
)

type TaskService struct {
	Repository repositories.TaskRepositoryInterface
}

func NewTaskService(repository repositories.TaskRepositoryInterface) *TaskService {
	return &TaskService{Repository: repository}
}

func FindAllPagination() {
	
}