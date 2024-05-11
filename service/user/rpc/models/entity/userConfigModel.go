package entity

import (
	"akita/panda-im/common/models"
	"akita/panda-im/common/models/ctype"
	"context"
	"gorm.io/gorm"
)

// UserConfModel 用户配置表
type UserConfModel struct {
	models.Model
	UserId        int64     `gorm:"column:user_id;comment:'用户ID';not null;uniqueIndex" json:"user_id"`
	UserModel     UserModel `gorm:"foreignKey:UserId" json:"-"`
	RecallMessage string    `gorm:"column:recall_message;comment:'撤回消息';default:'撤回了一条消息'" json:"recall_message"`
	Oline         bool      `gorm:"column:online;comment:'在线状态';default:true" json:"online"`
	FriendsOnline bool      `gorm:"column:friends_online;comment:'好友上线提醒'" json:"friends_online"`
	AllSounds     bool      `gorm:"column:all_sounds;comment:'所有声音'" json:"all_sounds"`
	SecureLink    bool      `gorm:"column:secure_link;comment:'安全链接'" json:"secure_link"`
	SavePwd       bool      `gorm:"column:save_pwd;comment:'记住密码'" json:"save_pwd"`
	SearchUser    int32     `gorm:"column:search_user;comment:'搜索用户'" json:"searchUser"` // 别人查找到你的方式，0不允许别人查找到我，1通过用户号找到我，2可以通过手机号搜索到我
	// 好友验证
	VerificationType     int32                       `gorm:"column:verification_type;comment:'好友验证方式'" json:"verification_type"` // 好友验证方式，0不允许任何人添加，1允许任何人添加，2需要验证消息，3需要回答问题，4需要正确回答问题
	VerificationQuestion *ctype.VerificationQuestion `json:"verification_question"`
}

func (U *UserConfModel) TableName() string {
	return "user_conf"
}

// GetConfigByUserId 通过用户ID查询用户配置表
func (m *UserDao) GetConfigByUserId(ctx context.Context, userId int64) (*UserConfModel, error) {
	var user UserConfModel
	err := m.db.WithContext(ctx).First(&user, userId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

// UpdateConfig 更新用户配置信息
func (m *UserDao) UpdateConfig(ctx context.Context, userId int64, userConf *UserConfModel) error {
	return m.db.WithContext(ctx).Model(&UserConfModel{}).Where("user_id = ?", userId).Updates(userConf).Error
}
