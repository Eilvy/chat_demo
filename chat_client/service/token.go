package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go_code/chat_demo/chat_client/model"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Typ      int64  `json:"typ"`
	jwt.StandardClaims
}

var (
	JwtKey = []byte("leiyv000")
	exp    time.Time
)

func CreateToken(user model.User, typ int64) (token string, err error) {

	if typ == model.AccessToken {
		//access Token
		exp = time.Now().Add(time.Hour * 1)
	} else if typ == model.RefreshToken {
		//refresh Token
		exp = time.Now().Add(time.Hour * 24 * 14)
	}
	claim := jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"typ":      typ,
		"exp":      exp.Unix(),
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = tok.SignedString(JwtKey)
	if err != nil {
		fmt.Println("signed token error : ", err.Error())
		return
	}
	return token, err
}
