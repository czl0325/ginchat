package service

import (
	"crypto/md5"
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"strconv"
	"time"
)

// UserRegister
// @Summary 注册用户
// @Tags 用户模块
// @param name formData string true "用户名"
// @param password formData string true "密码"
// @param passwordAgain formData string true "确认密码"
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

// UserLogin
// @Summary 登录用户
// @Tags 用户模块
// @param name formData string true "用户名"
// @param password formData string true "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("password")
	user := models.FindUserByNameAndPwd(name, pwd)
	if user.ID == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名或密码错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    user,
	})
}

// UserUpdate
// @Summary 更新用户
// @Tags 用户模块
// @Param id path int true "用户id"
// @param name formData string false "用户名"
// @param phone formData string false "手机"
// @param email formData string false "邮箱"
// @Success 200 {string} json{"code","message"}
// @Router /user/{id} [put]
func UserUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	avatar := c.PostForm("avatar")
	user := models.FindUserById(uint(id))
	if user.ID == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户不存在",
		})
		return
	}
	if name != "" {
		user := models.FindUserByName(name)
		if user.ID != 0 {
			c.JSON(200, gin.H{
				"code":    -1,
				"message": "用户名已存在",
			})
			return
		}
		user.Name = name
	}
	if phone != "" {
		user.Phone = phone
	}
	if email != "" {
		user.Email = email
	}
	if avatar != "" {
		user.Avatar = avatar
	}
	validate := validator.New()
	validate.RegisterValidation("phone", utils.PhoneValidator)
	if err := validate.Struct(user); err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "参数校验错误",
		})
		return
	}
	if err := models.UpdateUser(user).Error; err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "更新失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "更新成功",
	})
}

// UserDelete
// @Summary 删除用户
// @Tags 用户模块
// @Param id path int true "用户id"
// @Success 200 {string} json{"code","message"}
// @Router /user/{id} [delete]
func UserDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	user.ID = uint(id)
	err := models.DeleteUser(user).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "删除失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "删除成功",
	})
}
