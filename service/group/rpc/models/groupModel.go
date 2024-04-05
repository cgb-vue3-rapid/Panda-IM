package models

import (
	"akita/panda-im/common/models"
	"akita/panda-im/common/models/ctype"
)

// GroupModel 群组表结构
type GroupModel struct {
	models.Model
	//GroupInfo    `json:"groupInfo"`
	//GroupSetting `json:"groupSetting"`
	Title                string                       `gorm:"column:title;comment:'群名，非空且加索引'" json:"title" gorm:"not null;index"`
	Abstract             string                       `gorm:"column:abstract;comment:'简介'" json:"abstract"`
	Avatar               string                       `gorm:"column:avatar;comment:'群头像'" json:"avatar"`
	Creator              int                          `gorm:"column:creator;comment:'群主，非空且加索引'" json:"creator" gorm:"not null;index"`
	IsSearch             bool                         `gorm:"column:is_search;comment:'是否可以被搜索'" json:"isSearch"`
	Verification         int8                         `gorm:"column:verification;comment:'群验证方式'" json:"verification"`
	VerificationQuestion []ctype.VerificationQuestion `gorm:"column:verification_question;comment:'验证问题'" json:"verification_question"`
	IsInvitation         bool                         `gorm:"column:is_invitation;comment:'是否可以被邀请'; default:true" json:"is_invitation"`
	ISTemporarySession   bool                         `gorm:"column:is_temporary_session;comment:'是否是临时会话'" json:"is_temporary_session"`
	ISDisableSendMsg     bool                         `gorm:"column:is_disable_send_msg;comment:'是否禁止发送消息'" json:"is_disable_send_msg"`
	Size                 int32                        `gorm:"column:size;comment:'群规模';not null;default:100'" json:"size"` // 群规模（100、500、1000、1500、2000）
}

func (G *GroupModel) TableName() string {
	return "group"
}

//// GroupInfo  群信息
//type GroupInfo struct {
//	Title    string `gorm:"column:title;comment:'群名，非空且加索引'" json:"title" gorm:"not null;index"`
//	Abstract string `gorm:"column:abstract;comment:'简介'" json:"abstract"`
//	Avatar   string `gorm:"column:avatar;comment:'群头像'" json:"avatar"`
//	Creator  int    `gorm:"column:creator;comment:'群主，非空且加索引'" json:"creator" gorm:"not null;index"`
//}
//
//// GroupSetting 群设置
//type GroupSetting struct {
//	IsSearch             bool                         `gorm:"column:is_search;comment:'是否可以被搜索'" json:"isSearch"`
//	Verification         int8                         `gorm:"column:verification;comment:'群验证方式'" json:"verification"`
//	VerificationQuestion []ctype.VerificationQuestion `gorm:"column:verification_question;comment:'验证问题'" json:"verification_question"`
//	IsInvitation         bool                         `gorm:"column:is_invitation;comment:'是否可以被邀请'; default:true" json:"is_invitation"`
//	ISTemporarySession   bool                         `gorm:"column:is_temporary_session;comment:'是否是临时会话'" json:"is_temporary_session"`
//	ISDisableSendMsg     bool                         `gorm:"column:is_disable_send_msg;comment:'是否禁止发送消息'" json:"is_disable_send_msg"`
//}

//func (G *GroupModel) TableName() string {
//	return "group"
//}
