package admin

import (
	"strconv"
	"strings"
	"xiaomi/models"
)

type GoodsTypeAttrController struct {
	BaseController
}

func (c *GoodsTypeAttrController) Get() {

	cateId, err1 := c.GetInt("cate_id")
	if err1 != nil {
		c.Error("非法请求", "/goodsType")
	}
	//获取当前的类型
	goodsType := models.GoodsType{Id: cateId}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType

	//查询当前类型下面的商品类型属性
	goodsTypeAttr := []models.GoodsTypeAttribute{}
	models.DB.Where("cate_id=?", cateId).Find(&goodsTypeAttr)
	c.Data["goodsTypeAttrList"] = goodsTypeAttr

	c.TplName = "admin/goodsTypeAttribute/index.html"
}

func (c *GoodsTypeAttrController) Add() {

	cateId, err1 := c.GetInt("cate_id")
	if err1 != nil {
		c.Error("非法请求", "/goodsType")
	}

	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsTypeList"] = goodsType
	c.Data["cateId"] = cateId
	c.TplName = "admin/goodsTypeAttribute/add.html"
}

func (c *GoodsTypeAttrController) DoAdd() {

	title := c.GetString("title")
	cateId, err1 := c.GetInt("cate_id")
	attrType, err2 := c.GetInt("attr_type")
	attrValue := c.GetString("attr_value")
	sort, err4 := c.GetInt("sort")
	if err1 != nil || err2 != nil {
		c.Error("非法请求", "/goodsType")
		return
	}
	if strings.Trim(title, " ") == "" {
		c.Error("商品类型属性名称不能为空", "/goodsTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}
	if err4 != nil {
		c.Error("排序值错误", "/goodsTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}
	goodsTypeAttr := models.GoodsTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		AddTime:   int(models.GetUnix()),
		Sort:      sort,
	}
	models.DB.Create(&goodsTypeAttr)
	c.Success("增加成功", "/goodsTypeAttribute?cate_id="+strconv.Itoa(cateId))

}

func (c *GoodsTypeAttrController) Edit() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("非法请求", "/goodType")
		return
	}
	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	models.DB.Find(&goodsTypeAttr)
	c.Data["goodsTypeAttr"] = goodsTypeAttr
	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsTypeList"] = goodsType
	c.TplName = "admin/goodsTypeAttribute/edit.html"
}

func (c *GoodsTypeAttrController) DoEdit() {
	id, err := c.GetInt("id")
	title := c.GetString("title")
	cateId, err1 := c.GetInt("cate_id")
	attrType, err2 := c.GetInt("attr_type")
	attrValue := c.GetString("attr_value")
	sort, err4 := c.GetInt("sort")
	if err != nil || err1 != nil || err2 != nil {
		c.Error("非法请求", "/goodsTypeAttribute")
		return
	}
	if strings.Trim(title, " ") == "" {
		c.Error("商品类型属性名称不能为空", "/goodsTypeAttribute/edit?cate_id="+strconv.Itoa(id))
		return
	}
	if err4 != nil {
		c.Error("排序值错误", "/goodsTypeAttribute/edit?cate_id="+strconv.Itoa(id))
		return
	}
	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	models.DB.Find(&goodsTypeAttr)
	goodsTypeAttr.Title = title
	goodsTypeAttr.CateId = cateId
	goodsTypeAttr.AttrType = attrType
	goodsTypeAttr.AttrValue = attrValue
	goodsTypeAttr.Sort = sort
	err3 := models.DB.Save(&goodsTypeAttr).Error
	if err3 != nil {
		c.Error("修改数据失败", "/goodsTypeAttribute/edit?cate_id="+strconv.Itoa(id))
	}
	c.Success("修改数据成功", "/goodsTypeAttribute?cate_id="+strconv.Itoa(cateId))
}
func (c *GoodsTypeAttrController) Delete() {
	id, err := c.GetInt("id")
	cateId, err1 := c.GetInt("cate_id")
	if err != nil {
		c.Error("传入参数错误", "/goodsTypeAttribute?cate_id="+strconv.Itoa(cateId))
		return
	}
	if err1 != nil {
		c.Error("非法请求", "/goodsType")
	}
	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	models.DB.Delete(&goodsTypeAttr)
	c.Success("删除数据成功", "/goodsTypeAttribute?cate_id="+strconv.Itoa(cateId))
}
