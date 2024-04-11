package entity

import (
	"akita/panda-im/common/models"
	"context"
	"gorm.io/gorm"
)

// UserModel 用户表
type UserModel struct {
	models.Model
	NickName     string `gorm:"column:nickname;type:varchar(255);comment:'用户昵称';not null" json:"nickname"`
	Mobile       string `gorm:"column:mobile;type:varchar(255);comment:'用户手机号';index" json:"mobile"`
	PassWord     string `gorm:"column:password;type:varchar(255);comment:'用户密码';not null" json:"password"`
	Avatar       string `gorm:"column:avatar;type:varchar(255);comment:'用户头像地址'" json:"avatar"`
	Gender       int32  `gorm:"column:gender;type:tinyint(1);comment:'用户性别（0未知，1男，2女)';default:0" json:"gender"`
	Abstract     string `gorm:"column:abstract;type:text;comment:'用户简介'" json:"abstract"`
	IP           string `gorm:"column:ip;type:varchar(45);comment:'用户IP'" json:"ip"`
	Addr         string `gorm:"column:Addr;type:varchar(255);comment:'用户地址'" json:"addr"`
	Role         int32  `gorm:"column:role;type:tinyint(1);comment:'用户角色(1:普通用户,2:管理员)';default:0" json:"role"`
	ThirdPartyID string `gorm:"column:third_party_id;size:64;type:varchar(255);comment:'第三方登录ID'" json:"third_party_id"`
}

func (U *UserModel) TableName() string {
	return "user"
}

type UserModelDao struct {
	db *gorm.DB
}

func NewUserModelDao(db *gorm.DB) *UserModelDao {
	return &UserModelDao{
		db: db,
	}
}

func (m *UserModelDao) FindByMobile(ctx context.Context, mobile string) (*UserModel, error) {
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
func (m *UserModelDao) Insert(ctx context.Context, data *UserModel) error {
	return m.db.WithContext(ctx).Create(&data).Error
}

func (m *UserModelDao) FindByID(ctx context.Context, id int64) (*UserModel, error) {
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
func (m *UserModelDao) GetUserInfoByID(ctx context.Context, userID int64) (*UserModel, error) {
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
