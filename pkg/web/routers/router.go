package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/web/handlers/register"
	wx "github.com/quxiaolong/account/pkg/web/handlers/wechat"
	"github.com/quxiaolong/account/pkg/web/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(middlewares.Recover())
	r.Use(middlewares.Authentication())
	r.Use(middlewares.Monitor())
	v1 := r.Group("/v1")

	auth := v1.Group("/auth")
	{
		// register handler for router
		auth.POST("/digits", register.SendRegDigitalHandler)
		auth.POST("/validate/digits", register.VerifyDigitalHandler)
		auth.POST("/refresh/token", register.RefreshToken)

	}

	weChat := v1.Group("/wx")
	{
		weChat.POST("/login", wx.Login)
	}

	r.GET("/hello", register.HelloHandler)


	return r
}
