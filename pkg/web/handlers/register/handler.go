package register

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/controllers"
	"github.com/quxiaolong/account/pkg/errs"
	"github.com/quxiaolong/account/pkg/models"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"net/http"
)


func SendRegDigitalHandler (c *gin.Context) {
	var payload RegisterSchema
	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errs.BadArgs.WithData(err))
	}

	user := &models.User{}
	storage.Mysql.GetSlave().Where("phone = ?", payload.Phone).First(user)

	if user.ID != 0 {
		panic(errs.AlreadyExist.WithData(errs.BaseError{Reason:  "请更换手机号或直接登录"}))
	}

	if err := controllers.SendRegDigital(payload.Phone); err != nil {
		panic(errs.DigitalFailed.WithData(err))
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}


func VerifyDigitalHandler (c *gin.Context) {
	var payload VerifyDigitalSchema

	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errs.BadArgs.WithData(err))
	}

	if !controllers.VerifyDigital(payload.Phone, payload.Digital){
		panic(errs.VerifyDigitalFailed)
	}

	if err := controllers.CreateUser(payload.Phone); err != nil {
		panic("")
	}

	// TODO Set AccessToken to Cookie or response content
	// Web: x_l0: access token -> cookie
	// Mobile or WecChat: {"access_token": "access_token", expire_at: "", refresh_token: ""}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}