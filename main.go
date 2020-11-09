package main

import (
	"encoding/gob"
	"xiaomi/models"
	_ "xiaomi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/astaxie/beego/session/redis"
)

func init() {
	gob.Register(models.Manager{})
}

func main() {
	beego.AddFuncMap("unixToDate", models.UnixToDate)
	beego.AddFuncMap("unix64ToDate", models.Unix64ToDate)
	models.DB.LogMode(true)
	beego.AddFuncMap("setting", models.GetSettingFromColumn)
	beego.AddFuncMap("formatImg", models.FormatImg)
	beego.AddFuncMap("mul", models.Mul)
	beego.AddFuncMap("formatAttr", models.FormatAttr)

	//后台配置允许跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"tsung.top"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true, //是否允许cookie
	}))

	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	beego.Run()
	defer models.DB.Close()
}
