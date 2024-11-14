package user

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username   string `json:"username" gorm:"size:20"`
	SessionKey string `json:"session_key" gorm:"unique;not null"`
	Openid     string `json:"openid" gorm:"size:36;unique;not null"`
	Picture    string `json:"picture"`
}
