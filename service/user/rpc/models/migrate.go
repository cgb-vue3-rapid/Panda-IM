package models

import (
	"akita/panda-im/service/user/rpc/models/entity"
	"gorm.io/gorm"
)

// Migrate 用于执行数据库迁移
func Migrate(db *gorm.DB) error {
	// 在这里添加你的表迁移逻辑，例如：
	err := db.AutoMigrate(
		&entity.UserModel{},
		&entity.UserConfModel{},
		//&entity.FriendVerifyModel{},
		//&entity.UserFriendModel{},
	)
	if err != nil {
		return err
	}
	return nil
}
