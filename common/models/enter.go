package models

import "time"

// Model 结构体定义
type Model struct {
	// ID 字段，类型为 uint，作为主键
	ID uint `gorm:"primary_key" json:"id" comment:"ID 主键"`
	// 创建时间字段，类型为 time.Time
	CreatedAt time.Time `json:"created_at" comment:"记录创建时间"`
	// 更新时间字段，类型为 time.Time
	UpdatedAt time.Time `json:"updated_at" comment:"记录更新时间"`
	// 删除时间字段，类型为 time.Time
	DeletedAt time.Time `json:"deleted_at" comment:"记录删除时间"`
	// 是否删除字段，类型为 bool
	IsDeleted bool `json:"is_deleted" comment:"标记是否已删除"`
}
