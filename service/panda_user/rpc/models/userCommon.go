package models

// VerificationQuestion 验证问题、答案
type VerificationQuestion struct {
	Question string `gorm:"column:question;comment:'验证问题'" json:"question"` // 验证问题
	Answer   string `gorm:"column:answer;comment:'验证答案'" json:"answer"`     // 验证答案
}
