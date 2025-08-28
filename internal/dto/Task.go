package dto

import "time"

type TaskResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	UserId      int64  `json:"userId"`
}

type TaskCreateRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required,min=10"`
	UserId      int64  `json:"userId" validate:"required,min=0"`
}

type TaskUpdateRequest struct {
	Id          int64  `json:"id" validate:"min=0"`
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required,min=10"`
}

type TaskFindAllRequest struct {
	Limit int `query:"limit" validate:"required,min=0,max=100"`
	Page  int `query:"page" validate:"required,min=1"`
}

type TaskByIdRequest struct {
	Id int64 `json:"id" validate:"min=0"`
}
