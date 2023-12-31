package models

import "gorm.io/gorm"

// Contact 人员关系
type Contact struct {
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁 /群 ID
	Type     int  //对应的类型  1好友  2群  3xx
	Desc     string
	gorm.Model
}

func (contact *Contact) TableName() string {
	return "contact"
}
