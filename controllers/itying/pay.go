package itying

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"xiaomi/models"

	"github.com/astaxie/beego"
	"github.com/objcoding/wxpay"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/smartwalle/alipay/v3"
)

type PayController struct {
	BaseController
}

func (c *PayController) Alipay() {
	AliId, err1 := c.GetInt("id")
	if err1 != nil {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	orderitem := []models.OrderItem{}
	models.DB.Where("order_id=?", AliId).Find(&orderitem)
	var privateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCPUi2yj4RJcqAKfjV/AxjqIjjUq007pGhLClbrwEoCVhgnQU9bAVFnPlDaVdO2xBfu8D/gHCwV9czZKJh51yg1kwE6Yv/hSd966PnGS9NszfxxWfTbWeC6DIHZj66nTECK+vYWX36tKxIG+juzXyfoAuyL1h58oFgF/Zwa1FzpKDTcMUIE+npf3DpMS0Uatf6TsREzVQQs3i0WVxYY3lv1Dmualr6Q3GAFE1j/xt/STE993uh8MVeLS+RcTrrjilPSVH5Z0DLAjkDSH7XUK1lIHASpgOrddEJ8MeT8L7bUJ8nDs5qQ8zfbPVJaXfsF2NYS162HkIW2bl0r5c5Mudt7AgMBAAECggEAWUHVweHNgiyH7WECkhJsvswHVrNEi0NtzGYpEfOUY/YYXsI22Lduaf0OP5u6GZXwTdeEAF+rORX2uLumkiLkINFnr2QedcEbFCHqBIwOpTF36WQbsUw9P8EwUT1BiWFcxPFctzxL2S78sCnBaol1gfHoPYJhRD5b84cpZDAjmPSJk1XAtAtKChUIskLBAsCvwlGHbx/6UQwM9eKgwo0Y67MCPW9wBjE9bRFWBfaeLszEVu3nKyOKLwcUGDXrbmBS8bj9YtyqTG7RjKZmIGuJtPKehEjlNn2ALYYMHUA5VSdVP4LrBiLVLQE5tTDedAzr3uHuWOgLaMeDnfwQg8Hj6QKBgQDGaCoZ3JRg4PbGetduHpGfWOstwZKNwqkDeYvpA03GcRYhzhqG98qdTGEYaLGkszQrh0ZcYk7fs8jsR5Lu2WTzImtiAYDCDEjSDHj5N5PfWR0r2pxrUvX8shun4QKpz4QQ0RUjeujZ2hJHkeCviF2+k5IBDvz6YtBo+H3IgR9QjwKBgQC47IVtuvlYZ0/2TAcAt40YhEgLZOr3NuR4eSnx76zf/8vRmHDEfIqvUdIzzJy3RVYT3uXiB7DYwaHr76ouP3lMOhgLgUlK4Dt8L3UbMP6Asr+6D/uggVmlIHKK76HZIdBL5nGQvTOvwE+fmEivlR5QV+cczqQdCYyNZXETmgAkVQKBgF32JMIcqZR71cLHmFDJX1Okq7P+sWY7YwmHPZA7hVDOa5nU3tE+dpEqA+2oX0DNsY5PwS2tTQc6QJRNjTNadymCCnLenVjIso/vYjc8b+ZdcKg9HsjhACgNPXWy5S0AXt4L9sPXyICreu60EkFvBl5jysh/jaUSuPqNfBxBsk/XAoGBAKYxkzzp4/wKZXfSHh0L2VemUuVCnlTtVWncYtEXeQObbX8CBJ7h2vXzj/mTs2iWfOTA11NLXCmB5FcZfpWv4ACc2U1FtSwA2BUkxZdZcfESNHMwuBEpDvrzbV3mPUvaMsxz366YC+Kw8B5biz+ZwbOtPHzMTfv2wAW3nGdkaSo9AoGBALW8kQ6QwXu3gB7q8Es1zCvSg07tnJks57nJDTtOOg6oc52RvKjndWHT6aFwp/YrESeHNvoudRhPm5nKUTbKDiLixDXcrvEU3b70dJGwWHkBdYU/s5J2o5oaclwiyiUptnIp2eA+Cu+wH1eREGmrOAdeBSFOfdjcfJTFmOCoalJR" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New("2021001186696588", privateKey, true)
	client.LoadAppPublicCertFromFile("crt/appCertPublicKey_2021001186696588.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("crt/alipayRootCert.crt")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("crt/alipayCertPublicKey_RSA2.crt")       // 加载支付宝公钥证书

	//普通公钥模式下配置
	// var privateKey = "MIIEowIBAAKCAQEAmZlUXdYU5qhpSe7bmv+NAEQehXq2wkV2pAhr9zU2fKJSB77V1Z9079SYTs0yGX3XpGKqVYVGZu6F3js9QX1/PPfELHrxOa4mHXQz+TryniXx1e1ng4kDedy1SyS8xTBeGCvReIqAtMUeAo/hedgvdtgYEvY5hZqbV7b3zEpEkq/XgvgaH5nmYR/fDnocwQt5OpDlXNGMrfhU4g0nvQHzb8Kcd2YSjkDB37o/yV9i0Wq8vJqPdWCJOGxCOafy1Qt1VfPF6s9h/9V8mWTGrMGyaUUgz/0ILdlt/bHzRSR9K5qa0x9mEfnGZRj14x3WGIlNfSmS88a4wR0iXP+UtEASXwIDAQABAoIBABoUDV3tNhk/aLjzw/daAh+UcTYqcpMjZhRNlb8gGsMocBL+lKGzdBAwITfn4OSxGAbB9beVbDGXt8TWe/z9iLfaPUVsDj7D0ZbYnuZm2sB9IsU2jIepoJx1G5bJgv9bye4CqorzwQxwFztKIHcmfFCKOfQmN/f2Gv/WgdX+mgvpaNPfHn3gnnNpmhq2Mb8QiNhLvPMhKXORk08a2xGLUEO43Iw7YSg8lZ4L6So/ROSPCtyErAgVTxAkfKoFyJAaHykKPeeLt9Zy5aPndmizIWwCMXb0+Zq7bDmKS3yvq8Nr95ZjOc09GD1nit2vWLCPJxNc2lzkSBxIAmDmXta0+LECgYEA6LMl02AgB47zxXfvhDSpW3SABrzU8GyOnZjcYdf2cZkDtgx0QE3g5E2YeWFcSny3+x3ug6nVjjIscZaBap3DxwMft2PrGX7Bfmecg3hJ+TF7KLFuAkmG+qqzkPeox9YTs/aRBRVBu9oo+Io44HqryIoehCpX59+CjqznhLEX8FcCgYEAqPqR3wJPI38XqGLKoxpI20l99rGD/mj2J3PvFXGtti3e5ZB4irsxO8U38vHEw/2kEXEJZW+2trS486pXaKE8JfrRItjaCUbgUSO+3lCpdkl9gXfimEr0xD2zPf92aAcxqQiG65I+k2Ch5ba2YmteTGi55SB6FuEAvHtGhRf7iTkCgYEAlFQ1pVJduFOwIcx8uaoT1j8hqKnPll2sXtrkh93wsqKV0gKIS8EYvI6VxbGA8d4kLIb81aJ5hUWIPPNyFTLxa7cbDXw8jSjWUCvdgZQ4mwamed73v698weXzxlGHnbJhJtLhx/qvxv2eJid9b+HiBFe+cgLHu/8mKqoefd+g4csCgYAAsOWfz9abAo4KNj015Ymeu/Iz7A3qIGvBRYwYvlpDgHSE485aYuGUqP3NlIeFdagSGjA7pfVNUfffpzasStyAG0J3rgNWPl/0dPz208WdojdNLDxU+xl9I/NzsXO+gSkG0+4ZUIPI/oAq/FBKnr3H+jWoZjWZmlnya16idLKmoQKBgEU5tCv8klIdPLwHu0UOrMWtbPnWABzCu2eJdova0zJblHE5KTgu3U2td/wEfiDJVYpnK3XSZkxhjJWKxxtWSQwUjZwK00pc2J/wxW1wPq7uXI0/jiJT/C4BDwDor20wMm8hqOgdCxPT1yIbD25WUF7QVRBUuX/wQHJfdrsPKzUe" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	// var client, err = alipay.New("2021001144694156", privateKey, true)
	// client.LoadAliPayPublicKey("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhb/KlxYfhRE8KRp92MQM8ZB8NVjoM9LYFOnPIuNtcMZVA8ld7ybDP2FiA+QEE7wLGqMImwl1Y4xzkrTLCjHVC8fdR8ZvzZR2I3ZOrARerI9+RbkCfT+7YLv55+A+WTHEyiB+v7PfXVTT28s0CHNLPXMyQD1u8UVEQEpbMSs8hH3pJF55Li7kc5VvJpV3RVO9TXZTVAA5mSp9FvO3u+47IJDgFVLnqqHh6ETL1nHVpxiAY2LGer+RWpVYD8v+We+VWsrfJP7bO0xr2pwizldepo8YNYPgcIAIwd7KiveypL1pA0xWgSjUHzrkVh1j/nSnvJgKSdydU/VRcaVt/Mt8wwIDAQAB")

	// 将 key 的验证调整到初始化阶段
	if err != nil {
		fmt.Println(err)
		return
	}

	//计算总价格
	// var TotalAmount float64
	// for i := 0; i < len(orderitem); i++ {
	// 	TotalAmount = TotalAmount + orderitem[i].ProductPrice
	// }
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://tsung.top:8001/alipayNotify"
	p.ReturnURL = "http://tsung.top:8001/alipayReturn"
	p.TotalAmount = "0.01"
	p.Subject = "订单——" + time.Now().Format("200601021504")
	p.OutTradeNo = "WF" + time.Now().Format("200601021504") + "_" + strconv.Itoa(AliId)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var url, err4 = client.TradePagePay(p)
	if err4 != nil {
		fmt.Println(err4)
	}
	var payURL = url.String()
	c.Redirect(payURL, 302)

}

func (c *PayController) AlipayNotify() {
	var privateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCPUi2yj4RJcqAKfjV/AxjqIjjUq007pGhLClbrwEoCVhgnQU9bAVFnPlDaVdO2xBfu8D/gHCwV9czZKJh51yg1kwE6Yv/hSd966PnGS9NszfxxWfTbWeC6DIHZj66nTECK+vYWX36tKxIG+juzXyfoAuyL1h58oFgF/Zwa1FzpKDTcMUIE+npf3DpMS0Uatf6TsREzVQQs3i0WVxYY3lv1Dmualr6Q3GAFE1j/xt/STE993uh8MVeLS+RcTrrjilPSVH5Z0DLAjkDSH7XUK1lIHASpgOrddEJ8MeT8L7bUJ8nDs5qQ8zfbPVJaXfsF2NYS162HkIW2bl0r5c5Mudt7AgMBAAECggEAWUHVweHNgiyH7WECkhJsvswHVrNEi0NtzGYpEfOUY/YYXsI22Lduaf0OP5u6GZXwTdeEAF+rORX2uLumkiLkINFnr2QedcEbFCHqBIwOpTF36WQbsUw9P8EwUT1BiWFcxPFctzxL2S78sCnBaol1gfHoPYJhRD5b84cpZDAjmPSJk1XAtAtKChUIskLBAsCvwlGHbx/6UQwM9eKgwo0Y67MCPW9wBjE9bRFWBfaeLszEVu3nKyOKLwcUGDXrbmBS8bj9YtyqTG7RjKZmIGuJtPKehEjlNn2ALYYMHUA5VSdVP4LrBiLVLQE5tTDedAzr3uHuWOgLaMeDnfwQg8Hj6QKBgQDGaCoZ3JRg4PbGetduHpGfWOstwZKNwqkDeYvpA03GcRYhzhqG98qdTGEYaLGkszQrh0ZcYk7fs8jsR5Lu2WTzImtiAYDCDEjSDHj5N5PfWR0r2pxrUvX8shun4QKpz4QQ0RUjeujZ2hJHkeCviF2+k5IBDvz6YtBo+H3IgR9QjwKBgQC47IVtuvlYZ0/2TAcAt40YhEgLZOr3NuR4eSnx76zf/8vRmHDEfIqvUdIzzJy3RVYT3uXiB7DYwaHr76ouP3lMOhgLgUlK4Dt8L3UbMP6Asr+6D/uggVmlIHKK76HZIdBL5nGQvTOvwE+fmEivlR5QV+cczqQdCYyNZXETmgAkVQKBgF32JMIcqZR71cLHmFDJX1Okq7P+sWY7YwmHPZA7hVDOa5nU3tE+dpEqA+2oX0DNsY5PwS2tTQc6QJRNjTNadymCCnLenVjIso/vYjc8b+ZdcKg9HsjhACgNPXWy5S0AXt4L9sPXyICreu60EkFvBl5jysh/jaUSuPqNfBxBsk/XAoGBAKYxkzzp4/wKZXfSHh0L2VemUuVCnlTtVWncYtEXeQObbX8CBJ7h2vXzj/mTs2iWfOTA11NLXCmB5FcZfpWv4ACc2U1FtSwA2BUkxZdZcfESNHMwuBEpDvrzbV3mPUvaMsxz366YC+Kw8B5biz+ZwbOtPHzMTfv2wAW3nGdkaSo9AoGBALW8kQ6QwXu3gB7q8Es1zCvSg07tnJks57nJDTtOOg6oc52RvKjndWHT6aFwp/YrESeHNvoudRhPm5nKUTbKDiLixDXcrvEU3b70dJGwWHkBdYU/s5J2o5oaclwiyiUptnIp2eA+Cu+wH1eREGmrOAdeBSFOfdjcfJTFmOCoalJR" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New("2021001186696588", privateKey, true)

	client.LoadAppPublicCertFromFile("crt/appCertPublicKey_2021001186696588.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("crt/alipayRootCert.crt")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("crt/alipayCertPublicKey_RSA2.crt")       // 加载支付宝公钥证书

	if err != nil {
		fmt.Println(err)
		return
	}

	req := c.Ctx.Request
	req.ParseForm()
	ok, err := client.VerifySign(req.Form)
	if !ok || err != nil {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	rep := c.Ctx.ResponseWriter
	var noti, _ = client.GetTradeNotification(req)
	if noti != nil {
		fmt.Println("交易状态为:", noti.TradeStatus)
		if string(noti.TradeStatus) == "TRADE_SUCCESS" {
			order := models.Order{}
			temp := strings.Split(noti.OutTradeNo, "_")[1]
			id, _ := strconv.Atoi(temp)
			models.DB.Where("id=?", id).Find(&order)
			order.PayStatus = 1
			order.OrderStatus = 1
			models.DB.Save(&order)
		}
	}
	alipay.AckNotification(rep) // 确认收到通知消息
}
func (c *PayController) AlipayReturn() {
	c.Redirect("/user/order", 302)
}

func (c *PayController) WxPay() {
	WxId, err := c.GetInt("id")
	if err != nil {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	orderitem := []models.OrderItem{}
	models.DB.Where("order_id=?", WxId).Find(&orderitem)
	//1、配置基本信息
	account := wxpay.NewAccount(
		"wx7bf3787c783116e4",
		"1502539541",
		"zhongyuantengitying6666666666666",
		false,
	)
	client := wxpay.NewClient(account)
	var price int64
	for i := 0; i < len(orderitem); i++ {
		price = 1
	}
	//2、获取ip地址   订单号等信息
	ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	template := "200601021504"
	tradeNo := time.Now().Format(template)
	//3、调用统一下单
	params := make(wxpay.Params)
	params.SetString("body", "订单——"+time.Now().Format(template)).
		SetString("out_trade_no", tradeNo+"_"+strconv.Itoa(WxId)).
		SetInt64("total_fee", price).
		SetString("spbill_create_ip", ip).
		SetString("notify_url", "http://tsung.top:8001/wxpay/notify").
		// SetString("trade_type", "APP")
		SetString("trade_type", "NATIVE") //网站支付需要改为NATIVE

	p, err1 := client.UnifiedOrder(params)
	beego.Info(p)
	if err1 != nil {
		beego.Error(err1)
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	//4、获取code_url 生成支付二维码

	var pngObj []byte
	beego.Info(p)
	pngObj, _ = qrcode.Encode(p["code_url"], qrcode.Medium, 256)
	c.Ctx.WriteString(string(pngObj))
}

func (c *PayController) WxPayNotify() {
	//1、获取表单传过来的xml数据  在配置文件里设置 copyrequestbody = true
	xmlStr := string(c.Ctx.Input.RequestBody)
	postParams := wxpay.XmlToMap(xmlStr)
	beego.Info(postParams)

	//2、校验签名
	account := wxpay.NewAccount(
		"wx7bf3787c783116e4",
		"1502539541",
		"zhongyuantengitying6666666666666",
		false,
	)
	client := wxpay.NewClient(account)
	isValidate := client.ValidSign(postParams)
	// xml解析
	params := wxpay.XmlToMap(xmlStr)
	beego.Info(params)
	if isValidate == true {
		if params["return_code"] == "SUCCESS" {
			idStr := strings.Split(params["out_trade_no"], "_")[1]
			id, _ := strconv.Atoi(idStr)
			order := models.Order{}
			models.DB.Where("id=?", id).Find(&order)
			order.PayStatus = 1
			order.PayType = 1
			order.OrderStatus = 1
			models.DB.Save(&order)
		}
	} else {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
}
