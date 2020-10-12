package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {
	mysqladmin := beego.AppConfig.String("mysqladmin")
	mysqlpwd := beego.AppConfig.String("mysqlpwd")
	mysqldb := beego.AppConfig.String("mysqldb")
	DB, err = gorm.Open("mysql", mysqladmin+":"+mysqlpwd+"@/"+mysqldb+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		beego.Error(err)
		beego.Error("连接MySql数据库失败")
	} else {
		beego.Info("连接MySql数据库成功")
	}
}
