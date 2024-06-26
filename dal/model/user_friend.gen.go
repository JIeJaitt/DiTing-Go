// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserFriend = "user_friend"

// UserFriend 用户联系人表
type UserFriend struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                             // id
	UID          int64     `gorm:"column:uid;not null;comment:uid" json:"uid"`                                               // uid
	FriendUID    int64     `gorm:"column:friend_uid;not null;comment:好友uid" json:"friend_uid"`                               // 好友uid
	DeleteStatus int32     `gorm:"column:delete_status;not null;comment:逻辑删除(0-正常,1-删除)" json:"delete_status"`               // 逻辑删除(0-正常,1-删除)
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP(3);comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP(3);comment:修改时间" json:"update_time"` // 修改时间
}

// TableName UserFriend's table name
func (*UserFriend) TableName() string {
	return TableNameUserFriend
}
