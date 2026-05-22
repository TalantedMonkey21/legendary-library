package models

import "time"

type Note struct {
	ID        uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
