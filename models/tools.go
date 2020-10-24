package models

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/gomarkdown/markdown"
	. "github.com/hunterhug/go_image"
)

func UnixToDate(timestamp int) string {

	t := time.Unix(int64(timestamp), 0)

	return t.Format("2006-01-02 15:04:05")
}

//2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		beego.Info(err)
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// GetSettingFromColumn("SiteTitle")
func GetSettingFromColumn(columnName string) string {

	//redis file
	setting := Setting{}
	DB.First(&setting)
	//反射来获取
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func GetOrderId() string {
	template := "200601021504"
	return time.Now().Format(template) + GetRandomNum()
}
func ResizeImage(filename string) {
	extName := path.Ext(filename)
	resizeImage := strings.Split(beego.AppConfig.String("resizeImageSize"), ",")

	for i := 0; i < len(resizeImage); i++ {
		w := resizeImage[i]
		width, _ := strconv.Atoi(w)
		savepath := filename + "_" + w + "x" + w + extName
		err := ThumbnailF2F(filename, savepath, width, width)
		if err != nil {
			beego.Error(err)
		}
	}

}

//格式化图片
func FormatImg(picName string) string {
	ossStatus, err := beego.AppConfig.Bool("ossStatus")
	if err != nil {
		//判断目录前面是否有/
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName
	}
	if ossStatus {
		return beego.AppConfig.String("ossDomain") + "/" + picName
	} else {
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName

	}

}

/*
# 二级标题

<h2>二级标题<h2>
*/
func FormatAttr(str string) string {
	md := []byte(str)
	htmlByte := markdown.ToHTML(md, nil, nil)
	return string(htmlByte)
}

//乘法的函数
func Mul(price float64, num int) float64 {
	return price * float64(num)
}

//封装一个生产随机数的方法
func GetRandomNum() string {
	var str string
	for i := 0; i < 4; i++ {
		current := rand.Intn(10) //0-9   "math/rand"
		str += strconv.Itoa(current)
	}
	return str
}

func SendMsg(str string) {

	//1、正式环境
	// clnt := ypclnt.New("62c24eee15fxxxxxxxxxxxxxxxe0beca9e0a") //apikey https://www.yunpian.com/官网获取
	// param := ypclnt.NewParam(2061690)                          //模板id
	// param[ypclnt.MOBILE] = "15029745801"
	// param[ypclnt.TEXT] = "【IT营】您的验证码是888897"
	// r := clnt.Sms().SingleSend(param)
	// fmt.Println(r)

	//2、测试阶段保存短信到文件里面
	ioutil.WriteFile("code.txt", []byte(str), 06666)
}

func SendEmail(str string) {
	beego.Info("发送验证邮件成功")
}
