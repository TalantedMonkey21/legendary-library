package entity

import "time"

type Note struct {
	ID        uint
	UserID	  uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
