package services

import (
	"context"
	"errors"
	"fmt"
	"leafall/todo-service/internal/dto"
	"leafall/todo-service/internal/mappers"
	"leafall/todo-service/internal/models"
	"leafall/todo-service/internal/repositories"
	"leafall/todo-service/utils"
	"leafall/todo-service/utils/exceptions"

	"github.com/golang-jwt/jwt/v5"
)

type UserServiceInterface interface {
	FindAllPagination(ctx context.Context, dto *dto.UserFindAllRequest) ([]dto.UserShortResponse, int64, error)
	FindById(ctx context.Context, dto *dto.UserFindByIdRequest) (*dto.UserFullResponse, int64, error)
	Create(ctx context.Context, dto *dto.UserCreateRequest) (*dto.UserShortResponse, error)
	Update(ctx context.Context, dto *dto.UserUpdateRequest) (*dto.UserShortResponse, error)
	UpdatePassword(ctx context.Context, dto *dto.UserUpdatePasswordRequest) (*dto.UserShortResponse, error)
	ValidateAccessToken(accessToken string) (jwt.MapClaims, bool, error)
	GenerateAccessTokenFromRefreshToken(dtos *dto.RefreshAccessTokenRequest) (*dto.TokenResponse, error)
	Authorizate(context context.Context, dtoObj *dto.AuthorizationRequest) (*dto.TokenResponse, error)
	Delete(ctx context.Context, dto *dto.UserDeleteByIdRequest) (bool, error)
	Logout(ctx context.Context, dto *dto.LogoutRequest) (error)
}

type UserService struct {
	Repository repositories.UserRepositoryInterface
	UserMapper *mappers.UserMapper
	AccessTokenService TokenServiceInterface
	RefreshTokenService TokenServiceInterface
	TokenRepository repositories.TokenRepositoryInterface
}

func NewUserService(repository repositories.UserRepositoryInterface, 
	userMapper *mappers.UserMapper, 
	accessTokenService TokenServiceInterface, 
	refreshTokenService TokenServiceInterface, 
	tokenRepository repositories.TokenRepositoryInterface) *UserService {
	return &UserService{Repository: repository, 
		UserMapper: userMapper,
		AccessTokenService: accessTokenService,
		RefreshTokenService: refreshTokenService,
		TokenRepository: tokenRepository,}
}

func (u *UserService) FindAllPagination(ctx context.Context, dto *dto.UserFindAllRequest) ([]dto.UserShortResponse, int64, error) {
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return nil, 0, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	users, count, err := u.Repository.FindAllPagination(ctx, dto.Limit, dto.Page)
	if err != nil {
		return nil, 0, err
	}
	
	return u.UserMapper.MapToShortList(users), count, nil
}

func (u *UserService) FindById(ctx context.Context, dto *dto.UserFindByIdRequest) (*dto.UserFullResponse, int64, error) {
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return nil, 0, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	users, count, err := u.Repository.FindByIdWithTasks(ctx, dto.Id, dto.Limit, dto.Page)
	if err != nil {
		return nil, 0, exceptions.NewNotFound("User", err.Error(), dto.Id)
	}
	
	return u.UserMapper.MapToFull(users), count, nil
}

func (u *UserService) Create(ctx context.Context, dto *dto.UserCreateRequest) (*dto.UserShortResponse, error) {
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return nil, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	user := u.UserMapper.MapFromCreateRequestToEntity(dto)
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	
	err = u.Repository.Create(ctx, user)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s: %s", "Error! Database creator: ", err.Error()))
	}
	
	return u.UserMapper.MapToShort(user), nil
}

func (u *UserService) Update(ctx context.Context, dto *dto.UserUpdateRequest) (*dto.UserShortResponse, error) {
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return nil, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
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
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return nil, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	user, err := u.Repository.FindById(ctx, dto.Id)
	if err != nil {
		return nil, exceptions.NewNotFound("User", err.Error(), dto.Id)
	}
	if dto.NewPassword != dto.RepeatPassword {
		return nil, exceptions.NewBadRequestSimple("Password incorrect", err.Error())
	}
	if !utils.EqualPasswordHash(dto.OldPassword, user.Password) {
		return nil, exceptions.NewBadRequestSimple("Password incorrect", err.Error())
	}
	user.Password, err = utils.HashPassword(dto.NewPassword)
	err = u.Repository.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return u.UserMapper.MapToShort(user), nil
}

func (u *UserService) ValidateAccessToken(accessToken string) (jwt.MapClaims, bool, error) {
	token, err := u.AccessTokenService.ParseToken(accessToken)
	if err != nil {
		 return nil, false, err
	}
	claims := token.Claims.(jwt.MapClaims);
	return claims, token.Valid, nil
}

func (u *UserService) Logout(ctx context.Context, dto *dto.LogoutRequest) (error) {
	err := validate.StructCtx(ctx, dto);
	if err != nil {
		return exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	return u.TokenRepository.DeleteToken(ctx, dto.RefreshToken)
}

func (u *UserService) GenerateAccessTokenFromRefreshToken(dtos *dto.RefreshAccessTokenRequest) (*dto.TokenResponse, error) {
	err := validate.Struct(dtos)
	if err != nil {
		return nil, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	refreshToken, err := u.RefreshTokenService.ParseToken(dtos.RefreshToken)
	if err != nil {
		 return nil, err
	}
	refreshClaims := refreshToken.Claims.(jwt.MapClaims);
	userId := refreshClaims["user_id"].(int64)
	accessToken, err := u.AccessTokenService.GenerateToken(userId)
	if err != nil {
		 return nil, err
	}
	response := &dto.TokenResponse{RefreshToken: dtos.RefreshToken, AccessToken: accessToken}
	return response, nil
}

func (u *UserService) Authorizate(context context.Context, dtoObj *dto.AuthorizationRequest) (*dto.TokenResponse, error) {
	err := validate.StructCtx(context, dtoObj)
	if err != nil {
		return nil, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	user, err := u.Repository.FindByLogin(context, dtoObj.Login)
	if err != nil {
		return nil, exceptions.NewBadRequestSimple("Incorrect login or password", "")
	}

	if !utils.EqualPasswordHash(dtoObj.Password, user.Password) {
		return nil, exceptions.NewBadRequestSimple("Incorrect login or password", "")
	}
	refreshToken, err := u.RefreshTokenService.GenerateToken(user.Id)
	if err != nil {
		return nil, err
	}
	accessToken, err := u.AccessTokenService.GenerateToken(user.Id)
	if err != nil { 
		return nil, err
	}
	token := models.Token{ Token: refreshToken, UserId: user.Id}
	err = u.TokenRepository.Create(context, &token, u.RefreshTokenService.GetDuration())
	if err != nil {
		return nil, err
	}
	response := dto.TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	return &response, nil
}

func (u *UserService) Delete(ctx context.Context, dto *dto.UserDeleteByIdRequest) (bool, error) {
	err := validate.StructCtx(ctx, dto)
	if err != nil {
		return false, exceptions.NewBadRequestSimple("Validation has bad format", err.Error())
	}
	err = u.Repository.DeleteById(ctx, dto.Id)
	
	return err == nil, exceptions.NewNotFound("User", err.Error(), dto.Id)
}