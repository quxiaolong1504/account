package wechat

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/controllers"
	"github.com/quxiaolong/account/pkg/errs"
	"github.com/quxiaolong/account/pkg/utils"
)

func Login(c *gin.Context) {
	code := c.Param("code")
	ret, err := utils.WeChatCli.GetMiniProgram().Code2Session(code)
	if err != nil {
		panic(errs.WeChatLoginFailed)
	}
	user, isNew := controllers.GetOrCrateUserByUnionid(ret.UnionID)
	controllers.WriteSession2Resp(user, isNew, c)
}