package entity

import "akita/panda-im/common/models"

// AuthModel 用户表
type AuthModel struct {
	models.Model
	NickName string `gorm:"column:nickname;type:varchar(255);comment:'用户昵称'" json:"nick_name"`
	PassWord string `gorm:"column:password;type:varchar(255);comment:'用户密码';not null" json:"password"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);comment:'用户头像地址'" json:"avatar"`
	Gender   int8   `gorm:"column:gender;type:tinyint(1);comment:'用户性别（0未知，1男，2女)';default:0" json:"gender"`
	Abstract string `gorm:"column:abstract;type:text;comment:'用户简介'" json:"abstract"`
	IP       string `gorm:"column:ip;type:varchar(45);comment:'用户IP'" json:"ip"`
	Addr     string `gorm:"column:Addr;type:varchar(255);comment:'用户地址'" json:"addr"`
}

func (U *AuthModel) TableName() string {
	return "auth"
}
