package api

import "github.com/astaxie/beego"

type V2Controller struct {
	beego.Controller
}

func (c *V2Controller) Get() {
	c.Ctx.WriteString("api v2")
}
