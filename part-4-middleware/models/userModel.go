package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id     		int 	`gorm:"primaryKey`
	Name   		string	`json:"name" form:"name"`
	Email 		string	`gorm:"unique" json:"email" form:"email"`
	Password 	string	`json:"password" form: "password"`
}

type UserAPI struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
}