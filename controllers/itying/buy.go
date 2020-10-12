package itying

import (
	"fmt"
	"strconv"
	"xiaomi/models"
)

type BuyController struct {
	BaseController
}

func (c *BuyController) Checkout() {

	c.SuperInit()
	//1、获取要结算的商品
	cartList := []models.Cart{}
	orderList := []models.Cart{} //要结算的商品
	models.Cookie.Get(c.Ctx, "cartList", &cartList)

	var allPrice float64
	//执行计算总价
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}
	//判断去结算页面有没有要结算的商品
	if len(orderList) == 0 {
		c.Redirect("/", 302)
		return
	}

	c.Data["orderList"] = orderList
	c.Data["allPrice"] = allPrice

	//2、获取收货地址
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	addressList := []models.Address{}
	models.DB.Where("uid=?", user.Id).Order("default_address desc").Find(&addressList)
	c.Data["addressList"] = addressList

	//3、防止重复提交订单 生成签名
	orderSign := models.Md5(models.GetRandomNum())
	c.SetSession("orderSign", orderSign)
	c.Data["orderSign"] = orderSign

	c.TplName = "itying/buy/checkout.html"
}

/*
提交订单
   1、获取收货地址信息
   2、获取购买商品的信息
   3、把订单信息放在订单表，把商品信息放在商品表
   4、删除购物车里面的选中数据
*/
func (c *BuyController) DoOrder() {
	//0、防止重复提交订单
	orderSign := c.GetString("orderSign")
	sessionOrderSign := c.GetSession("orderSign")
	if sessionOrderSign != orderSign {
		c.Redirect("/", 302)
		return
	}
	c.DelSession("orderSign")

	// 1、获取收货地址信息
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)

	addressResult := []models.Address{}
	models.DB.Where("uid=? AND default_address=1", user.Id).Find(&addressResult)

	if len(addressResult) > 0 {

		// 2、获取购买商品的信息   orderList就是要购买的商品信息
		cartList := []models.Cart{}
		orderList := []models.Cart{} //要结算的商品
		models.Cookie.Get(c.Ctx, "cartList", &cartList)
		var allPrice float64
		for i := 0; i < len(cartList); i++ {
			if cartList[i].Checked {
				allPrice += cartList[i].Price * float64(cartList[i].Num)
				orderList = append(orderList, cartList[i])
			}
		}

		//  3、把订单信息放在订单表，把商品信息放在商品表
		order := models.Order{
			OrderId:     models.GetOrderId(),
			Uid:         user.Id,
			AllPrice:    allPrice,
			Phone:       addressResult[0].Phone,
			Name:        addressResult[0].Name,
			Address:     addressResult[0].Address,
			Zipcode:     addressResult[0].Zipcode,
			PayStatus:   0,
			PayType:     0,
			OrderStatus: 0,
			AddTime:     int(models.GetUnix()),
		}
		err := models.DB.Create(&order).Error
		if err == nil {
			for i := 0; i < len(orderList); i++ {
				orderItem := models.OrderItem{
					OrderId:      order.Id,
					Uid:          user.Id,
					ProductTitle: orderList[i].Title,
					ProductId:    orderList[i].Id,
					ProductImg:   orderList[i].GoodsImg,
					ProductPrice: orderList[i].Price,
					ProductNum:   orderList[i].Num,
					GoodsVersion: orderList[i].GoodsVersion,
					GoodsColor:   orderList[i].GoodsColor,
					AddTime:      int(models.GetUnix()),
				}
				err := models.DB.Create(&orderItem).Error
				if err != nil {
					fmt.Println(err)
				}
			}
			// 4、删除购物车里面的选中数据

			noSelectedCartList := []models.Cart{}
			for i := 0; i < len(cartList); i++ {
				if !cartList[i].Checked {
					noSelectedCartList = append(noSelectedCartList, cartList[i])
				}
			}
			models.Cookie.Set(c.Ctx, "cartList", noSelectedCartList)
			c.Redirect("/buy/confirm?id="+strconv.Itoa(order.Id), 302)

		} else {
			//非法请求
			c.Redirect("/", 302)
		}
	} else {
		//非法请求
		c.Redirect("/", 302)
	}

}

//去结算
func (c *BuyController) Confirm() {
	c.SuperInit()
	id, err := c.GetInt("id")
	if err != nil {
		c.Redirect("/", 302)
		return
	}
	//获取用户信息
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)

	//获取主订单信息
	order := models.Order{}
	models.DB.Where("id=?", id).Find(&order)
	c.Data["order"] = order
	//判断当前数据是否合法
	if user.Id != order.Uid {
		c.Redirect("/", 302)
		return
	}

	//获取主订单下面的商品信息
	orderItem := []models.OrderItem{}
	models.DB.Where("order_id=?", id).Find(&orderItem)
	c.Data["orderItem"] = orderItem

	c.TplName = "itying/buy/confirm.html"
}

//获取订单状态
func (c *BuyController) OrderPayStatus() {
	//1、获取订单号
	id, err := c.GetInt("id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.ServeJSON()
		return
	}
	//2、查询订单
	order := models.Order{}
	models.DB.Where("id=?", id).Find(&order)

	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	//3、判断当前数据是否合法
	if user.Id != order.Uid {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.ServeJSON()
		return
	}

	//4、判断订单的支付状态
	if order.PayStatus == 1 && order.OrderStatus == 1 {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"message": "已支付",
		}
		c.ServeJSON()

	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "未支付",
		}
		c.ServeJSON()
	}

}
