package admin

import (
	"strings"
	"xiaomi/models"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.TplName = "admin/login/login.html"
}

// func (c *LoginController) SetYzm() {
// 	yzm := models.GetRandomNum()
// 	models.YzmClient.Put("yanzhengma", yzm, time.Second*60)
// 	c.Goto("/login")
// }

func (c *LoginController) DoLogin() {
	var flag = models.Cpt.VerifyReq(c.Ctx.Request)
	if flag {
		username := strings.Trim(c.GetString("username"), "")
		password := models.Md5(strings.Trim(c.GetString("password"), ""))
		// yanzhengma := c.GetString("yanzhengma")
		// returnval, _ := (models.YzmClient.Get("yanzhengma")).([]uint8)
		// if yanzhengma == string(returnval) && yanzhengma != "" {
		manager := []models.Manager{}
		models.DB.Where("username=? AND password=? AND status=1", username, password).Find(&manager)
		if len(manager) == 1 {
			c.SetSession("userinfo", manager[0])
			c.Success("登陆成功", "/")
		} else {
			c.Error("无登陆权限或用户名密码错误", "/login")
			// 	}
			// } else {
			// 	c.Error("手机验证码错误", "/login")
		}
	} else {
		c.Error("验证码错误", "/login")
	}
}

func (c *LoginController) LoginOut() {
	c.DelSession("userinfo")
	c.Success("退出登录成功,将返回登陆页面！", "/login")
}
