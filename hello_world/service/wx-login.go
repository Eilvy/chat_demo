package service

import (
	"fmt"
	"go_code/hello_world/model"
)

func WxLoginService(code string) {
	AcTokenJson, err := WxGetOpenid(code)
	if err != nil {
		model.Logger.Error("获取AcTokenJson 错误：", err.Error())
		return
	}
	fmt.Println(AcTokenJson.Openid)
}
