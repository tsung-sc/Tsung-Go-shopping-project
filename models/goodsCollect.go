package models

import (
	_ "github.com/jinzhu/gorm"
)

type GoodsCollect struct {
	Id      int
	UserId  int
	GoodId  int
	AddTime string
}

func (GoodsCollect) TableName() string {
	return "goods_collect"
}
