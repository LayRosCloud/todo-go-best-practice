package services

import (
	"context"
	"errors"
	"leafall/todo-service/internal/dto"
	"leafall/todo-service/internal/mappers"
	"leafall/todo-service/internal/repositories"
	"leafall/todo-service/utils"
	"leafall/todo-service/utils/exceptions"
)

type UserService struct {
	Repository repositories.UserRepositoryInterface
	UserMapper mappers.UserMapper
}

func NewUserService(repository repositories.UserRepositoryInterface, userMapper mappers.UserMapper) *UserService {
	return &UserService{Repository: repository, UserMapper: userMapper}
}

func (u *UserService) FindAllPagination(ctx context.Context, dto *dto.UserFindAllRequest) ([]dto.UserShortResponse, int64, error) {
	users, count, err := u.Repository.FindAllPagination(ctx, dto.Limit, dto.Page)
	if err != nil {
		return nil, 0, err
	}
	
	return u.UserMapper.MapToShortList(users), count, nil
}

func (u *UserService) FindById(ctx context.Context, dto *dto.UserFindByIdRequest) (*dto.UserFullResponse, int64, error) {
	users, count, err := u.Repository.FindByIdWithTasks(ctx, dto.Id, dto.Limit, dto.Page)
	if err != nil {
		return nil, 0, exceptions.NewNotFound("User", err.Error(), dto.Id)
	}
	
	return u.UserMapper.MapToFull(users), count, nil
}

func (u *UserService) Create(ctx context.Context, dto *dto.UserCreateRequest) (*dto.UserShortResponse, error) {
	var err error
	user := u.UserMapper.MapFromCreateRequestToEntity(dto)
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	
	err = u.Repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return u.UserMapper.MapToShort(user), nil
}

func (u *UserService) Update(ctx context.Context, dto *dto.UserUpdateRequest) (*dto.UserShortResponse, error) {
	user, err := u.Repository.FindById(ctx, dto.Id)
	if err != nil {
		return nil, err
	}
	u.UserMapper.AssignFromUpdateRequestToEntity(dto, user)
	
	err = u.Repository.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return u.UserMapper.MapToShort(user), nil
}

func (u *UserService) UpdatePassword(ctx context.Context, dto *dto.UserUpdatePasswordRequest) (*dto.UserShortResponse, error) {
	user, err := u.Repository.FindById(ctx, dto.Id)
	if err != nil {
		return nil, exceptions.NewNotFound("User", err.Error(), dto.Id)
	}
	if dto.NewPassword != dto.OldPassword {
		return nil, errors.New("Password incorrect")
	}
	
	err = u.Repository.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return u.UserMapper.MapToShort(user), nil
}

func (u *UserService) Delete(ctx context.Context, dto *dto.UserDeleteByIdRequest) (bool, error) {
	err := u.Repository.DeleteById(ctx, dto.Id)
	
	return err == nil, exceptions.NewNotFound("User", err.Error(), dto.Id)
}