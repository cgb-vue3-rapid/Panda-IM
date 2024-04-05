package models

import (
	"akita/panda-im/common/models"
	"akita/panda-im/common/models/ctype"
)

// FriendVerifyModel 好友验证表
type FriendVerifyModel struct {
	models.Model
	SendUserID           uint                         `gorm:"column:send_user_id;comment:'发送方ID';not null;uniqueIndex" json:"send_user_id"`
	RevUserID            uint                         `gorm:"column:rev_user_id;comment:'接收方ID';not null;uniqueIndex" json:"rev_user_id"`
	SendUserModel        UserModel                    `gorm:"foreignKey:SendUserId" json:"-"`
	RevUserModel         UserModel                    `gorm:"foreignKey:RevUserId" json:"-"`
	Status               int8                         `gorm:"column:status;comment:'验证状态（1同意，2拒绝，3忽略）';not null" json:"status"`        // 验证状态（1同意，2拒绝，3忽略）
	AdditionalMessages   string                       `gorm:"column:additional_messages;comment:'附加消息'" json:"additional_messages"`    // 附加消息
	VerificationQuestion []ctype.VerificationQuestion `gorm:"column:VerificationQuestion;comment:'验证问题'" json:"verification_question"` // 验证问题为3和4的时候需要
}

func (U *FriendVerifyModel) TableName() string {
	return "friend_verify"
}
