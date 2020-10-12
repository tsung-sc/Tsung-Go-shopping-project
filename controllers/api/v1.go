package api

import (
	"xiaomi/models"

	"github.com/astaxie/beego"
)

type V1Controller struct {
	beego.Controller
}

func (c *V1Controller) Get() {
	c.Ctx.WriteString("api v1")
}

func (c *V1Controller) Nav() {
	nav := []models.Nav{}
	models.DB.Find(&nav)
	c.Data["json"] = nav
	c.ServeJSON()
}
