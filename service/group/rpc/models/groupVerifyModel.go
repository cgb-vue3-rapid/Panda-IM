package models

import (
	"akita/panda-im/common/models"
	"akita/panda-im/common/models/ctype"
	userModels "akita/panda-im/service/user/rpc/models"
)

// GroupVerifyModel 加群验证表
type GroupVerifyModel struct {
	models.Model
	UserID               uint                         `gorm:"column:user_id;comment:'用户ID';not null" json:"user_id"`
	GroupID              uint                         `gorm:"column:group_id;comment:'群ID';not null" json:"group_id"`
	GroupModel           GroupModel                   `gorm:"foreignKey:GroupId;not null" json:"-"`
	UserModel            userModels.UserModel         `gorm:"foreignKey:UserId" json:"-"`
	Status               int8                         `gorm:"column:status;comment:'验证状态（1同意，2拒绝，3忽略）';not null" json:"status"`        // 验证状态（1同意，2拒绝，3忽略）
	AdditionalMessages   string                       `gorm:"column:additional_messages;comment:'附加消息'" json:"additional_messages"`    // 附加消息
	VerificationQuestion []ctype.VerificationQuestion `gorm:"column:VerificationQuestion;comment:'验证问题'" json:"verification_question"` // 验证问题为3和4的时候需要
}

// GroupExitModel 退群表
type GroupExitModel struct {
	models.Model
	UserID     uint                 `gorm:"column:user_id;comment:'用户ID';not null" json:"user_id"`  // 用户ID
	GroupID    uint                 `gorm:"column:group_id;comment:'群ID';not null" json:"group_id"` // 群ID
	GroupModel GroupModel           `gorm:"foreignKey:GroupId;not null" json:"-"`
	UserModel  userModels.UserModel `gorm:"foreignKey:UserId" json:"-"`
	Reason     string               `gorm:"column:reason;comment:'退群原因'" json:"reason"`       // 退群原因
	Timestamp  int64                `gorm:"column:timestamp;comment:'退群时间'" json:"timestamp"` // 退群时间
}

func (m *GroupVerifyModel) TableName() string {
	return "group_verify"
}

func (m *GroupExitModel) TableName() string {
	return "group_exit"
}
