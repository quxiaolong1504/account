package register

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/controllers"
	"github.com/quxiaolong/account/pkg/errs"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"net/http"
	"time"
)


func SendRegDigitalHandler (c *gin.Context) {
	var payload RegisterSchema
	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errs.BadArgs.WithData(err))
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
	user, isNew := controllers.GetOrCreateUser(payload.Phone)

	sess := storage.WriteNewSession(user.UID, c.Writer)
	c.JSON(http.StatusOK, gin.H{"is_new": isNew, "token": gin.H{
		"token": sess.ID(),
		"expired_at": time.Now().Add(sess.Timeout()),
		"token_name": "q_x0",}})
}

func HelloHandler(c *gin.Context) {
	uid := c.Value("uid")
	c.JSON(http.StatusInternalServerError, gin.H{"uid": uid})
}