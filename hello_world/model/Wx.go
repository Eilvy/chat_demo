package model

import (
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"time"
)

var Logger logger.Logger

const (
	WX_APPID          = "wx34dc31d821104107"
	WX_APPKEY         = "d0db71a9944f10333a5a25d40988039d"
	UrlAccessOpenIdWx = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	UrlAccessTokenWx  = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

type WxLoginRequest struct {
	Code string `json:"code" binding:"required"`
}

// WxError 微信服务器返回错误结构体
type WxError struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (err WxError) Error() string {
	return fmt.Sprintf("{errcod: %d, errmsg: %s}", err.Errcode, err.Errmsg)
}

// AccessToken 微信返回的accessToken结构体、
// 文档地址：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html#access-token-%E7%9A%84%E5%AD%98%E5%82%A8%E4%B8%8E%E6%9B%B4%E6%96%B0
type AccessToken struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	CreateTime  time.Time `json:"create_time"` //用于后续判断过期自动续期
}
