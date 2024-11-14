package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username string
	Nickname string
	Password string
}

type Register struct {
	*gorm.Model
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
