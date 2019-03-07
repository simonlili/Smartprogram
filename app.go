package Smartprogram

import (
	"github.com/chengxiaoer/Smartprogram/auth"
	"github.com/chengxiaoer/Smartprogram/context"
)

type Config struct {
	AppID     string
	AppKey    string
	AppSecret string
}

type App struct {
	ctx  *context.Context
	auth *auth.Auth
}

func NewApp(appid string,app_key string,app_secret string) (app *App) {
	ctx := new(context.Context)
	ctx.AppID = appid
	ctx.AppKey = app_key
	ctx.AppSecret = app_secret
	return &App{ctx: ctx}
}

func (app *App) GetAuth() *auth.Auth {
	if app.auth == nil {
		app.auth = auth.NewAuth(app.ctx)
	}
	return app.auth
}
