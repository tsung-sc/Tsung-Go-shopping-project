package routers

import (
	"xiaomi/controllers/admin"
	"xiaomi/middleware"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/"+beego.AppConfig.String("adminPath"),
		beego.NSBefore(middleware.AdminAuth),
		//后台管理
		beego.NSRouter("/", &admin.MainController{}),
		beego.NSRouter("/welcome", &admin.MainController{}, "get:Welcome"),
		beego.NSRouter("/main/changestatus", &admin.MainController{}, "get:ChangeStatus"),
		beego.NSRouter("/main/editnum", &admin.MainController{}, "get:EditNum"),
		beego.NSRouter("/login", &admin.LoginController{}),
		// beego.NSRouter("/login/yanzhengma", &admin.LoginController{}, "get:SetYzm"),
		beego.NSRouter("/login/dologin", &admin.LoginController{}, "post:DoLogin"),
		beego.NSRouter("/login/loginout", &admin.LoginController{}, "get:LoginOut"),
		beego.NSRouter("/focus", &admin.FocusController{}),
		//管理员管理
		beego.NSRouter("/manager", &admin.ManagerController{}),
		beego.NSRouter("/manager/add", &admin.ManagerController{}, "get:Add"),
		beego.NSRouter("/manager/edit", &admin.ManagerController{}, "get:Edit"),
		beego.NSRouter("/manager/doadd", &admin.ManagerController{}, "post:DoAdd"),
		beego.NSRouter("/manager/doedit", &admin.ManagerController{}, "post:DoEdit"),
		beego.NSRouter("/manager/delete", &admin.ManagerController{}, "get:Delete"),
		//部门管理
		beego.NSRouter("/role", &admin.RoleController{}),
		beego.NSRouter("/role/add", &admin.RoleController{}, "get:Add"),
		beego.NSRouter("/role/doadd", &admin.RoleController{}, "post:DoAdd"),
		beego.NSRouter("/role/edit", &admin.RoleController{}, "get:Edit"),
		beego.NSRouter("/role/doedit", &admin.RoleController{}, "post:DoEdit"),
		beego.NSRouter("/role/delete", &admin.RoleController{}, "get:Delete"),
		beego.NSRouter("/role/auth", &admin.RoleController{}, "get:Auth"),
		beego.NSRouter("/role/doauth", &admin.RoleController{}, "post:DoAuth"),
		//权限管理
		beego.NSRouter("/access", &admin.AccessController{}),
		beego.NSRouter("/access/add", &admin.AccessController{}, "get:Add"),
		beego.NSRouter("/access/edit", &admin.AccessController{}, "get:Edit"),
		beego.NSRouter("/access/doadd", &admin.AccessController{}, "post:DoAdd"),
		beego.NSRouter("/access/doedit", &admin.AccessController{}, "post:DoEdit"),
		beego.NSRouter("/access/delete", &admin.AccessController{}, "get:Delete"),
		//轮播图管理
		beego.NSRouter("/focus", &admin.FocusController{}),
		beego.NSRouter("/focus/add", &admin.FocusController{}, "get:Add"),
		beego.NSRouter("/focus/edit", &admin.FocusController{}, "get:Edit"),
		beego.NSRouter("/focus/doadd", &admin.FocusController{}, "post:DoAdd"),
		beego.NSRouter("/focus/doedit", &admin.FocusController{}, "post:DoEdit"),
		beego.NSRouter("/focus/delete", &admin.FocusController{}, "get:Delete"),
		//商品分类管理
		beego.NSRouter("/goodsCate", &admin.GoodsCateController{}),
		beego.NSRouter("/goodsCate/add", &admin.GoodsCateController{}, "get:Add"),
		beego.NSRouter("/goodsCate/edit", &admin.GoodsCateController{}, "get:Edit"),
		beego.NSRouter("/goodsCate/doadd", &admin.GoodsCateController{}, "post:DoAdd"),
		beego.NSRouter("/goodsCate/doedit", &admin.GoodsCateController{}, "post:DoEdit"),
		beego.NSRouter("/goodsCate/delete", &admin.GoodsCateController{}, "get:Delete"),
		//商品类型管理
		beego.NSRouter("/goodsType", &admin.GoodsTypeController{}),
		beego.NSRouter("/goodsType/add", &admin.GoodsTypeController{}, "get:Add"),
		beego.NSRouter("/goodsType/edit", &admin.GoodsTypeController{}, "get:Edit"),
		beego.NSRouter("/goodsType/doadd", &admin.GoodsTypeController{}, "post:DoAdd"),
		beego.NSRouter("/goodsType/doedit", &admin.GoodsTypeController{}, "post:DoEdit"),
		beego.NSRouter("/goodsType/delete", &admin.GoodsTypeController{}, "get:Delete"),
		//商品属性管理
		beego.NSRouter("/goodsTypeAttribute", &admin.GoodsTypeAttrController{}),
		beego.NSRouter("/goodsTypeAttribute/add", &admin.GoodsTypeAttrController{}, "get:Add"),
		beego.NSRouter("/goodsTypeAttribute/edit", &admin.GoodsTypeAttrController{}, "get:Edit"),
		beego.NSRouter("/goodsTypeAttribute/doadd", &admin.GoodsTypeAttrController{}, "post:DoAdd"),
		beego.NSRouter("/goodsTypeAttribute/doedit", &admin.GoodsTypeAttrController{}, "post:DoEdit"),
		beego.NSRouter("/goodsTypeAttribute/delete", &admin.GoodsTypeAttrController{}, "get:Delete"),
		//商品管理
		beego.NSRouter("/goods", &admin.GoodsController{}),
		beego.NSRouter("/goods/add", &admin.GoodsController{}, "get:Add"),
		beego.NSRouter("/goods/edit", &admin.GoodsController{}, "get:Edit"),
		beego.NSRouter("/goods/doadd", &admin.GoodsController{}, "post:DoAdd"),
		beego.NSRouter("/goods/doedit", &admin.GoodsController{}, "post:DoEdit"),
		beego.NSRouter("/goods/delete", &admin.GoodsController{}, "get:Delete"),
		beego.NSRouter("/goods/doUpload", &admin.GoodsController{}, "post:DoUpload"),
		beego.NSRouter("/goods/getGoodsTypeAttribute", &admin.GoodsController{}, "get:GetGoodsTypeAttribute"),
		beego.NSRouter("/goods/changeGoodsImageColor", &admin.GoodsController{}, "get:ChangeGoodsImageColor"),
		beego.NSRouter("/goods/removeGoodsImage", &admin.GoodsController{}, "get:RemoveGoodsImage"),
		//订单管理
		beego.NSRouter("/order", &admin.OrderController{}),
		beego.NSRouter("/order/detail", &admin.OrderController{}, "get:Detail"),
		beego.NSRouter("/order/edit", &admin.OrderController{}, "get:Edit"),
		beego.NSRouter("/order/doEdit", &admin.OrderController{}, "post:DoEdit"),
		beego.NSRouter("/order/delete", &admin.OrderController{}, "get:Delete"),
		//导航管理
		beego.NSRouter("/nav", &admin.NavController{}),
		beego.NSRouter("/nav/add", &admin.NavController{}, "get:Add"),
		beego.NSRouter("/nav/edit", &admin.NavController{}, "get:Edit"),
		beego.NSRouter("/nav/doadd", &admin.NavController{}, "post:DoAdd"),
		beego.NSRouter("/nav/doedit", &admin.NavController{}, "post:DoEdit"),
		beego.NSRouter("/nav/delete", &admin.NavController{}, "get:Delete"),
		//系统设置
		beego.NSRouter("/setting", &admin.SettingController{}),
		beego.NSRouter("/setting/doedit", &admin.SettingController{}, "post:DoEdit"),
	)
	beego.AddNamespace(ns)
}
