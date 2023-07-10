package models

import (
	"crypto/md5"
	"fmt"
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Name          string `size:20;unique_index;not null`
	Password      string `not null`
	Phone         string `validate:"phone"`
	Email         string `validate:"email"`
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

func FindUserById(id uint) User {
	user := User{}
	utils.Db.Where("id = ?", id).First(&user)
	return user
}

func FindUserByNameAndPwd(name, password string) User {
	user := User{}
	utils.Db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		return user
	}
	pwdMd5 := fmt.Sprintf("%x", md5.Sum([]byte(password+user.Salt)))
	user = User{}
	utils.Db.Where("name = ? and password=?", name, pwdMd5).First(&user)
	return user
}

func UpdateUser(user User) *gorm.DB {
	return utils.Db.Updates(&user)
}

func DeleteUser(user User) *gorm.DB {
	return utils.Db.Delete(&user)
}
