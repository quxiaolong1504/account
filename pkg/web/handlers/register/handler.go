package register

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/controllers"
	"github.com/quxiaolong/account/pkg/errs"
	"github.com/quxiaolong/account/pkg/models"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"net/http"
)

func SendRegDigitalHandler(c *gin.Context) {
	var payload RegisterSchema
	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errs.BadArgs.WithData(err))
	}

	if err := controllers.SendRegDigital(payload.Phone); err != nil {
		panic(errs.DigitalFailed.WithData(err))
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func VerifyDigitalHandler(c *gin.Context) {
	var payload VerifyDigitalSchema

	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errs.BadArgs.WithData(err))
	}

	if !controllers.VerifyDigital(payload.Phone, payload.Digital) {
		panic(errs.VerifyDigitalFailed)
	}
	user, isNew := controllers.GetOrCreateUser(payload.Phone)
	controllers.WriteSession2Resp(user, isNew, c)
}

func RefreshToken(c *gin.Context) {
	uid, _ := c.Get("uid")
	user := &models.User{}
	storage.Mysql.GetSlave().Where("uid = ?", uid).First(user)
	controllers.WriteSession2Resp(user, false, c)
}

func HelloHandler(c *gin.Context) {
	uid := c.Value("uid")
	c.JSON(http.StatusOK, gin.H{"uid": uid})
}
