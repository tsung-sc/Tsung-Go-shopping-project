package admin

import (
	"fmt"
	"math"
	"strconv"
	"xiaomi/models"
)

type NavController struct {
	BaseController
}

func (c *NavController) Get() {
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 3
	//查询数据
	nav := []models.Nav{}
	models.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&nav)
	//查询nav表里面的数量
	var count int
	models.DB.Table("nav").Count(&count)
	// if len(nav) == 0 {
	// 	prvPage := page - 1
	// 	if prvPage == 0 {
	// 		prvPage = 1
	// 	}
	// 	c.Goto("/nav?page=" + strconv.Itoa(prvPage))
	// }
	c.Data["navList"] = nav
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.TplName = "admin/nav/index.html"
}

func (c *NavController) Add() {
	c.TplName = "admin/nav/add.html"
}

func (c *NavController) DoAdd() {
	title := c.GetString("title")
	link := c.GetString("link")
	position, _ := c.GetInt("position")
	isOpennew, _ := c.GetInt("is_opennew")
	relation := c.GetString("relation")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")

	nav := models.Nav{
		Title:     title,
		Link:      link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}

	err := models.DB.Create(&nav).Error
	if err != nil {
		c.Error("增加数据失败", "/nav/add")
	} else {
		c.Success("增加成功", "/nav")
	}
}

func (c *NavController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/nav")
		return
	}
	nav := models.Nav{Id: id}
	models.DB.Find(&nav)
	c.Data["nav"] = nav
	c.Data["prevPage"] = c.Ctx.Request.Referer()
	c.TplName = "admin/nav/edit.html"
}

func (c *NavController) DoEdit() {

	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/nav")
		return
	}
	title := c.GetString("title")
	link := c.GetString("link")
	position, _ := c.GetInt("position")
	isOpennew, _ := c.GetInt("is_opennew")
	relation := c.GetString("relation")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")
	prevPage := c.GetString("prevPage")
	fmt.Println("-----------------------", relation)
	//修改
	nav := models.Nav{Id: id}
	models.DB.Find(&nav)
	nav.Title = title
	nav.Link = link
	nav.Position = position
	nav.IsOpennew = isOpennew
	nav.Relation = relation
	nav.Sort = sort
	nav.Status = status

	err2 := models.DB.Save(&nav).Error
	if err2 != nil {
		c.Error("修改数据失败", "/nav/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改数据成功", prevPage)
	}

}

func (c *NavController) Delete() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/nav")
		return
	}
	nav := models.Nav{Id: id}
	models.DB.Delete(&nav)

	c.Success("删除数据成功", c.Ctx.Request.Referer())
}
