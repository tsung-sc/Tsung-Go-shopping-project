package itying

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"xiaomi/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	c.SuperInit()
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	c.Data["user"] = user
	time := time.Now().Hour()
	if time >= 12 && time <= 18 {
		c.Data["Hello"] = "尊敬的用户下午好"
	} else if time >= 6 && time < 12 {
		c.Data["Hello"] = "尊敬的用户上午好！"
	} else {
		c.Data["Hello"] = "深夜了，注意休息哦！"
	}
	order := []models.Order{}
	models.DB.Where("uid=?", user.Id).Find(&order)
	var wait_pay int
	var wait_rec int
	for i := 0; i < len(order); i++ {
		if order[i].PayStatus == 0 {
			wait_pay += 1
		}
		if order[i].OrderStatus >= 2 && order[i].OrderStatus < 4 {
			wait_rec += 1
		}
	}
	c.Data["wait_pay"] = wait_pay
	c.Data["wait_rec"] = wait_rec
	c.TplName = "itying/user/welcome.html"
}

func (c *UserController) OrderList() {
	c.SuperInit()
	//1、获取当前用户
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	//2、获取当前用户下面的订单信息 并分页

	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	pageSize := 2
	//3、获取搜索关键词
	where := "uid=?"
	keywords := c.GetString("keywords")
	if keywords != "" {
		orderitem := []models.OrderItem{}
		models.DB.Where("product_title like ?", "%"+keywords+"%").Find(&orderitem)
		var str string
		for i := 0; i < len(orderitem); i++ {
			if i == 0 {
				str += strconv.Itoa(orderitem[i].OrderId)
			} else {
				str += "," + strconv.Itoa(orderitem[i].OrderId)
			}
		}
		where += " AND id in (" + str + ")"
	}
	//获取筛选条件
	orderStatus, err := c.GetInt("order_status")
	if err == nil {
		where += " AND order_status=" + strconv.Itoa(orderStatus)
		c.Data["orderStatus"] = orderStatus
	} else {
		c.Data["orderStatus"] = "nil"
	}
	//3、总数量
	var count int
	models.DB.Where(where, user.Id).Table("order").Count(&count)
	order := []models.Order{}
	models.DB.Where(where, user.Id).Offset((page - 1) * pageSize).Limit(pageSize).Preload("OrderItem").Order("add_time desc").Find(&order)

	c.Data["order"] = order
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.Data["keywords"] = keywords
	c.TplName = "itying/user/order.html"
}
func (c *UserController) OrderInfo() {
	c.SuperInit()
	id, _ := c.GetInt("id")
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	order := models.Order{}
	models.DB.Where("id=? AND uid=?", id, user.Id).Preload("OrderItem").Find(&order)
	c.Data["order"] = order
	if order.OrderId == "" {
		c.Redirect("/", 302)
	}
	c.TplName = "itying/user/order_info.html"
}

func (c *UserController) BindMail() {
	c.TplName = "itying/pass/bindmail.html"
}

func (c *UserController) Comment() {
	order_id, err := c.GetInt("id")
	if err != nil {
		c.Redirect(c.Ctx.Request.Referer(), 302)
		return
	}
	user := models.User{}
	ok := models.Cookie.Get(c.Ctx, "userinfo", &user)
	if ok == false {
		c.Redirect(c.Ctx.Request.Referer(), 302)
		return
	}
	realuser := models.DB.First(&user).RowsAffected
	if realuser != 1 {
		c.Redirect(c.Ctx.Request.Referer(), 302)
		return
	}
	order := models.Order{}
	realorder := models.DB.Where("id=? AND uid=? AND pay_status=1 AND order_status=4", order_id, user.Id).First(&order).RowsAffected
	if realorder != 1 {
		c.Redirect(c.Ctx.Request.Referer(), 302)
		return
	}
	orderitem := models.OrderItem{}
	models.DB.Where("order_id=?", order_id).Find(&orderitem)
	c.Data["orderitem"] = orderitem
	c.TplName = "itying/user/comment.html"
}

func (c *UserController) GetCollect() {
	c.SuperInit()
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	pageSize := 2
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	collect := []models.GoodsCollect{}
	models.DB.Where("user_id=?", user.Id).Find(&collect)
	var goodsId []int
	for i := 0; i < len(collect); i++ {
		goodsId = append(goodsId, collect[i].GoodId)
	}
	goods := []models.Goods{}
	models.DB.Where("id in (?)", goodsId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&goods)
	var count int
	models.DB.Where("id in (?)", goodsId).Table("goods").Count(&count)
	c.Data["collect"] = goods
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.TplName = "itying/user/order_collect.html"
}

func (c *UserController) DoComment() {
	order_id, err := c.GetInt("order_id")
	text := c.GetString("text")
	str := c.GetString("str")
	realuser := models.User{}
	ok := models.Cookie.Get(c.Ctx, "userinfo", &realuser)
	if ok == false {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "非法会员",
		}
		c.ServeJSON()
		return
	}
	order := models.Order{}
	realorder := models.DB.Where("id=? AND uid=? AND pay_status=1 AND order_status=4 AND is_comment=0", order_id, realuser.Id).First(&order).RowsAffected
	if realorder != 1 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "非法评论",
		}
		c.ServeJSON()
		return
	}
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "订单号错误",
		}
		c.ServeJSON()
		return
	}
	star, err := c.GetInt("index")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "评论出错",
		}
		c.ServeJSON()
		return
	}

	orderItem := models.OrderItem{}
	models.DB.Where("order_id=?", order_id).First(&orderItem)
	comment := models.Comment{}
	var impress []string
	if strings.Contains(str, "a1") {
		impress = append(impress, "价格实惠 ")
	}
	if strings.Contains(str, "a2") {
		impress = append(impress, "交付准时 ")
	}
	if strings.Contains(str, "a3") {
		impress = append(impress, "包装精美 ")
	}
	if strings.Contains(str, "a4") {
		impress = append(impress, "服务态度友善 ")
	}
	if strings.Contains(str, "a5") {
		impress = append(impress, "能力待提高 ")
	}
	if strings.Contains(str, "a6") {
		impress = append(impress, "延迟送达 ")
	}
	for _, i := range impress {
		comment.Str += i
	}
	comment.OrderId = orderItem.OrderId
	comment.UserId = orderItem.Uid
	comment.Star = star
	comment.Text = text
	comment.GoodId = orderItem.ProductId
	comment.AddTime = models.GetUnix()
	err = models.DB.Debug().Create(&comment).Error
	if err != nil {
		fmt.Println(err)
	}
	models.DB.Model(&order).Update("is_comment", 1)
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"msg":     "评论成功",
	}
	c.ServeJSON()
}

func (c *UserController) UserComment() {
	c.SuperInit()
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	pageSize := 2
	user := models.User{}
	ok := models.Cookie.Get(c.Ctx, "userinfo", &user)
	if ok == false {
		c.Redirect(c.Ctx.Request.Referer(), 302)
		return
	}
	realuser := models.DB.First(&user).RowsAffected
	if realuser != 1 {
		c.Redirect(c.Ctx.Request.Referer(), 302)
		return
	}
	Comment := []models.Comment{}
	var count int
	models.DB.Table("comment").Where("user_id=?", user.Id).Count(&count)
	models.DB.Where("user_id=?", user.Id).Order("add_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&Comment)
	for i := 0; i < len(Comment); i++ {
		OrderItem := models.OrderItem{}
		models.DB.Where("order_id=?", Comment[i].OrderId).First(&OrderItem)
		Comment[i].OrderItem = OrderItem
	}
	c.Data["comment"] = Comment
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.TplName = "itying/user/userComment.html"
}
