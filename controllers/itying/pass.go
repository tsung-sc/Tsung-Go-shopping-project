package itying

import (
	"regexp"
	"strings"
	"xiaomi/models"
)

type PassController struct {
	BaseController
}

func (c *PassController) Login() {
	c.Data["prevPage"] = c.Ctx.Request.Referer()
	c.TplName = "itying/pass/login.html"
}

func (c *PassController) DoLogin() {
	phone := c.GetString("phone")
	password := c.GetString("password")
	photo_code := c.GetString("photo_code")
	photoCodeId := c.GetString("photoCodeId")
	identifyFlag := models.Cpt.Verify(photoCodeId, photo_code)
	if !identifyFlag {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "输入的图形验证码不正确",
		}
		c.ServeJSON()
		return
	}
	password = models.Md5(password)
	user := []models.User{}
	models.DB.Where("phone=? AND password=?", phone, password).Find(&user)
	if len(user) > 0 {
		models.Cookie.Set(c.Ctx, "userinfo", user[0])
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"msg":     "用户登陆成功",
		}
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "用户名或密码不正确",
		}
		c.ServeJSON()
		return
	}
}

func (c *PassController) LoginOut() {
	models.Cookie.Remove(c.Ctx, "userinfo", "")
	c.Redirect(c.Ctx.Request.Referer(), 302)
}
func (c *PassController) RegisterStep1() {
	c.TplName = "itying/pass/register_step1.html"
}

func (c *PassController) RegisterStep2() {
	sign := c.GetString("sign")
	photo_code := c.GetString("photo_code")
	//验证图形验证码和前面是否正确
	sessionPhotoCode := c.GetSession("photo_code")
	if photo_code != sessionPhotoCode {
		c.Redirect("/pass/registerStep1", 302)
		return
	}
	userTemp := []models.UserTemp{}
	models.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) > 0 {
		c.Data["sign"] = sign
		c.Data["photo_code"] = photo_code
		c.Data["phone"] = userTemp[0].Phone
		c.TplName = "itying/pass/register_step2.html"
	} else {
		c.Redirect("/pass/registerStep1", 302)
		return
	}
}
func (c *PassController) RegisterStep3() {
	sign := c.GetString("sign")
	sms_code := c.GetString("sms_code")
	sessionSmsCode := c.GetSession("sms_code")
	if sms_code != sessionSmsCode {
		c.Redirect("/pass/registerStep1", 302)
		return
	}
	userTemp := []models.UserTemp{}
	models.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) > 0 {
		c.Data["sign"] = sign
		c.Data["sms_code"] = sms_code
		c.TplName = "itying/pass/register_step3.html"
	} else {
		c.Redirect("/pass/registerStep1", 302)
		return
	}
}

func (c *PassController) SendCode() {
	phone := c.GetString("phone")
	photo_code := c.GetString("photo_code")
	photoCodeId := c.GetString("photoCodeId")
	if photoCodeId == "resend" {
		//session里面验证验证码是否合法
		sessionPhotoCode := c.GetSession("photo_code")
		if sessionPhotoCode != photo_code {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"msg":     "输入的图形验证码不正确,非法请求",
			}
			c.ServeJSON()
			return
		}
	}
	if !models.Cpt.Verify(photoCodeId, photo_code) {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "输入的图形验证码不正确",
		}
		c.ServeJSON()
		return
	}

	c.SetSession("photo_code", photo_code)
	pattern := `^[\d]{11}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(phone) {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "手机号格式不正确",
		}
		c.ServeJSON()
		return
	}
	user := []models.User{}
	models.DB.Where("phone=?", phone).Find(&user)
	if len(user) > 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "此用户已存在",
		}
		c.ServeJSON()
		return
	}

	add_day := models.GetDay()
	ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	sign := models.Md5(phone + add_day) //签名
	sms_code := models.GetRandomNum()
	userTemp := []models.UserTemp{}
	models.DB.Where("add_day=? AND phone=?", add_day, phone).Find(&userTemp)
	var sendCount int
	models.DB.Where("add_day=? AND ip=?", add_day, ip).Table("user_temp").Count(&sendCount)
	//验证IP地址今天发送的次数是否合法
	if sendCount <= 10 {
		if len(userTemp) > 0 {
			//验证当前手机号今天发送的次数是否合法
			if userTemp[0].SendCount < 5 {
				models.SendMsg(sms_code)
				c.SetSession("sms_code", sms_code)
				oneUserTemp := models.UserTemp{}
				models.DB.Where("id=?", userTemp[0].Id).Find(&oneUserTemp)
				oneUserTemp.SendCount += 1
				models.DB.Save(&oneUserTemp)
				c.Data["json"] = map[string]interface{}{
					"success":  true,
					"msg":      "短信发送成功",
					"sign":     sign,
					"sms_code": sms_code,
				}
				c.ServeJSON()
				return
			} else {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"msg":     "当前手机号今天发送短信数已达上限",
				}
				c.ServeJSON()
				return
			}

		} else {
			models.SendMsg(sms_code)
			c.SetSession("sms_code", sms_code)
			//发送验证码 并给userTemp写入数据
			oneUserTemp := models.UserTemp{
				Ip:        ip,
				Phone:     phone,
				SendCount: 1,
				AddDay:    add_day,
				AddTime:   int(models.GetUnix()),
				Sign:      sign,
			}
			models.DB.Create(&oneUserTemp)
			c.Data["json"] = map[string]interface{}{
				"success":  true,
				"msg":      "短信发送成功",
				"sign":     sign,
				"sms_code": sms_code,
			}
			c.ServeJSON()
			return
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "此IP今天发送次数已经达到上限，明天再试",
		}
		c.ServeJSON()
		return
	}

}

func (c *PassController) ValidateSmsCode() {
	sign := c.GetString("sign")
	sms_code := c.GetString("sms_code")

	userTemp := []models.UserTemp{}
	models.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) == 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "参数错误",
		}
		c.ServeJSON()
		return
	}

	sessionSmsCode := c.GetSession("sms_code")
	if sessionSmsCode != sms_code {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "输入的短信验证码错误",
		}
		c.ServeJSON()
		return
	}

	nowTime := models.GetUnix()
	if (nowTime-int64(userTemp[0].AddTime))/1000/60 > 15 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "验证码已过期",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"msg":     "验证成功",
	}
	c.ServeJSON()
}

func (c *PassController) DoRegister() {
	sign := c.GetString("sign")
	sms_code := c.GetString("sms_code")
	password := c.GetString("password")
	rpassword := c.GetString("rpassword")
	sessionSmsCode := c.GetSession("sms_code")
	if sms_code != sessionSmsCode {
		c.Redirect("/pass/registerStep1", 302)
		return
	}
	if len(password) < 6 {
		c.Redirect("/pass/registerStep1", 302)
	}
	if password != rpassword {
		c.Redirect("/pass/registerStep1", 302)
	}
	userTemp := []models.UserTemp{}
	models.DB.Where("sign=?", sign).Find(&userTemp)
	ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	if len(userTemp) > 0 {
		user := models.User{
			Phone:    userTemp[0].Phone,
			Password: models.Md5(password),
			LastIp:   ip,
		}
		models.DB.Create(&user)

		models.Cookie.Set(c.Ctx, "userinfo", user)
		c.Redirect("/", 302)
	} else {
		c.Redirect("/pass/registerStep1", 302)
	}

}
