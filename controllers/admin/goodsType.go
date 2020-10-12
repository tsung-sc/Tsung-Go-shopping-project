package admin

import (
	"strconv"
	"strings"
	"xiaomi/models"
)

type GoodsTypeController struct {
	BaseController
}

func (c *GoodsTypeController) Get() {
	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsTypeList"] = goodsType
	c.TplName = "admin/goodsType/index.html"
}

func (c *GoodsTypeController) Add() {
	c.TplName = "admin/goodsType/add.html"
}

func (c *GoodsTypeController) DoAdd() {
	title := strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")
	status, err := c.GetInt("status")
	if title == "" {
		c.Error("标题不能为空", "/goodsType/add")
		return
	}
	if err != nil {
		c.Error("传入参数不正确", "/goodsType/add")
		return
	}
	goodsTypeList := []models.GoodsType{}
	models.DB.Where("title=?", title).Find(&goodsTypeList)
	if len(goodsTypeList) != 0 {
		c.Error("该商品已存在！", "/goodsType/add")
		return
	}
	goodsType := models.GoodsType{}
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status
	goodsType.AddTime = int(models.GetUnix())
	err1 := models.DB.Create(&goodsType).Error
	if err1 != nil {
		c.Error("增加商品类型失败", "/goodsType/add")
	} else {
		c.Success("增加商品类型成功", "/goodsType")
	}
}

func (c *GoodsTypeController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/goodsType")
		return
	}
	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType
	c.TplName = "admin/goodsType/edit.html"
}

func (c *GoodsTypeController) DoEdit() {
	title := strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")
	status, err := c.GetInt("status")
	id, err1 := c.GetInt("id")
	if err != nil || err1 != nil {
		c.Error("传入参数错误", "/goodsType")
		return
	}
	if title == "" {
		c.Error("标题不能为空", "/goodsType/edit?id="+strconv.Itoa(id))
		return
	}
	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status
	err2 := models.DB.Save(&goodsType).Error
	if err2 != nil {
		c.Error("修改商品类型失败", "/goodsType/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改商品类型成功", "/goodsType")
	}
}

func (c *GoodsTypeController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/goodsType")
		return
	}
	goodsType := models.GoodsType{Id: id}
	models.DB.Delete(&goodsType)
	c.Success("删除数据成功", "/goodsType")
}
