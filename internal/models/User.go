package models

import "time"

type User struct {
	Id        int64     `json:"id" db:"id"`
	Login     string    `json:"login" db:"login"`
	Password  string    `json:"_" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	Tasks     []Task    `json:"tasks"`
}