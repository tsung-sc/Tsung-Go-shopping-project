package middleware

import (
	"net/url"
	"strings"
	"xiaomi/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func AdminAuth(ctx *context.Context) {
	pathname := ctx.Request.URL.String()
	userinfo, ok := ctx.Input.Session("userinfo").(models.Manager)
	if !(ok && userinfo.Username != "") {
		if pathname != "/"+beego.AppConfig.String("adminPath")+"/login" && pathname != "/"+beego.AppConfig.String("adminPath")+"/login/dologin" && pathname != "/"+beego.AppConfig.String("adminPath")+"/login/yanzhengma" {
			ctx.Redirect(302, "/"+beego.AppConfig.String("adminPath")+"/login")
		}
	} else {
		pathname = strings.Replace(pathname, "/"+beego.AppConfig.String("adminPath"), "", 1)
		urlPath, _ := url.Parse(pathname)
		if userinfo.IsSuper == 0 && !excludeAuthPath(string(urlPath.Path)) {
			roleId := userinfo.RoleId
			roleAccess := []models.RoleAccess{}
			models.DB.Where("role_id=?", roleId).Find(&roleAccess)
			roleAccessMap := make(map[int]int)
			for _, v := range roleAccess {
				roleAccessMap[v.AccessId] = v.AccessId
			}
			access := models.Access{}
			models.DB.Where("url=?", urlPath.Path).Find(&access)
			if _, ok := roleAccessMap[access.Id]; !ok {
				ctx.WriteString("没有权限")
				return
			}
		}
	}
}
func excludeAuthPath(urlPath string) bool {
	excludeAuthPathSlice := strings.Split(beego.AppConfig.String("excludeAuthPath"), ",")
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
