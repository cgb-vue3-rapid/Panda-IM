package ctype

import (
	"database/sql/driver"
	"encoding/json"
)

// VerificationQuestion 验证问题、答案
type VerificationQuestion struct {
	Question string `gorm:"column:question;comment:'验证问题'" json:"question"` // 验证问题
	Answer   string `gorm:"column:answer;comment:'验证答案'" json:"answer"`     // 验证答案
}

// Scan 从数据库中检索数据时使用Scan方法
func (c *VerificationQuestion) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

// Value 将数据准备好以存储到数据库中
func (c VerificationQuestion) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
