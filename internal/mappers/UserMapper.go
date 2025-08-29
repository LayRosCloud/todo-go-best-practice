package mappers

import (
	"leafall/todo-service/internal/dto"
	"leafall/todo-service/internal/models"
)

type UserMapper struct {
	TaskMapper *TaskMapper
}

func NewUserMapper(taskMapper *TaskMapper) *UserMapper {
	return &UserMapper{TaskMapper: taskMapper}
}

func (m *UserMapper) MapToFull(user *models.User) (*dto.UserFullResponse) {
	tasks := m.TaskMapper.MapToFullList(user.Tasks);
	return &dto.UserFullResponse{
		Id: user.Id,
		Login: user.Login,
		CreatedAt: user.CreatedAt,
		Tasks: tasks,
	}
}

func (m *UserMapper) MapToShort(user *models.User) (*dto.UserShortResponse) {
	return &dto.UserShortResponse{
		Id: user.Id,
		Login: user.Login,
		CreatedAt: user.CreatedAt,
	}
}

func (m *UserMapper) MapToShortList(users []models.User) ([]dto.UserShortResponse) {
	responses := make([]dto.UserShortResponse, len(users))
	for index, item := range users {
		responses[index] = *m.MapToShort(&item)
	}
	return responses
}

func (m *UserMapper) MapFromCreateRequestToEntity(dto *dto.UserCreateRequest) (*models.User) {
	return &models.User {
		Id: 0,
		Login: dto.Login,
		Password: dto.Password,
	}
}

func (m *UserMapper) AssignFromUpdateRequestToEntity(dto *dto.UserUpdateRequest, user *models.User) {
	user.Login = dto.Login
}