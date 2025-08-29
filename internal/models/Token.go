package models

import "time"

type Token struct {
	Id           int64     `json:"id" db:"id"`
	Token        string    `json:"token" db:"token"`
	UserId       int64     `json:"userId" db:"user_id"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	ExpiredAt time.Time `json:"experationAt" db:"expired_at"`
}