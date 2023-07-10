package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Name          string `size:20;unique_index;not null`
	Password      string `not null`
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time `default:NULL`
	HeartbeatTime time.Time `default:NULL`
	LogoutTime    time.Time `default:NULL`
	IsLogout      bool      `default:false`
	DeviceInfo    string
	gorm.Model
}

func (user *User) TableName() string {
	return "user"
}

func CreateUser(user User) error {
	return utils.Db.Create(&user).Error
}

func FindUserByName(name string) User {
	user := User{}
	utils.Db.Where("name = ?", name).First(&user)
	return user
}
