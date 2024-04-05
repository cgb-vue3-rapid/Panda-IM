package models

import (
	"akita/panda-im/common/models"
	"akita/panda-im/common/models/ctype"
	userModels "akita/panda-im/service/user/rpc/models"
)

// GroupChatModel 群聊消息表
type GroupChatModel struct {
	models.Model
	SendUserID     int64                 `gorm:"column:send_user_id;comment:'发送用户ID';not null;uniqueIndex" json:"send_user_id"` // 发送用户ID
	SendUserModel  *userModels.UserModel `gorm:"foreignKey:SendUserID;references:ID" json:"send_user_model"`
	GroupID        int64                 `gorm:"column:group_id;comment:'群ID';not null" json:"group_id"`
	GroupModel     GroupModel            `gorm:"foreignKey:GroupID;references:ID" json:"group_model"`
	MessageType    int8                  `gorm:"column:message_type;comment:'消息类型(0: 文本类型，1: 图片消息，2: 视频消息，3: 文件消息，4: 语音消息，5: 语言通话，6: 视频通话，7: 撤回消息，8: 回复消息，9: 引用消息，10：@消息)';not null" json:"message_type"`
	MessageContent string                `gorm:"column:message_content;comment:'消息内容'" json:"message_content"` // 消息内容
	MessagePreview string                `gorm:"column:message_preview;comment:'消息预览'" json:"message_preview"` // 消息预览
	SystemPrompt   ctype.SystemPrompt    `json:"system_prompt"`                                                // 系统提示
	//SendTime       int64  `gorm:"column:send_time;comment:'发送时间';not null" json:"send_time"`             // 发送时间
}

func (m *GroupChatModel) TableName() string {
	return "group_chat"
}
