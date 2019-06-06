package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/models"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"net/http"
	"time"
)

func WriteSession2Resp(user *models.User, isNew bool,  c *gin.Context) {
	sess := storage.NewSession(user.UID, c.Writer)
	c.JSON(http.StatusOK, gin.H{"is_new": isNew, "token": gin.H{
		"token": sess.ID(),
		"expired_at": time.Now().Add(sess.Timeout()),
		"token_name": "q_x0",}})
}