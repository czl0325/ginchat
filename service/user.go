package service

import (
	"crypto/md5"
	"fmt"
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

// UserRegister
// @Summary 注册用户
// @Tags 用户模块
// @param name formData string true "用户名"
// @param password formData string false "密码"
// @param passwordAgain formData string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	name := c.PostForm("name")
	pwd1 := c.PostForm("password")
	pwd2 := c.PostForm("passwordAgain")
	if pwd1 != pwd2 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "两次密码不一致",
		})
		return
	}
	user := models.FindUserByName(name)
	if user.ID != 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名已存在",
		})
		return
	}
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Name = name
	user.Salt = salt
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(pwd1+salt)))
	user.LoginTime = time.Now()
	user.LogoutTime = time.Now()
	user.HeartbeatTime = time.Now()
	err := models.CreateUser(user)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "注册失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "注册成功",
		"data":    user,
	})
}
