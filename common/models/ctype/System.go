package ctype

import (
	"database/sql/driver"
	"encoding/json"
)

// SystemPrompt 系统提示消息
type SystemPrompt struct {
	// 系统提示（0: 涉黄，1: 政治，2: 色情，3: 暴力，4: 恐怖，5: 反动，6: 其他）
	Message string `gorm:"column:system_message;comment:'系统提示消息（0: 涉黄，1: 政治，2: 色情，3: 暴力，4: 恐怖，5: 反动，6: 其他）'" json:"message"`
}

// Scan 从数据库中检索数据时使用Scan方法
func (c *SystemPrompt) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

// Value 将数据准备好以存储到数据库中
func (c SystemPrompt) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
