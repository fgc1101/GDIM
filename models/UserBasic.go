package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name       string
	Password   string
	Phone      string
	Email      string
	Identity   string
	ClientIp   string
	ClientPort string
}

func (table *UserBasic) TableName() string {
	return "im_user_basic"
}
