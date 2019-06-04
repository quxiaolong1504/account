package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/controllers/login"
	"net/http"
)

type RegisterHandler struct {

}

type RegisterSchema struct {
	Phone string `json:"phone" binding:"required"`
}

func (h RegisterHandler) Post (c *gin.Context) {
	var payload RegisterSchema
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := login.SendVerifyCode(c, payload.Phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}