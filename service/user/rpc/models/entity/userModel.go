package entity

import (
	"akita/panda-im/common/models"
	"context"
)

type Gender int32

const (
	Unknown Gender = iota
	Male
	Female
)

var genderLabels = map[Gender]string{
	Unknown: "未设置",
	Male:    "男",
	Female:  "女",
}

func (g Gender) String() string {
	return genderLabels[g]
}

// UserModel 用户表
type UserModel struct {
	models.Model
	NickName string `gorm:"column:nickname;type:varchar(255);comment:'用户昵称';not null" json:"nickname"`
	Mobile   string `gorm:"column:mobile;type:varchar(255);comment:'用户手机号';index" json:"mobile"`
	PassWord string `gorm:"column:password;type:varchar(255);comment:'用户密码';not null" json:"-"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);comment:'用户头像地址'" json:"avatar"`
	Gender   Gender `gorm:"column:gender;type:tinyint(1);comment:'用户性别（0未知，1男，2女)';default:0" json:"gender"`
	Abstract string `gorm:"column:abstract;type:text;comment:'用户简介'" json:"abstract"`
	IP       string `gorm:"column:ip;type:varchar(45);comment:'用户IP'" json:"ip"`
	Addr     string `gorm:"column:Addr;type:varchar(255);comment:'用户地址'" json:"addr"`
	Role     int32  `gorm:"column:role;type:tinyint(1);comment:'用户角色(1:普通用户,2:管理员)';default:0" json:"role"`
	// 添加 UserConfModel 的引用字段
	UserConf *UserConfModel `gorm:"foreignKey:UserId" json:"-"`
}

func (U *UserModel) TableName() string {
	return "user"
}

// FindByMobile 根据手机号查询
func (m *UserDao) FindByMobile(ctx context.Context, mobile string) (*UserModel, error) {
	var userModel UserModel
	err := m.db.WithContext(ctx).
		Where("mobile = ?", mobile).
		First(&userModel).
		Error

	if err != nil {
		return nil, err
	}
	return &userModel, nil
}

// Insert 插入
func (m *UserDao) Insert(ctx context.Context, data *UserModel) error {
	return m.db.WithContext(ctx).Create(&data).Error
}

// FindByID 根据id查询
func (m *UserDao) FindByID(ctx context.Context, id int64) (*UserModel, error) {
	var userModel UserModel
	err := m.db.WithContext(ctx).
		Where("id = ?", id).
		First(&userModel).
		Error

	if err != nil {
		return nil, err
	}
	return &userModel, nil
}

// GetUserInfoByID 获取用户信息
func (m *UserDao) GetUserInfoByID(ctx context.Context, userID int64) (*UserModel, error) {
	var userModel UserModel

	err := m.db.WithContext(ctx).
		Where("id = ?", userID).
		First(&userModel).
		Error

	if err != nil {
		return nil, err
	}

	return &userModel, nil
}

// UpdateUserInfo 跟新用户信息
func (m *UserDao) UpdateUserInfo(ctx context.Context, userID int64, data *UserModel) error {
	return m.db.WithContext(ctx).Model(&UserModel{}).
		Where("id = ?", userID).
		Updates(data).
		Error
}
