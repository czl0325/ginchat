package models

import "gorm.io/gorm"

// Group 群信息
type Group struct {
	Name    string
	OwnerId uint
	Icon    string
	Type    int
	Desc    string
	gorm.Model
}

func (group *Group) TableName() string {
	return "group"
}
