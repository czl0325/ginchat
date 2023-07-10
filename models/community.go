package models

import "gorm.io/gorm"

type Community struct {
	Name    string
	OwnerId uint
	Img     string
	Desc    string
	gorm.Model
}

func (community *Community) TableName() string {
	return "community"
}
