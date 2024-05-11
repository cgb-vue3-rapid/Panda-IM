package models

import "time"

// Model 结构体定义
type Model struct {
	// ID 字段，类型为 uint，作为主键
	ID int64 `gorm:"primaryKey;comment:'ID主键';not null;autoIncrement;" json:"id"`
	// 创建时间字段，类型为 time.Time
	CreatedAt time.Time `gorm:"comment:'记录创建时间'" json:"created_at"`
	// 更新时间字段，类型为 time.Time
	UpdatedAt time.Time `gorm:"comment:'记录更新时间'" json:"updated_at"`
	// 删除时间字段，类型为 time.Time
	DeletedAt time.Time `gorm:"comment:'记录删除时间';default:null" json:"deleted_at"`
	// 是否删除字段，类型为 bool
	IsDeleted bool `gorm:"comment:'标记是否已删除(0:未删除，1：删除)';index;default:0" json:"is_deleted"`
}
