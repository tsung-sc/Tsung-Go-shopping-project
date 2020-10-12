package models

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

var Cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	Cpt = captcha.NewWithFilter("/captcha/", store)
	Cpt.ChallengeNums = 4
	Cpt.StdWidth = 100
	Cpt.StdHeight = 40
}
