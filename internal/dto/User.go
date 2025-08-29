package dto

import "time"

// UserShortResponse DTO для получения только данных о пользователе
type UserShortResponse struct {
	// Id идентификатор пользователя
	Id        int64          `json:"id" example:"1"`
	// Login логин пользователя
	Login     string         `json:"login" example:"my_login"`
	// CreatedAt дата создания пользователя
	CreatedAt time.Time      `json:"createdAt" example:"2010-10-10"`
}

// UserFullResponse DTO для получения данных о пользователе и его задачах
type UserFullResponse struct {
	// Id идентификатор пользователя
	Id        int64          `json:"id" example:"1"`
	// Login логин пользователя
	Login     string         `json:"login" example:"my_login"`
	// CreatedAt дата создания пользователя
	CreatedAt time.Time      `json:"createdAt" example:"2010-10-10"`
	// Tasks все задачи пользователя
	Tasks     []TaskResponse `json:"tasks"`
}
// UserCreateRequest DTO для регистрации пользователя
type UserCreateRequest struct {
	// Login логин пользователя (обязательное поле)
	Login    string `json:"login" example:"my_login" validate:"required,min=3,max=30"`
	// Password пароль пользователя (обязательное поле)
	Password string `json:"password" example:"12345678" validate:"required,min=8,max=30"`
}

type UserUpdateRequest struct {
	Id       int64  `json:"id" validate:"required,min=0"`
	Login    string `json:"login" validate:"required,min=3,max=30"`
}

type UserUpdatePasswordRequest struct {
	Id       int64  `json:"id" validate:"required,min=0"`
	OldPassword    string `json:"oldPassword" validate:"required,min=8,max=30"`
	NewPassword    string `json:"newPassword" validate:"required,min=8,max=30"`
	RepeatPassword    string `json:"repeatPassword" validate:"required,min=8,max=30"`
}

type UserFindByIdRequest struct {
	Id int64 `query:"id"`
	Limit int `query:"limit" validate:"required,min=0,max=100"`
	Page  int `query:"page" validate:"required,min=1"`
}

type UserFindAllRequest struct {
	Limit int `query:"limit" validate:"required,min=0,max=100"`
	Page  int `query:"page" validate:"required,min=1"`
}

type UserDeleteByIdRequest struct {
	Id int64 `query:"id" validate:"required,min=0"`
}

type AuthorizationRequest struct {
	Login string `json:"login" validate:"required,min=3,max=30"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}