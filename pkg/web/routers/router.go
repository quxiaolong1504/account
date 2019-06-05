package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/web/handlers"
	"github.com/quxiaolong/account/pkg/web/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(middlewares.Recover())

	// register handler for router
	r.POST("/register", handlers.SendRegDigitalHandler)
	r.POST("/register/verify", handlers.VerifyDigitalHandler)
	return r
}
