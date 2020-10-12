package admin

import (
	"os"
	"strconv"
	"xiaomi/models"

	"github.com/astaxie/beego"
)

type FocusController struct {
	BaseController
}

func (c *FocusController) Get() {
	focus := []models.Focus{}
	models.DB.Find(&focus)
	c.Data["focusList"] = focus
	c.TplName = "admin/focus/index.html"
}

func (c *FocusController) Add() {
	c.TplName = "admin/focus/add.html"
}

func (c *FocusController) DoAdd() {
	focusType, err1 := c.GetInt("focus_type")
	title := c.GetString("title")
	link := c.GetString("link")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")
	if err1 != nil || err3 != nil {
		c.Error("非法请求", "/focus")
		return
	}
	if err2 != nil {
		c.Error("排序表单里面输入的数据不合法", "/focus/add")
		return
	}
	focusImgSrc, err4 := c.UploadImg("focus_img")
	if err4 == nil {
		focus := models.Focus{
			Title:     title,
			FocusType: focusType,
			FocusImg:  focusImgSrc,
			Link:      link,
			Sort:      sort,
			Status:    status,
			AddTime:   int(models.GetUnix()),
		}
		models.DB.Create(&focus)
		c.Success("增加轮播图成功", "/focus")
	} else {
		c.Error("增加轮播图失败", "/focus/add")
		return
	}
}

func (c *FocusController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("非法请求", "/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	c.Data["focus"] = focus
	c.TplName = "admin/focus/edit.html"
}

func (c *FocusController) DoEdit() {
	id, err := c.GetInt("id")
	focusType, err1 := c.GetInt("focus_type")
	title := c.GetString("title")
	link := c.GetString("link")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")
	if err != nil || err1 != nil || err3 != nil {
		c.Error("非法请求", "/focus")
		return
	}
	if err2 != nil {
		c.Error("排序表单里面输入的数据不合法", "/focus/edit?id="+strconv.Itoa(id))
		return
	}
	focusImgSrc, _ := c.UploadImg("focus_img")
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImgSrc != "" {
		focus.FocusImg = focusImgSrc
	}
	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		c.Error("修改轮播图失败", "/focus/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改轮播图成功", "/focus")
}

func (c *FocusController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	address := "D:/gowork/src/xiaomi/" + focus.FocusImg
	test := os.Remove(address)
	if test != nil {
		beego.Error(test)
		c.Error("删除物理机上图片错误", "/focus")
		return
	}
	models.DB.Delete(&focus)
	c.Success("删除轮播图成功", "/focus")
}
