package models

import (
	"akita/panda-im/common/models"
	"akita/panda-im/common/models/ctype"
	userModels "akita/panda-im/service/user/rpc/models"
)

// ChatModel 聊天记录表
type ChatModel struct {
	models.Model
	SendUserID      int64                `gorm:"column:send_user_id;comment:'发送消息的用户 ID';not null" json:"send_user_id"` // 发送消息的用户 ID
	SendUserModel   userModels.UserModel `gorm:"foreignKey:SendUserID;references:ID" json:"send_user_model"`
	RevUserID       int                  `gorm:"column:rev_user_id;comment:'接收消息的用户 ID';not null" json:"rev_user_id"` // 接收用户 ID
	MessageType     int8                 `gorm:"column:message_type;comment:'消息类型(0: 文本类型，1: 图片消息，2: 视频消息，3: 文件消息，4: 语音消息，5: 语言通话，6: 视频通话，7: 撤回消息，8: 回复消息，9: 引用消息)';not null" json:"message_type"`
	MessagesPreview string               `gorm:"column:messages_preview;comment:'消息预览'" json:"messages_preview"` // 消息预览
	MessageContent  ctype.MessageType    `gorm:"column:message_content;comment:'消息内容'" json:"message_content"`   // 消息内容
	SystemPrompt    ctype.SystemPrompt   `gorm:"column:system_prompt;comment:'系统提示'" json:"system_prompt"`       // 系统提示（0: 涉黄，1: 政治，2: 色情，3: 暴力，4: 恐怖，5: 反动，6: 其他）
}

func (C *ChatModel) TableName() string {
	return "chat"
}
