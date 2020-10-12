package itying

import (
	"math"
	"strconv"
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
