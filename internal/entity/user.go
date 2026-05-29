package entity

import "time"

type User struct {
	ID uint
	Name string
	PasswordHash string
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
}