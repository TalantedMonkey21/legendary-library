package models

import (
	
)

// type Note struct {
// 	gorm.Model
// 	Id int `json:"id"`
// 	Title string `json:"title"`
// 	Body string `json:"body"`
// 	Created_at time.Time `json:"created_at"`
// }

type Note struct {
	Id int `json:"id" gorm:"primary_key"`
	Content string `json:"content" gorm:"not null"`
}