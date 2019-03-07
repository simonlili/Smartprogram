package auth

import (
"encoding/json"
"fmt"
"github.com/chengxiaoer/Smartprogram/context"
"github.com/chengxiaoer/Smartprogram/utils"
)

const (
	sessionKeyURL = "https://openapi.baidu.com/nalogin/getSessionKeyByCode"
)

type Auth struct {
	*context.Context
}

type RespGetSessionKey struct {
	utils.CommonError
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

func NewAuth(ctx *context.Context) (auth *Auth) {
	auth = new(Auth)
	auth.Context = ctx
	return
}

func (auth *Auth) GetSessionKey(code string) (resp RespGetSessionKey, err error) {
	params := map[string]string{
		"code":      code,
		"client_id": auth.AppKey,
		"sk":        auth.AppSecret,
	}
	respByte, err := utils.HttpPost(sessionKeyURL, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(respByte, &resp)
	if resp.ErrMsg != "" {
		err = fmt.Errorf("get session key error, code:%v, msg:%s, desc:%s", resp.ErrCode, resp.ErrMsg, resp.ErrDesc)
	}
	return
}