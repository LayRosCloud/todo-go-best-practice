package dto

import "time"

type UserShortResponse struct {
	Id        int64          `json:"id"`
	Login     string         `json:"login"`
	CreatedAt time.Time      `json:"createdAt"`
	Tasks     []TaskResponse `json:"tasks"`
}

type UserFullResponse struct {
	Id        int64          `json:"id"`
	Login     string         `json:"login"`
	CreatedAt time.Time      `json:"createdAt"`
	Tasks     []TaskResponse `json:"tasks"`
}

type UserCreateRequest struct {
	Login    string `json:"login" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdateRequest struct {
	Id       int64  `json:"id" validate:"required,min=0"`
	Login    string `json:"login" validate:"required,min=3"`
}

type UserUpdatePasswordRequest struct {
	Id       int64  `json:"id" validate:"required,min=0"`
	OldPassword    string `json:"oldPassword" validate:"required,min=8"`
	NewPassword    string `json:"newPassword" validate:"required,min=8"`
	RepeatPassword    string `json:"repeatPassword" validate:"required,min=8"`
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
	Id int64 `query:"id"`
}