package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/controllers"
	"github.com/quxiaolong/account/pkg/errs"
	"net/http"
)


type RegisterSchema struct {
	Phone string `json:"phone" binding:"required"`
}

func SendRegDigitalHandler (c *gin.Context) {
	var payload RegisterSchema
	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errs.BadArgs.WithData(err))
	}

	if err := digital.SendRegDigital(payload.Phone); err != nil {
		panic(errs.DigitalFailed.WithData(err))
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}


type VerifyDigitalSchema struct {
	Phone string `json:"phone" binding:"required"`
	Digital string `json:"digital" binding:"required"`

}
func VerifyDigitalHandler (c *gin.Context) {
	var payload VerifyDigitalSchema

	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(errs.BadArgs.WithData(err))
	}

	if !digital.VerifyDigital(payload.Phone, payload.Digital){
		panic(errs.VerifyDigitalFailed)
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}