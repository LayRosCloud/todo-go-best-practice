package services

import (
	"context"
	"leafall/todo-service/internal/dto"
	"leafall/todo-service/internal/mappers"
	"leafall/todo-service/internal/repositories"
	"leafall/todo-service/utils/exceptions"

	"github.com/go-playground/validator/v10"
)

type TaskService struct {
	Repository repositories.TaskRepositoryInterface
	Mapper mappers.TaskMapper
}

var validate = validator.New()

func NewTaskService(repository repositories.TaskRepositoryInterface, mapper mappers.TaskMapper) *TaskService {
	return &TaskService{Repository: repository, Mapper: mapper}
}

func (s *TaskService) FindAllByUserIdPagination(ctx context.Context, dto *dto.TaskFindAllRequest) ([]dto.TaskResponse, int64, error) {
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return nil, 0, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	tasks, totalCount, err := s.Repository.FindAllByUserId(ctx, dto.UserId, dto.Limit, dto.Page)
	if err != nil {
		return nil, 0, err
	}
	return s.Mapper.MapToFullList(tasks), totalCount, nil
}

func (s *TaskService) FindById(ctx context.Context, dto *dto.TaskByIdRequest) (*dto.TaskResponse, error) {
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return nil, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	task, err := s.Repository.FindById(ctx, dto.Id)
	if err != nil {
		return nil, exceptions.NewNotFound("Task", err.Error(), dto.Id)
	}
	return s.Mapper.MapToFull(task), nil
}

func (s *TaskService) Create(ctx context.Context, dto *dto.TaskCreateRequest) (*dto.TaskResponse, error) {
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return nil, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	task := s.Mapper.MapToEntity(dto)
	err = s.Repository.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return s.Mapper.MapToFull(task), nil
}

func (s *TaskService) Update(ctx context.Context, dto *dto.TaskUpdateRequest) (*dto.TaskResponse, error) {
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return nil, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	task, err := s.Repository.FindById(ctx, dto.Id)
	if err != nil {
		return nil, exceptions.NewNotFound("Task", err.Error(), dto.Id)
	}
	s.Mapper.AssignUpdateAndEntity(task, dto)
	err = s.Repository.Update(ctx, task)
	if err != nil {
		return nil, err
	}
	return s.Mapper.MapToFull(task), nil
}

func (s *TaskService) Delete(ctx context.Context, dto *dto.TaskByIdRequest) (bool, error) {
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return false, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	err = s.Repository.DeleteById(ctx, dto.Id)
	
	return err == nil, nil
}