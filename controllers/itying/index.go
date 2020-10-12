package itying

import (
	"fmt"
	"time"
	"xiaomi/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {

	//调用功能功能
	c.SuperInit()

	//开始时间
	startTime := time.Now().UnixNano()

	//获取轮播图 注意获取的时候要写地址
	focus := []models.Focus{}
	if hasFocus := models.CacheDb.Get("focus", &focus); hasFocus == true {
		c.Data["focusList"] = focus
	} else {
		models.DB.Where("status=1 AND focus_type=1").Order("sort desc").Find(&focus)
		c.Data["focusList"] = focus
		models.CacheDb.Set("focus", focus)
	}
	//获取楼层数据
	//手机
	redisPhone := []models.Goods{}
	if hasPhone := models.CacheDb.Get("phone", &redisPhone); hasPhone == true {
		c.Data["phoneList"] = redisPhone
	} else {
		phone := models.GetGoodsByCategory(1, "hot", 8)
		c.Data["phoneList"] = phone
		models.CacheDb.Set("phone", phone)
	}
	//电视
	redisTv := []models.Goods{}
	if hasTv := models.CacheDb.Get("tv", &redisTv); hasTv == true {
		c.Data["tvList"] = redisTv
	} else {
		tv := models.GetGoodsByCategory(4, "best", 8)
		c.Data["tvList"] = tv
		models.CacheDb.Set("tv", tv)
	}

	//结束时间
	endTime := time.Now().UnixNano()

	fmt.Println("执行时间", endTime-startTime)

	c.TplName = "itying/index/index.html"
}
