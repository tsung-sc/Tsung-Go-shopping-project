package models

import (
	_ "github.com/jinzhu/gorm"
)

type Comment struct {
	Id        int
	UserId    int
	OrderId   int
	Star      int
	Str       string
	Text      string
	GoodId    int
	AddTime   int64
	OrderItem OrderItem
	UserName  string `gorm:"-"`
}

func (Comment) TableName() string {
	return "comment"
}
