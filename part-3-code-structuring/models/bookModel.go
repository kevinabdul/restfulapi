package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id     		int 	`gorm:"primaryKey`
	Title   	string	`json:"title" form:"title"`
	Author 		string	`json:"author" form:"author"`
	Content 	string	`json:"content" form: "content"`
}

type BookAPI struct {
	Id 		int 	`json:"id" form:"id"`
	Title 	string	`json:"title" form:"title"`
	Author 	string	`json:"author" form:"author"`
}