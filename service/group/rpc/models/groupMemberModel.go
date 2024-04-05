package models

import (
	"akita/panda-im/common/models"
	userModels "akita/panda-im/service/user/rpc/models"
)

// GroupMemberModel 群成员表
type GroupMemberModel struct {
	models.Model
	GroupId            uint                 `gorm:"column:group_id;comment:'群ID'" json:"groupId"` // 群ID
	GroupModel         GroupModel           `gorm:"foreignKey:GroupId;not null" json:"-"`
	UserModel          userModels.UserModel `gorm:"foreignKey:UserId" json:"-"`
	UserId             uint                 `gorm:"column:user_id;comment:'用户ID'" json:"userId"`                                // 用户ID
	Role               int8                 `gorm:"column:role;comment:'角色(0:普通成员,1:管理员,2:群主)';not null;default:0" json:"role"` // 角色(0:普通成员,1:管理员,2:群主)
	MemberNickName     string               `gorm:"column:member_nick_name;comment:'群昵称'" json:"nickName"`                      // 群昵称
	ForbiddenWordsTime int64                `gorm:"column:forbidden_words_time;comment:'禁言时间'" json:"forbiddenWordsTime"`       // 禁言时间
}

func (g *GroupMemberModel) TableName() string {
	return "group_member"
}
