package entity

import (
	"akita/panda-im/common/models"
	"gorm.io/gorm"
)

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

// IsFriend 判断是否是好友
func (m *UserDao) IsFriend(db *gorm.DB, s, r int) bool {
	var count int64
	// 查询好友关系表中是否存在 A 和 B 之间的好友关系
	err := db.Model(m).
		Where("(send_user_id = ? AND rev_user_id = ?) OR (send_user_id = ? AND rev_user_id = ?)", s, r, r, s).
		Count(&count).
		Error
	if err != nil {
		// 发生错误，可能是数据库查询失败
		return false
	}
	// 如果好友关系存在（count > 0），则说明 A 和 B 是好友关系，否则不是
	return count > 0
}
