package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/chat_demo/chat_client/dao"
	"go_code/chat_demo/chat_client/model"
	"go_code/chat_demo/chat_client/resps"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

func Register(c *gin.Context) {
	u := model.Register{}
	err := c.ShouldBind(&u)
	if err != nil {
		log.Println("shouldBind error : ", err)
		return
	}
	err = dao.FindUserByName(u.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("DB error : ", err)
		return
	}

	//加密用户密码
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("password hash error : ", err.Error())
		return
	}
	err = dao.CreateUser(model.User{
		Username: u.Username,
		Password: string(password),
	})

	resps.OK(c)
}

func Login(c *gin.Context) {
	u := model.User{}
	err := c.ShouldBind(&u)
	if err != nil {
		log.Println("shouldBind error : ", err)
		return
	}

	accessToken, err := CreateToken(u, model.AccessToken)
	if err != nil {
		log.Println("create token error : ", err)
		return
	}
	refreshToken, err := CreateToken(u, model.RefreshToken)
	if err != nil {
		log.Println("create token error : ", err)
		return
	}

	resps.OKWithData(c, map[string]string{
		"accessToken":   accessToken,
		"refreshToken":  refreshToken,
		"welcome back ": u.Username,
	})
}
