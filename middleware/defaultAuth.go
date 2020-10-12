package middleware

import (
	"xiaomi/models"

	"github.com/astaxie/beego/context"
)

func DefaultAuth(ctx *context.Context) {
	//判断前端用户有没有登陆
	user := models.User{}
	models.Cookie.Get(ctx, "userinfo", &user)
	if len(user.Phone) != 11 {
		ctx.Redirect(302, "/pass/login")
	}
}
