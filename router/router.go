package router

import (
	"ginchat/docs"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.Index)

	r.POST("/user/register", service.UserRegister)
	r.POST("/user/login", service.UserLogin)
	r.PUT("/user/:id", service.UserUpdate)
	r.DELETE("/user/:id", service.UserDelete)

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
