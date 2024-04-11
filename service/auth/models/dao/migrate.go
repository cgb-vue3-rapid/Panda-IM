package dao

import (
	"akita/panda-im/service/auth/models/entity"
	"gorm.io/gorm"
)

// Migrate 用于执行数据库迁移
func Migrate(db *gorm.DB) error {
	// 在这里添加你的表迁移逻辑，例如：
	err := db.AutoMigrate(
		&entity.UserModel{},
	)
	if err != nil {
		return err
	}
	return nil
}
