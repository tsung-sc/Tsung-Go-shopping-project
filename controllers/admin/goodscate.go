package admin

import (
	"os"
	"strconv"
	"xiaomi/models"

	"github.com/astaxie/beego"
)

type GoodsCateController struct {
	BaseController
}

func (c *GoodsCateController) Get() {
	goodsCate := []models.GoodsCate{}
	models.DB.Preload("GoodsCateItem").Where("pid=0").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	c.TplName = "admin/goodsCate/index.html"
}

func (c *GoodsCateController) Add() {
	goodsCate := []models.GoodsCate{}
	models.DB.Where("pid=0").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	c.TplName = "admin/goodsCate/add.html"
}

func (c *GoodsCateController) DoAdd() {
	title := c.GetString("title")
	pid, err1 := c.GetInt("pid")
	link := c.GetString("link")
	template := c.GetString("template")
	subTitle := c.GetString("sub_title")
	keywords := c.GetString("keywords")
	description := c.GetString("description")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")

	if err1 != nil || err3 != nil {
		c.Error("传入参数类型不正确", "/goodsCate/add")
		return
	}
	if err2 != nil {
		c.Error("排序值必须是整数", "/goodsCate/add")
		return
	}
	uploadDir, _ := c.UploadImg("cate_img")
	goodsCate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		SubTitle:    subTitle,
		Link:        link,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     uploadDir,
		Sort:        sort,
		Status:      status,
		AddTime:     int(models.GetUnix()),
	}
	err := models.DB.Create(&goodsCate).Error
	if err != nil {
		c.Error("增加失败", "/goodsCate/add")
		return
	}
	c.Success("增加成功", "/goodsCate")
}

func (c *GoodsCateController) Edit() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/goodsCate")
		return
	}
	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)
	c.Data["goodsCate"] = goodsCate
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid=0").Find(&goodsCateList)
	c.Data["goodsCateList"] = goodsCateList
	c.TplName = "admin/goodsCate/edit.html"
}

func (c *GoodsCateController) DoEdit() {
	id, err := c.GetInt("id")
	title := c.GetString("title")
	pid, err1 := c.GetInt("pid")
	link := c.GetString("link")
	template := c.GetString("template")
	subTitle := c.GetString("sub_title")
	keywords := c.GetString("keywords")
	description := c.GetString("description")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")
	if err != nil || err1 != nil || err3 != nil {
		c.Error("传入参数类型不正确", "/goodsCate/edit")
		return
	}
	if err2 != nil {
		c.Error("排序值必须是整数", "/goodsCate/edit?id="+strconv.Itoa(id))
		return
	}
	uploadDir, _ := c.UploadImg("cate_img")
	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)
	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status
	if uploadDir != "" {
		goodsCate.CateImg = uploadDir
	}
	err5 := models.DB.Save(&goodsCate).Error
	if err5 != nil {
		c.Error("修改失败", "/goodsCate/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改成功", "/goodsCate")
}

func (c *GoodsCateController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/goodCate")
		return
	}
	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)
	address := "D:/gowork/src/xiaomi/" + goodsCate.CateImg
	test := os.Remove(address)
	if test != nil {
		beego.Error(test)
		c.Error("删除物理机上图片错误", "/goodCate")
		return
	}
	if goodsCate.Pid == 0 {
		goodsCate2 := []models.GoodsCate{}
		models.DB.Where("pid=?", goodsCate.Id).Find(&goodsCate2)
		if len(goodsCate2) > 0 {
			c.Error("请删除当前顶级分类下面的商品！", "/goodsCate")
			return
		}
	}
	models.DB.Delete(&goodsCate)
	c.Success("删除成功", "/goodsCate")
}
