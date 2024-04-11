package entity

import (
	"akita/panda-im/common/models"
)

// UserModel 用户表
type UserModel struct {
	models.Model
	//UserName     string `gorm:"column:username;type:varchar(255);comment:'用户名';index" json:"username"`
	NickName     string `gorm:"column:nickname;type:varchar(255);comment:'用户昵称';not null" json:"nickname"`
	Mobile       string `gorm:"column:mobile;type:varchar(255);comment:'用户手机号';index" json:"mobile"`
	PassWord     string `gorm:"column:password;type:varchar(255);comment:'用户密码';not null" json:"password"`
	Avatar       string `gorm:"column:avatar;type:varchar(255);comment:'用户头像地址'" json:"avatar"`
	Gender       int8   `gorm:"column:gender;type:tinyint(1);comment:'用户性别（0未知，1男，2女)';default:0" json:"gender"`
	Abstract     string `gorm:"column:abstract;type:text;comment:'用户简介'" json:"abstract"`
	IP           string `gorm:"column:ip;type:varchar(45);comment:'用户IP'" json:"ip"`
	Addr         string `gorm:"column:Addr;type:varchar(255);comment:'用户地址'" json:"addr"`
	Role         int8   `gorm:"column:role;type:tinyint(1);comment:'用户角色(1:普通用户,2:管理员)';default:0" json:"role"`
	ThirdPartyID string `gorm:"column:third_party_id;size:64;type:varchar(255);comment:'第三方登录ID'" json:"third_party_id"`
}

func (U *UserModel) TableName() string {
	return "user"
}
