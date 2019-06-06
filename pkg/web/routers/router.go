package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/web/handlers/register"
	"github.com/quxiaolong/account/pkg/web/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(middlewares.Recover())
	r.Use(middlewares.Authentication())
	v1 := r.Group("/v1")

	auth := v1.Group("/auth")
	{
		// register handler for router
		auth.POST("/digits", register.SendRegDigitalHandler)
		auth.POST("/validate/digits", register.VerifyDigitalHandler)
		auth.GET("/hello", register.HelloHandler)
	}


	return r
}
