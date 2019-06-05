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

	// register handler for router
	r.POST("/register", register.SendRegDigitalHandler)
	r.POST("/register/verify", register.VerifyDigitalHandler)
	return r
}
