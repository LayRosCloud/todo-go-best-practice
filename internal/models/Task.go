package models

import "time"

type Task struct {
	Id          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
	UserId      int64     `json:"userId" db:"user_id"`
	FinishedAt  time.Time `json:"finishedAt" db:"finished_at"`
}
