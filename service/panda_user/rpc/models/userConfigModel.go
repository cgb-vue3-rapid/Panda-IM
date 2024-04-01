package models

import "akita/panda-im/common/models"

// UserConfModel 用户配置表
type UserConfModel struct {
	models.Model
	UserId               int64                  `gorm:"column:user_id;comment:'用户ID';primary_key" json:"user_id"`
	RecallMessage        *string                `gorm:"column:recall_message;comment:'撤回消息'" json:"recall_message"`
	Oline                bool                   `gorm:"column:online;comment:'在线状态'" json:"online"`
	FriendsOnline        bool                   `gorm:"column:friends_online;comment:'好友在线'" json:"friends_online"`
	AllSounds            bool                   `gorm:"column:all_sounds;comment:'所有声音'" json:"all_sounds"`
	SecureLink           bool                   `gorm:"column:secure_link;comment:'安全链接'" json:"secure_link"`
	SavePwd              bool                   `gorm:"column:save_pwd;comment:'记住密码'" json:"save_pwd"`
	SearchUser           int8                   `gorm:"column:search_user;comment:'搜索用户'" json:"searchUser"`                      // 别人查找到你的方式，0不允许别人查找到我，1通过用户号找到我，2可以通过昵称搜索到我
	FriendVerification   int8                   `gorm:"column:friend_verification;comment:'好友验证方式'" json:"friend_verification"`   // 好友验证方式，0不允许任何人添加，1允许任何人添加，2需要验证消息，3需要回答问题，4需要正确回答问题
	VerificationQuestion []VerificationQuestion `gorm:"column:verification_question;comment:'验证问题'" json:"verification_question"` // 验证问题为3和4的时候需要
}

func (U *UserConfModel) TableName() string {
	return "user_conf"
}
