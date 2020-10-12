package admin

import (
	"strconv"
	"xiaomi/models"
)

type AccessController struct {
	BaseController
}

func (c *AccessController) Get() {
	access := []models.Access{}
	models.DB.Preload("AccessItem").Where("module_id=0").Find(&access)
	c.Data["accessList"] = access
	c.TplName = "admin/access/index.html"
}

func (c *AccessController) Add() {
	access := []models.Access{}
	models.DB.Where("module_id=0").Find(&access)
	c.Data["accessList"] = access
	c.TplName = "admin/access/add.html"
}

func (c *AccessController) DoAdd() {
	moduleName := c.GetString("module_name")
	iType, err1 := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, err2 := c.GetInt("module_id")
	sort, err3 := c.GetInt("sort")
	description := c.GetString("description")
	status, err4 := c.GetInt("status")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		c.Error("传入参数错误", "/access/add")
		return
	}
	access := models.Access{
		ModuleName:  moduleName,
		Type:        iType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	err := models.DB.Create(&access).Error
	if err != nil {
		c.Error("增加数据失败", "/access/add")
		return
	}
	c.Success("增加数据成功", "/access")
}

func (c *AccessController) Edit() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	c.Data["access"] = access
	accessList := []models.Access{}
	models.DB.Where("module_id=0").Find(&accessList)
	c.Data["accessList"] = accessList
	c.TplName = "admin/access/edit.html"
}

func (c *AccessController) DoEdit() {
	moduleName := c.GetString("module_name")
	iType, err1 := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, err2 := c.GetInt("module_id")
	sort, err3 := c.GetInt("sort")
	description := c.GetString("description")
	status, err4 := c.GetInt("status")
	id, err5 := c.GetInt("id")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		c.Error("传入参数错误", "/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = iType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status
	err6 := models.DB.Save(&access).Error
	if err6 != nil {
		c.Error("修改权限失败", "/access/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改权限成功", "/access")
}

func (c *AccessController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	if access.ModuleId == 0 {
		access2 := []models.Access{}
		models.DB.Where("module_id=?", access.Id).Find(&access2)
		if len(access2) > 0 {
			c.Error("请删除当前顶级模块下面的菜单或操作！", "/access")
			return
		}
	}
	models.DB.Delete(&access)
	c.Success("删除成功", "/access")
}
