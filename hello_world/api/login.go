package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"go_code/hello_world/model"
	"go_code/hello_world/service"
)

func WXLogin(ctx context.Context, c *app.RequestContext) {
	req := new(model.WxLoginRequest)
	err := c.BindAndValidate(req)
	if err != nil {
		fmt.Println("catch Wx login id error: ", err)
		c.JSON(400, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	//Todo:微信小程序登录后端逻辑
	service.WxLoginService(req.Code)
	//log.Println("code is :", req.Code)
	c.JSON(200, utils.H{"success": true, "message": "Login successful"})
}
