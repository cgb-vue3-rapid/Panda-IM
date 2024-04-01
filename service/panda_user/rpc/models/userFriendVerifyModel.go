package models

import "akita/panda-im/common/models"

// FriendVerifyModel 是好友验证表
type FriendVerifyModel struct {
	models.Model
	AddFriendBySendUserID int                   `gorm:"column:add_friend_send_user_id;comment:'发起验证方'" json:"add_friend_by_send_user_id"` // 发起验证方
	AddFriendByRevUserID  int                   `gorm:"column:add_friend_rev_user_id;comment:'接受验证方'" json:"add_friend_by_rev_user_id"`   // 接受验证方
	AddFriendStatus       int8                  `gorm:"column:status;comment:'验证状态（1同意，2拒绝，3忽略）'" json:"status"`                          // 验证状态（1同意，2拒绝，3忽略）
	AdditionalMessages    string                `gorm:"column:additional_messages;comment:'附加消息'" json:"additional_messages"`             // 附加消息
	VerificationQuestion  *VerificationQuestion `json:"verification_question"`                                                            // 验证问题为3和4的时候需要
}

func (U *FriendVerifyModel) TableName() string {
	return "friend_verify"
}
