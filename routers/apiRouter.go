package routers

import (
	"xiaomi/controllers/api"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSRouter("/", &api.V1Controller{}),
		beego.NSRouter("/nav", &api.V1Controller{}, "get:Nav"),
	)
	beego.AddNamespace(ns)
}
