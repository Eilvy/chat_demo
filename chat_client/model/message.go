package model

import (
	"gorm.io/gorm"
)

type Message struct {
	*gorm.Model
	content string
	name    string
}
