package main

import (
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCreateSql(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:1111@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("创建User表失败")
	}
	err = db.AutoMigrate(&models.Group{})
	if err != nil {
		panic("创建Group表失败")
	}
	err = db.AutoMigrate(&models.Message{})
	if err != nil {
		panic("创建Message表失败")
	}
	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		panic("创建Contact表失败")
	}
	err = db.AutoMigrate(&models.Community{})
	if err != nil {
		panic("创建Community表失败")
	}
}
