package mappers

import (
	"leafall/todo-service/internal/dto"
	"leafall/todo-service/internal/models"
)

type TaskMapper struct {
}

func (m *TaskMapper) MapToFull(task *models.Task) (*dto.TaskResponse) {
	return &dto.TaskResponse{
		Id: task.Id,
		Name: task.Name,
		Description: task.Description,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		UserId: task.UserId,
		FinishedAt: task.FinishedAt,
	}
}

func (m *TaskMapper) MapToFullList(tasks []models.Task) ([]dto.TaskResponse) {
	responses := make([]dto.TaskResponse, len(tasks))
	for index, item := range tasks {
		responses[index] = *m.MapToFull(&item)
	}
	return responses
}

func (m *TaskMapper) AssignUpdateAndEntity(task *models.Task, dto *dto.TaskUpdateRequest) {
	task.Name = dto.Name
	task.Description = dto.Description
	task.FinishedAt = dto.FinishedAt
}

func (m *TaskMapper) MapToEntity(dto *dto.TaskCreateRequest) (*models.Task) {
	return &models.Task{
		Name: dto.Name,
		Description: dto.Description,
		UserId: dto.UserId,
	}
}