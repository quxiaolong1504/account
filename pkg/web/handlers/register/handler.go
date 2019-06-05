package register

import (
	"github.com/gin-gonic/gin"
	"github.com/icza/session"
	"github.com/quxiaolong/account/pkg/controllers"
	"github.com/quxiaolong/account/pkg/errs"
	"github.com/quxiaolong/account/pkg/models"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"net/http"
	"time"
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
	user, err := controllers.CreateUser(payload.Phone)
	if  err != nil {
		panic("")
	}

	sess := session.NewSessionOptions(&session.SessOptions{
		Timeout: time.Hour * 24 * 30,
	})
	sess.SetAttr("uid", user.UID)
	storage.SessMgr.Add(sess, c.Writer)
	c.JSON(http.StatusOK, gin.H{ "token": sess.ID(), "expired_at": time.Now().Add(sess.Timeout())})
}

func HelloHandler(c *gin.Context) {
	uid := c.Value("uid")
	c.JSON(http.StatusOK, gin.H{"uid": uid})
}