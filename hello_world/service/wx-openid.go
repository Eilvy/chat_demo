package service

import (
	"encoding/json"
	"fmt"
	"go_code/hello_world/model"
	"io"
	"net/http"
)

// SessionKey 请求session_key openid接口的返回格式
type SessionKey struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	model.WxError
}

func WxGetOpenid(code string) (sc SessionKey, err error) {
	url := fmt.Sprintf(model.UrlAccessOpenIdWx, model.WX_APPID, model.WX_APPKEY, code)

	//发送请求向微信服务器
	resp, err := http.Get(url)

	if err != nil {
		return SessionKey{}, err
	}

	//对返回体进行读取
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return SessionKey{}, err
	}

	//对返回体进行json解码
	err = json.Unmarshal(data, &sc)
	if err != nil {
		return SessionKey{}, err
	}

	return sc, nil
}
