package entity

import "akita/panda-im/common/models"

// UserFriendModel 用户好友表
type UserFriendModel struct {
	models.Model
	SendUserId       int64     `gorm:"column:send_user_id;comment:'发送用户ID';not null;uniqueIndex" json:"send_user_id"`
	SendUserModel    UserModel `gorm:"foreignKey:SendUserId" json:"-"`
	ReceiveUserId    int64     `gorm:"column:receive_user_id;comment:'接收用户ID';not null;uniqueIndex" json:"receive_user_id"`
	ReceiveUserModel UserModel `gorm:"foreignKey:ReceiveUserId" json:"-"`
	Notice           string    `gorm:"column:notice;comment:'备注'" json:"notice"`
}

func (U *UserFriendModel) TableName() string {
	return "user_friend"
}
