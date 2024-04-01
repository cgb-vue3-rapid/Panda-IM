package models

import "akita/panda-im/common/models"

// UserModel 用户配置表
type UserModel struct {
	models.Model
	NickName string `gorm:"column:nickname;comment:'用户昵称'" json:"nick_name"`
	PassWord string `gorm:"column:password;comment:'用户密码'" json:"password"`
	Avatar   string `gorm:"column:avatar;comment:'用户头像地址'" json:"avatar"`
	Gender   int8   `gorm:"column:gender;comment:'用户性别（0未知，1男，2女）'" json:"gender"`
	Abstract string `gorm:"column:abstract;comment:'用户简介'" json:"abstract"`
	IP       string `gorm:"column:ip;comment:'用户IP'" json:"ip"`
	Addr     string `gorm:"column:Addr;comment:'用户地址'" json:"addr"`
}

func (U *UserModel) TableName() string {
	return "user"
}
