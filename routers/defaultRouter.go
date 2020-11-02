package routers

import (
	"xiaomi/controllers/itying"
	"xiaomi/middleware"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &itying.IndexController{})
	beego.Router("/category_:id([0-9]+).html", &itying.ProductController{}, "get:CategoryList")
	beego.Router("/item_:id([0-9]+).html", &itying.ProductController{}, "get:ProductItem")
	beego.Router("/product/getImgList", &itying.ProductController{}, "get:GetImgList")
	beego.Router("/product/collect", &itying.ProductController{}, "get:Collect")
	beego.Router("/cart", &itying.CartController{})
	beego.Router("/cart/addCart", &itying.CartController{}, "get:AddCart")
	beego.Router("/cart/incCart", &itying.CartController{}, "get:IncCart")
	beego.Router("/cart/decCart", &itying.CartController{}, "get:DecCart")
	beego.Router("/cart/delCart", &itying.CartController{}, "get:DelCart")

	beego.Router("/cart/changeOneCart", &itying.CartController{}, "get:ChangeOneCart")
	beego.Router("/cart/changeAllCart", &itying.CartController{}, "get:ChangeAllCart")

	beego.Router("/pass/sendCode", &itying.PassController{}, "get:SendCode")
	beego.InsertFilter("/pass/MailsendCode", beego.BeforeRouter, middleware.DefaultAuth)
	beego.Router("/pass/MailsendCode", &itying.PassController{}, "get:MailSendCode")

	beego.Router("/pass/doRegister", &itying.PassController{}, "post:DoRegister")

	beego.InsertFilter("/pass/doBindMail", beego.BeforeRouter, middleware.DefaultAuth)
	beego.Router("/pass/doBindMail", &itying.PassController{}, "get:DoBindMail")

	beego.Router("/pass/validateSmsCode", &itying.PassController{}, "get:ValidateSmsCode")
	beego.Router("/pass/login", &itying.PassController{}, "get:Login")
	beego.Router("/pass/registerStep1", &itying.PassController{}, "get:RegisterStep1")
	beego.Router("/pass/registerStep2", &itying.PassController{}, "get:RegisterStep2")
	beego.Router("/pass/registerStep3", &itying.PassController{}, "get:RegisterStep3")
	beego.Router("/pass/login", &itying.PassController{}, "get:Login")
	beego.Router("/pass/doLogin", &itying.PassController{}, "post:DoLogin")
	beego.Router("/pass/loginOut", &itying.PassController{}, "get:LoginOut")

	//配置中间件判断权限
	beego.InsertFilter("/buy/*", beego.BeforeRouter, middleware.DefaultAuth)
	beego.Router("/buy/checkout", &itying.BuyController{}, "get:Checkout")
	beego.Router("/buy/doOrder", &itying.BuyController{}, "post:DoOrder")
	beego.Router("/buy/confirm", &itying.BuyController{}, "get:Confirm")
	beego.Router("/buy/orderPayStatus", &itying.BuyController{}, "get:OrderPayStatus")
	//配置中间件判断权限
	beego.InsertFilter("/address/*", beego.BeforeRouter, middleware.DefaultAuth)
	beego.Router("/address/addAddress", &itying.AddressController{}, "post:AddAddress")
	beego.Router("/address/getOneAddressList", &itying.AddressController{}, "get:GetOneAddressList")
	beego.Router("/address/doEditAddressList", &itying.AddressController{}, "post:DoEditAddressList")
	beego.Router("/address/changeDefaultAddress", &itying.AddressController{}, "get:ChangeDefaultAddress")
	//支付宝支付
	beego.Router("/alipay", &itying.PayController{}, "get:Alipay")
	beego.Router("/alipayNotify", &itying.PayController{}, "post:AlipayNotify")
	beego.Router("/alipayReturn", &itying.PayController{}, "get:AlipayReturn")

	//微信支付
	beego.Router("/wxpay", &itying.PayController{}, "get:WxPay")
	beego.Router("/wxpay/notify", &itying.PayController{}, "post:WxPayNotify")

	//配置中间件判断权限
	beego.InsertFilter("/user/*", beego.BeforeRouter, middleware.DefaultAuth)
	beego.Router("/user", &itying.UserController{})
	beego.Router("/user/order", &itying.UserController{}, "get:OrderList")
	beego.Router("/user/orderinfo", &itying.UserController{}, "get:OrderInfo")
	beego.Router("/user/bindmail", &itying.UserController{}, "get:BindMail")
	beego.Router("/user/collect", &itying.UserController{}, "get:GetCollect")

	//搜索
	// beego.Router("/search", &itying.SearchController{})
	// beego.Router("/search/getOne", &itying.SearchController{}, "get:GetOne")
	// beego.Router("/search/update", &itying.SearchController{}, "get:Update")
	// beego.Router("/search/delete", &itying.SearchController{}, "get:Delete")
	// beego.Router("/search/query", &itying.SearchController{}, "get:Query")
	// beego.Router("/search/filterQuery", &itying.SearchController{}, "get:FilterQuery")
	beego.Router("/search/goodsList", &itying.SearchController{}, "get:GoodsList")
}
