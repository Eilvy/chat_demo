package model

import (
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
)

type WxLoginRequest struct {
	Code string `json:"code" binding:"required"`
}

var Logger logger.Logger

const (
	WX_APPID          = "wx34dc31d821104107"
	WX_APPKEY         = "d0db71a9944f10333a5a25d40988039d"
	UrlAccessOpenIdWx = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type WxError struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (err WxError) Error() string {
	return fmt.Sprintf("{errcod: %d, errmsg: %s}", err.Errcode, err.Errmsg)
}
