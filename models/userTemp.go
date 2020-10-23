package models

import (
	_ "github.com/jinzhu/gorm"
)

type UserTemp struct {
	Id        int
	Nickname  string
	Ip        string
	Phone     string
	SendCount int
	AddDay    string
	AddTime   int
	Sign      string
}

func (UserTemp) TableName() string {
	return "user_temp"
}
