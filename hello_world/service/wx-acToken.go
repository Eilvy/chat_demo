package service

import (
	"encoding/json"
	"fmt"
	"go_code/hello_world/model"
	"io"
	"net/http"
)

// WxGetAcToken 获取微信accessToken
func WxGetAcToken(code string) (ac model.AccessToken, err error) {
	url := fmt.Sprintf(model.UrlAccessTokenWx, model.WX_APPID, model.WX_APPKEY)

	resp, err := http.Get(url)
	if err != nil {
		return model.AccessToken{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.AccessToken{}, err
	}

	err = json.Unmarshal(data, &ac)
	if err != nil {
		return model.AccessToken{}, err
	}

	return ac, nil
}
