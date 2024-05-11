package ctype

import (
	"database/sql/driver"
	"encoding/json"
)

// VerificationQuestion 验证问题、答案
type VerificationQuestion struct {
	Question1 *string `json:"question1"` // 验证问题
	Question2 *string `json:"question2"` // 验证问题
	Question3 *string `json:"question3"` // 验证问题
	Answer1   *string `json:"answer1"`   // 验证答案
	Answer2   *string `json:"answer2"`   // 验证答案
	Answer3   *string `json:"answer3"`   // 验证答案
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
