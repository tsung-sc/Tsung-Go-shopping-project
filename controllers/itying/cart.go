package itying

import (
	"strconv"
	"xiaomi/models"
)

type CartController struct {
	BaseController
}

func (c *CartController) Get() {
	c.SuperInit()
	cartList := []models.Cart{}
	models.Cookie.Get(c.Ctx, "cartList", &cartList)

	var allPrice float64
	//执行计算总价
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	c.Data["cartList"] = cartList
	c.Data["allPrice"] = allPrice

	c.TplName = "itying/cart/cart.html"

}

func (c *CartController) AddCart() {
	c.SuperInit()
	/*
	   购物车数据保持到哪里？：

	           1、购物车数据保存在本地    （cookie）

	           2、购物车数据保存到服务器(mysql)   （必须登录）

	           3、没有登录 购物车数据保存到本地 ， 登录成功后购物车数据保存到服务器  （用的最多）


	   增加购物车的实现逻辑：

	           1、获取增加购物车的数据  （把哪一个商品加入到购物车）

	           2、判断购物车有没有数据   （cookie）

	           3、如果购物车没有任何数据  直接把当前数据写入cookie

	           4、如果购物车有数据

	               4、1、判断购物车有没有当前数据

	                       有当前数据让当前数据的数量加1，然后写入到cookie

	                       如果没有当前数据直接写入cookie
	*/

	colorId, err1 := c.GetInt("color_id")
	goodsId, err2 := c.GetInt("goods_id")

	goods := models.Goods{}
	goodsColor := models.GoodsColor{}
	err3 := models.DB.Where("id=?", goodsId).Find(&goods).Error
	err4 := models.DB.Where("id=?", colorId).Find(&goodsColor).Error

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {

		c.Ctx.Redirect(302, "/item_"+strconv.Itoa(goods.Id)+".html")
		return
	}
	// 1、获取增加购物车的数据  （把哪一个商品加入到购物车）

	currentData := models.Cart{
		Id:           goodsId,
		Title:        goods.Title,
		Price:        goods.Price,
		GoodsVersion: goods.GoodsVersion,
		Num:          1,
		GoodsColor:   goodsColor.ColorName,
		GoodsImg:     goods.GoodsImg,
		GoodsGift:    goods.GoodsGift, /*赠品*/
		GoodsAttr:    "",              //根据自己的需求拓展
		Checked:      true,            /*默认选中*/
	}

	//  2、判断购物车有没有数据   （cookie）
	cartList := []models.Cart{}
	models.Cookie.Get(c.Ctx, "cartList", &cartList)
	if len(cartList) > 0 { //购物车有数据
		//4、判断购物车有没有当前数据
		if models.CartHasData(cartList, currentData) {
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr {
					cartList[i].Num = cartList[i].Num + 1
				}
			}
		} else {
			cartList = append(cartList, currentData)
		}
		models.Cookie.Set(c.Ctx, "cartList", cartList)

	} else {
		//3、如果购物车没有任何数据  直接把当前数据写入cookie
		cartList = append(cartList, currentData)
		models.Cookie.Set(c.Ctx, "cartList", cartList)
	}

	c.Data["goods"] = goods
	c.TplName = "itying/cart/addcart_success.html"
}

func (c *CartController) DecCart() {
	var flag bool
	var allPrice float64
	var currentAllPrice float64
	var num int

	goodsId, _ := c.GetInt("goods_id")
	goodsColor := c.GetString("goods_color")
	goodsAttr := ""

	cartList := []models.Cart{}
	models.Cookie.Get(c.Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == goodsAttr {
			if cartList[i].Num > 1 {
				cartList[i].Num = cartList[i].Num - 1
			}
			flag = true
			num = cartList[i].Num
			currentAllPrice = cartList[i].Price * float64(cartList[i].Num)
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	if flag {
		models.Cookie.Set(c.Ctx, "cartList", cartList)
		c.Data["json"] = map[string]interface{}{
			"success":         true,
			"message":         "修改数量成功",
			"allPrice":        allPrice,
			"currentAllPrice": currentAllPrice,
			"num":             num,
		}

	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
	}
	c.ServeJSON()

}

func (c *CartController) IncCart() {
	var flag bool
	var allPrice float64
	var currentAllPrice float64
	var num int

	goodsId, _ := c.GetInt("goods_id")
	goodsColor := c.GetString("goods_color")
	goodsAttr := ""

	cartList := []models.Cart{}
	models.Cookie.Get(c.Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == goodsAttr {
			cartList[i].Num = cartList[i].Num + 1
			flag = true
			num = cartList[i].Num
			currentAllPrice = cartList[i].Price * float64(cartList[i].Num)
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	if flag {
		models.Cookie.Set(c.Ctx, "cartList", cartList)
		c.Data["json"] = map[string]interface{}{
			"success":         true,
			"message":         "修改数量成功",
			"allPrice":        allPrice,
			"currentAllPrice": currentAllPrice,
			"num":             num,
		}

	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
	}
	c.ServeJSON()
}

func (c *CartController) ChangeOneCart() {
	var flag bool
	var allPrice float64
	goodsId, _ := c.GetInt("goods_id")
	goodsColor := c.GetString("goods_color")
	goodsAttr := ""

	cartList := []models.Cart{}
	models.Cookie.Get(c.Ctx, "cartList", &cartList)

	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == goodsAttr {
			cartList[i].Checked = !cartList[i].Checked
			flag = true
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	if flag {
		models.Cookie.Set(c.Ctx, "cartList", cartList)
		c.Data["json"] = map[string]interface{}{
			"success":  true,
			"message":  "修改状态成功",
			"allPrice": allPrice,
		}

	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
	}
	c.ServeJSON()
}

//全选反选
func (c *CartController) ChangeAllCart() {
	flag, _ := c.GetInt("flag")
	var allPrice float64
	cartList := []models.Cart{}
	models.Cookie.Get(c.Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if flag == 1 {
			cartList[i].Checked = true
		} else {
			cartList[i].Checked = false
		}
		//计算总价
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	models.Cookie.Set(c.Ctx, "cartList", cartList)

	c.Data["json"] = map[string]interface{}{
		"success":  true,
		"allPrice": allPrice,
	}
	c.ServeJSON()
}

func (c *CartController) DelCart() {

	goodsId, _ := c.GetInt("goods_id")
	goodsColor := c.GetString("goods_color")
	goodsAttr := ""

	cartList := []models.Cart{}
	models.Cookie.Get(c.Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == goodsAttr {
			//执行删除
			cartList = append(cartList[:i], cartList[(i+1):]...)
		}
	}
	models.Cookie.Set(c.Ctx, "cartList", cartList)

	c.Redirect("/cart", 302)
	// c.Ctx.WriteString("xxx")

}
