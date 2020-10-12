package admin

import "xiaomi/models"

type SettingController struct {
	BaseController
}

func (c *SettingController) Get() {
	setting := models.Setting{}
	models.DB.First(&setting)
	c.Data["setting"] = setting
	c.TplName = "admin/setting/index.html"
}

func (c *SettingController) DoEdit() {
	setting := models.Setting{}
	models.DB.Find(&setting)
	c.ParseForm(&setting)
	siteLogo, err := c.UploadImg("site_logo")
	if len(siteLogo) > 0 && err == nil {
		setting.SiteLogo = siteLogo
	}
	noPicture, err := c.UploadImg("no_picture")
	if len(noPicture) > 0 && err == nil {
		setting.NoPicture = noPicture
	}
	err = models.DB.Where("id=1").Save(&setting).Error
	if err != nil {
		c.Error("修改数据失败", "/setting")
		return
	}
	c.Success("修改数据成功", "/setting")
}
