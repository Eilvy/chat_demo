package dao

import (
	"context"
	"errors"
	"fmt"
	"go_code/chat_demo/chat_client/model"
	"gorm.io/gorm"
	"time"
)

func CreateUser(u model.User) (err error) {
	if err = DB.Create(&u).Error; err != nil {
		fmt.Println("crate user error :", err)
	}
	if err = RDB.Set(context.Background(), u.Username, u.Password, time.Hour*24).Err(); err != nil {
		fmt.Println("user into Redis error : ", err)
		return err
	}
	return nil
}

func FindUserByName(username string) (err error) {
	user := model.User{}
	err = DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("user not found : ", err)
			return err
		}
		fmt.Println("find user by username error :", err)
	}
	return err
}
